package repo

import (
	"c_program_edu-gin/internal/app/storage/model"
	ctx "c_program_edu-gin/internal/pkg/context"
	"gorm.io/gorm"
)

type QuestionRepo struct {
	ctx.Repo[model.Question]
}

func NewQuestions(db *gorm.DB) *QuestionRepo {
	return &QuestionRepo{Repo: ctx.NewRepo[model.Question](db)}
}
