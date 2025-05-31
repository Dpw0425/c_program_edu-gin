package service

import (
	"c_program_edu-gin/internal/app/schema"
	"c_program_edu-gin/internal/app/storage/repo"
	myErr "c_program_edu-gin/pkg/error"
	"context"
)

var _ ICompetitionService = (*CompetitionService)(nil)

type ICompetitionService interface {
	List(ctx context.Context) ([]*schema.WebCompetitionItem, int, error)
}

type CompetitionService struct {
	CompetitionRepo *repo.CompetitionRepo
}

func (c *CompetitionService) List(ctx context.Context) ([]*schema.WebCompetitionItem, int, error) {
	db := c.CompetitionRepo.DB.WithContext(ctx)
	var items []*schema.WebCompetitionItem

	err := db.Table("competitions").Scan(&items).Error
	if err != nil {
		return nil, 0, myErr.NotFound("", err.Error())
	}

	var count int64
	db.Table("competitions").Count(&count)

	var quantity int64
	for _, item := range items {
		db.Table("cpt_ques").Where("competition_id = ?", item.ID).Count(&quantity)
		item.Quantity = int(quantity)
	}

	return items, int(count), nil
}
