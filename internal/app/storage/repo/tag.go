package repo

import (
	"c_program_edu-gin/internal/app/storage/model"
	ctx "c_program_edu-gin/internal/pkg/context"
	"context"
	"gorm.io/gorm"
)

type TagRepo struct {
	ctx.Repo[model.Tag]
}

func NewTags(db *gorm.DB) *TagRepo {
	return &TagRepo{Repo: ctx.NewRepo[model.Tag](db)}
}

func (t *TagRepo) IsExist(ctx context.Context, name string) bool {
	if len(name) == 0 {
		return false
	}

	isExist, _ := t.Repo.IsExist(ctx, "name = ?", name)
	return isExist
}

func (t *TagRepo) Create(ctx context.Context, tag *model.Tag) error {
	if err := t.Repo.Create(ctx, tag); err != nil {
		return err
	}

	return nil
}
