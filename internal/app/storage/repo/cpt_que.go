package repo

import (
	"c_program_edu-gin/internal/app/storage/model"
	ctx "c_program_edu-gin/internal/pkg/context"
	"gorm.io/gorm"
)

type CptQueRepo struct {
	ctx.Repo[model.CptQue]
}

func NewCptQues(db *gorm.DB) *CptQueRepo {
	return &CptQueRepo{Repo: ctx.NewRepo[model.CptQue](db)}
}
