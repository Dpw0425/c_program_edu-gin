package repo

import (
	"c_program_edu-gin/internal/app/storage/model"
	ctx "c_program_edu-gin/internal/pkg/context"
	"gorm.io/gorm"
)

type BankQueRepo struct {
	ctx.Repo[model.BankQue]
}

func NewBankQue(db *gorm.DB) *BankQueRepo {
	return &BankQueRepo{Repo: ctx.NewRepo[model.BankQue](db)}
}
