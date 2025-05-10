package repo

import (
	"c_program_edu-gin/internal/app/storage/model"
	ctx "c_program_edu-gin/internal/pkg/context"
	"context"
	"gorm.io/gorm"
)

type QuestionRepo struct {
	ctx.Repo[model.Question]
}

func NewQuestions(db *gorm.DB) *QuestionRepo {
	return &QuestionRepo{Repo: ctx.NewRepo[model.Question](db)}
}

func (q *QuestionRepo) IsExist(ctx context.Context, title string) bool {
	if len(title) == 0 {
		return false
	}

	isExist, _ := q.Repo.IsExist(ctx, "title = ?", title)

	return isExist
}

func (q *QuestionRepo) Create(ctx context.Context, question *model.Question) error {
	if err := q.Repo.Create(ctx, question); err != nil {
		return err
	}

	return nil
}
