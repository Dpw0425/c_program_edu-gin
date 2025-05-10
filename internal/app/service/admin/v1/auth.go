package admin_service

import (
	"c_program_edu-gin/internal/app/schema"
	"c_program_edu-gin/internal/app/storage/model"
	"c_program_edu-gin/internal/app/storage/repo"
	myErr "c_program_edu-gin/pkg/error"
	admin "c_program_edu-gin/schema/genproto/admin/v1/auth"
	"c_program_edu-gin/utils/encrypt"
	"context"
	"strconv"
)

var _ IAuthService = (*AuthService)(nil)

type IAuthService interface {
	UserList(ctx context.Context, request *admin.UserListRequest) ([]*schema.UserItem, int, error)
	UpdateUser(ctx context.Context, request *admin.UpdateUserRequest) error
	DeleteUser(ctx context.Context, request *admin.DeleteUserRequest) error
	AdminList(ctx context.Context, request *admin.AdminListRequest) ([]*schema.AdminItem, int, error)
	AddAdmin(ctx context.Context, request *admin.AddAdminRequest) error
	UpdateAdmin(ctx context.Context, request *admin.UpdateAdminRequest) error
	DeleteAdmin(ctx context.Context, request *admin.DeleteAdminRequest) error
}

type AuthService struct {
	UserRepo  *repo.UserRepo
	AdminRepo *repo.AdminRepo
}

func (a *AuthService) UserList(ctx context.Context, request *admin.UserListRequest) ([]*schema.UserItem, int, error) {
	db := a.UserRepo.DB.WithContext(ctx)
	var items []*schema.UserItem

	err := db.Table("users").Select("user_id", "user_name", "student_id", "grade", "status", "email", "avatar").Where("user_name LIKE ?", "%"+request.Search+"%").Or("student_id LIKE ?", "%"+request.Search+"%").Limit(int(request.Number)).Offset(int((request.Page - 1) * request.Number)).Scan(&items).Error
	if err != nil {
		return nil, 0, myErr.NotFound("", err.Error())
	}

	var count int64
	db.Table("users").Count(&count)

	return items, int(count), nil
}

func (a *AuthService) UpdateUser(ctx context.Context, request *admin.UpdateUserRequest) error {
	uid, _ := strconv.ParseInt(request.UserId, 10, 64)

	_, err := a.UserRepo.UpdateByUserID(ctx, uid, map[string]any{
		"user_name": request.UserName,
		"status":    request.Status,
	})
	if err != nil {
		return myErr.BadRequest("", "修改失败！")
	}

	return nil
}

func (a *AuthService) DeleteUser(ctx context.Context, request *admin.DeleteUserRequest) error {
	db := a.AdminRepo.DB.WithContext(ctx)
	return db.Unscoped().Where("user_id = ?", request.UserId).Delete(&model.User{}).Error
}

func (a *AuthService) AdminList(ctx context.Context, request *admin.AdminListRequest) ([]*schema.AdminItem, int, error) {
	db := a.AdminRepo.DB.WithContext(ctx)
	var items []*schema.AdminItem

	err := db.Table("admins").Select("id", "user_name", "teacher_id", "permission", "status").Where("user_name LIKE ?", "%"+request.Search+"%").Or("teacher_id LIKE ?", "%"+request.Search+"%").Limit(int(request.Number)).Offset(int((request.Page - 1) * request.Number)).Scan(&items).Error
	if err != nil {
		return nil, 0, myErr.NotFound("", err.Error())
	}

	var count int64
	db.Table("admins").Count(&count)

	return items, int(count), nil
}

func (a *AuthService) AddAdmin(ctx context.Context, request *admin.AddAdminRequest) error {
	if a.AdminRepo.IsExist(ctx, request.TeacherId) {
		return myErr.BadRequest("", "账号已存在！")
	}

	return a.AdminRepo.Create(ctx, &model.Admin{
		UserName:   request.UserName,
		TeacherID:  request.TeacherId,
		Password:   encrypt.HashPassword(request.Password),
		Permission: request.Permission,
	})
}

func (a *AuthService) UpdateAdmin(ctx context.Context, request *admin.UpdateAdminRequest) error {
	_, err := a.AdminRepo.UpdateByID(ctx, uint(request.Id), map[string]any{
		"user_name":  request.UserName,
		"permission": request.Permission,
	})
	if err != nil {
		return myErr.BadRequest("", "修改失败！")
	}

	return nil
}

func (a *AuthService) DeleteAdmin(ctx context.Context, request *admin.DeleteAdminRequest) error {
	db := a.AdminRepo.DB.WithContext(ctx)
	return db.Unscoped().Where("id = ?", request.Id).Delete(&model.Admin{}).Error
}
