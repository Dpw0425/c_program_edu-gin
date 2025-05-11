package repo

import (
	"c_program_edu-gin/internal/app/storage/model"
	ctx "c_program_edu-gin/internal/pkg/context"
	"context"
	"gorm.io/gorm"
)

type TestDataRepo struct {
	ctx.Repo[model.TestData]
}

func NewTestData(db *gorm.DB) *TestDataRepo {
	return &TestDataRepo{Repo: ctx.NewRepo[model.TestData](db)}
}

func (t *TestDataRepo) Create(ctx context.Context, data *model.TestData) error {
	if err := t.Repo.Create(ctx, data); err != nil {
		return err
	}

	return nil
}
