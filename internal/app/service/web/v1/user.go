package service

import (
	"c_program_edu-gin/internal/app/schema"
	"c_program_edu-gin/internal/app/storage/model"
	"c_program_edu-gin/internal/app/storage/repo"
	myErr "c_program_edu-gin/pkg/error"
	"c_program_edu-gin/utils/encrypt"
	"c_program_edu-gin/utils/generator"
	"context"
	"gorm.io/gorm"
)

var _ IUserService = (*UserService)(nil)

type IUserService interface {
	Register(ctx context.Context, register *schema.UserRegister) error
	LoginByPassword(ctx context.Context, login *schema.UserLogin) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByID(ctx context.Context, UserID int64) (*model.User, error)
}

type UserService struct {
	UserRepo *repo.UserRepo
}

func (u *UserService) Register(ctx context.Context, sur *schema.UserRegister) error {
	if u.UserRepo.IsExist(ctx, sur.Email) {
		return myErr.BadRequest("", "账号已存在！")
	}

	return u.UserRepo.Create(ctx, &model.User{
		UserID:   generator.IDGenerator(),
		NickName: sur.NickName,
		Password: encrypt.HashPassword(sur.Password),
		Avatar:   sur.Avatar,
		Email:    sur.Email,
	})
}

func (u *UserService) LoginByPassword(ctx context.Context, sul *schema.UserLogin) (*model.User, error) {
	user, err := u.UserRepo.FindByEmail(ctx, sul.Email)
	if err != nil {
		if myErr.Equal(err, gorm.ErrRecordNotFound) {
			return nil, myErr.BadRequest("", "账号不存在！")
		}

		return nil, myErr.BadRequest("", err.Error())
	}

	if !encrypt.VerifyPassword(user.Password, sul.Password) {
		return nil, myErr.BadRequest("", "密码错误！")
	}

	return user, nil
}

func (u *UserService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := u.UserRepo.FindByEmail(ctx, email)
	if err != nil {
		if myErr.Equal(err, gorm.ErrRecordNotFound) {
			return nil, myErr.BadRequest("", "账号不存在！")
		}

		return nil, myErr.BadRequest("", err.Error())
	}

	return user, nil
}

func (u *UserService) GetUserByID(ctx context.Context, UserID int64) (*model.User, error) {
	user, err := u.UserRepo.FindByID(ctx, UserID)
	if err != nil {
		if myErr.Equal(err, gorm.ErrRecordNotFound) {
			return nil, myErr.BadRequest("", "账号不存在！")
		}

		return nil, myErr.BadRequest("", err.Error())
	}

	return user, nil
}
