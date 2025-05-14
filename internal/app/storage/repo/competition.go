package repo

import (
	"c_program_edu-gin/internal/app/storage/model"
	ctx "c_program_edu-gin/internal/pkg/context"
	"context"
	"gorm.io/gorm"
)

type CompetitionRepo struct {
	ctx.Repo[model.Competition]
}

func NewCompetition(db *gorm.DB) *CompetitionRepo {
	return &CompetitionRepo{Repo: ctx.NewRepo[model.Competition](db)}
}

func (c *CompetitionRepo) IsExist(ctx context.Context, name string) bool {
	if len(name) == 0 {
		return false
	}

	isExist, _ := c.Repo.IsExist(ctx, "name = ?", name)

	return isExist
}

func (c *CompetitionRepo) Create(ctx context.Context, competition *model.Competition) error {
	if err := c.Repo.Create(ctx, competition); err != nil {
		return err
	}

	return nil
}
