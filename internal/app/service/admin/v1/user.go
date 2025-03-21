package admin_service

import (
	"c_program_edu-gin/internal/app/schema"
	"c_program_edu-gin/internal/app/storage/model"
	"c_program_edu-gin/internal/app/storage/repo"
	myErr "c_program_edu-gin/pkg/error"
	"c_program_edu-gin/utils/encrypt"
	"context"
	"gorm.io/gorm"
)

var _ IAdminService = (*AdminService)(nil)

type IAdminService interface {
	Login(ctx context.Context, login *schema.AdminLogin) (*model.Admin, error)
}

type AdminService struct {
	AdminRepo *repo.AdminRepo
}

func (a *AdminService) Login(ctx context.Context, sal *schema.AdminLogin) (*model.Admin, error) {
	admin, err := a.AdminRepo.FindByTeacherID(ctx, sal.UserName)
	if err != nil {
		if myErr.Equal(err, gorm.ErrRecordNotFound) {
			return nil, myErr.BadRequest("", "账号不存在！")
		}

		return nil, myErr.BadRequest("", err.Error())
	}

	if !encrypt.VerifyPassword(admin.Password, sal.Password) {
		return nil, myErr.BadRequest("", "密码错误！")
	}
	return admin, nil
}
