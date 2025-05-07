package admin_service

import (
	"c_program_edu-gin/internal/app/schema"
	"c_program_edu-gin/internal/app/storage/repo"
	myErr "c_program_edu-gin/pkg/error"
	admin "c_program_edu-gin/schema/genproto/admin/v1/auth"
	"context"
)

var _ IAuthService = (*AuthService)(nil)

type IAuthService interface {
	UserList(ctx context.Context, request *admin.UserListRequest) ([]*schema.UserItem, error)
	AdminList(ctx context.Context, request *admin.AdminListRequest) ([]*schema.AdminItem, error)
}

type AuthService struct {
	UserRepo  *repo.UserRepo
	AdminRepo *repo.AdminRepo
}

func (a *AuthService) UserList(ctx context.Context, request *admin.UserListRequest) ([]*schema.UserItem, error) {
	db := a.UserRepo.DB.WithContext(ctx)
	var items []*schema.UserItem

	err := db.Table("users").Select("user_id", "user_name", "student_id", "grade", "status").Where("user_name LIKE ?", "%"+request.Search+"%").Or("student_id LIKE ?", "%"+request.Search+"%").Limit(int(request.Number)).Offset(int((request.Page - 1) * request.Number)).Scan(&items).Error
	if err != nil {
		return nil, myErr.NotFound("", err.Error())
	}

	return items, nil
}

func (a *AuthService) AdminList(ctx context.Context, request *admin.AdminListRequest) ([]*schema.AdminItem, error) {
	db := a.AdminRepo.DB.WithContext(ctx)
	var items []*schema.AdminItem

	err := db.Table("admins").Select("id", "user_name", "teacher_id", "permission", "status").Where("user_name LIKE ?", "%"+request.Search+"%").Or("teacher_id LIKE ?", "%"+request.Search+"%").Limit(int(request.Number)).Offset(int((request.Page - 1) * request.Number)).Scan(&items).Error
	if err != nil {
		return nil, myErr.NotFound("", err.Error())
	}

	return items, nil
}
