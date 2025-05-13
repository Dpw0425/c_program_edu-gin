package repo

import (
	"c_program_edu-gin/internal/app/storage/model"
	ctx "c_program_edu-gin/internal/pkg/context"
	"context"
	"gorm.io/gorm"
)

type BankRepo struct {
	ctx.Repo[model.QuestionBank]
}

func NewBanks(db *gorm.DB) *BankRepo {
	return &BankRepo{Repo: ctx.NewRepo[model.QuestionBank](db)}
}

func (b *BankRepo) IsExist(ctx context.Context, name string) bool {
	if len(name) == 0 {
		return false
	}

	isExist, _ := b.Repo.IsExist(ctx, "name = ?", name)

	return isExist
}
