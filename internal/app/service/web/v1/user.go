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
	Personal(ctx context.Context, UserID int64) (*schema.Personal, []*schema.TeamItem, error)
}

type UserService struct {
	UserRepo *repo.UserRepo
}

func (u *UserService) Register(ctx context.Context, sur *schema.UserRegister) error {
	if u.UserRepo.IsExist(ctx, sur.Email) {
		return myErr.BadRequest("", "账号已存在！")
	}

	if u.UserRepo.IsStudentIDExist(ctx, sur.StudentID) {
		return myErr.BadRequest("", "学号已被绑定！")
	}

	return u.UserRepo.Create(ctx, &model.User{
		UserID:    generator.IDGenerator(),
		UserName:  sur.UserName,
		Password:  encrypt.HashPassword(sur.Password),
		StudentID: sur.StudentID,
		Avatar:    sur.Avatar,
		Email:     sur.Email,
		Grade:     sur.Grade,
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

func (u *UserService) Personal(ctx context.Context, UserID int64) (*schema.Personal, []*schema.TeamItem, error) {
	db := u.UserRepo.DB.WithContext(ctx)

	user, err := u.UserRepo.FindByID(ctx, UserID)
	if err != nil {
		return nil, nil, myErr.BadRequest("", err.Error())
	}

	var teamList []*schema.TeamItem
	db.Table("teams").Where("manager = ?", user.UserID).Or("FIND_IN_SET(?,member)", user.UserID).Scan(&teamList)

	var question []*model.UserQue
	db.Table("user_ques").Where("user_id = ?", user.UserID).Scan(&question)

	var commitTimes int
	for _, item := range question {
		commitTimes += int(item.CommitTimes)
	}

	var competitionTimes int64
	db.Table("competitions").Where("FIND_IN_SET(?,contestant)", user.UserID).Count(&competitionTimes)
	for _, item := range teamList {
		competitionTimes += int64(item.CompetitionTimes)
	}

	return &schema.Personal{
		UserID:           user.UserID,
		UserName:         user.UserName,
		StudentID:        user.StudentID,
		Grade:            user.Grade,
		Status:           user.Status,
		Email:            user.Email,
		Avatar:           user.Avatar,
		CompetitionTimes: int(competitionTimes),
		CommitTimes:      commitTimes,
	}, teamList, nil
}
