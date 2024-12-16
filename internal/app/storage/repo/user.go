package repo

import (
	"c_program_edu-gin/internal/app/storage/model"
	ctx "c_program_edu-gin/internal/pkg/context"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type UserRepo struct {
	ctx.Repo[model.User]
}

func NewUsers(db *gorm.DB) *UserRepo {
	return &UserRepo{Repo: ctx.NewRepo[model.User](db)}
}

// 判断是否存在
func (u *UserRepo) IsExist(ctx context.Context, email string) bool {
	if len(email) == 0 {
		return false
	}

	isExist, _ := u.Repo.IsExist(ctx, "email = ?", email)
	return isExist
}

func (u *UserRepo) Create(ctx context.Context, user *model.User) error {
	if err := u.Repo.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

// 通过邮箱号查询
func (u *UserRepo) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	if len(email) == 0 {
		return nil, fmt.Errorf("email empty")
	}

	return u.Repo.FindByWhere(ctx, "email = ?", email)
}
