package admin_service

import (
	"c_program_edu-gin/internal/app/schema"
	"c_program_edu-gin/internal/app/storage/model"
	"c_program_edu-gin/internal/app/storage/repo"
	myErr "c_program_edu-gin/pkg/error"
	admin "c_program_edu-gin/schema/genproto/admin/v1/competition"
	"context"
	"time"
)

var _ ICompetitionService = (*CompetitionService)(nil)

type ICompetitionService interface {
	List(ctx context.Context, request *admin.CompetitionListRequest) ([]*schema.CompetitionItem, int, error)
	Add(ctx context.Context, userId uint, request *admin.AddCompetitionRequest) error
	Update(ctx context.Context, request *admin.UpdateCompetitionRequest) error
	Delete(ctx context.Context, request *admin.DeleteCompetitionRequest) error
}

type CompetitionService struct {
	CompetitionRepo *repo.CompetitionRepo
}

func (c *CompetitionService) List(ctx context.Context, request *admin.CompetitionListRequest) ([]*schema.CompetitionItem, int, error) {
	db := c.CompetitionRepo.DB.WithContext(ctx)
	var items []*schema.CompetitionItem

	err := db.Table("competitions").Where("name LIKE ?", "%"+request.Search+"%").Limit(int(request.Number)).Offset(int((request.Page - 1) * request.Number)).Scan(&items).Error
	if err != nil {
		return nil, 0, myErr.NotFound("", err.Error())
	}

	var count int64
	db.Table("competitions").Count(&count)

	return items, int(count), nil
}

func (c *CompetitionService) Add(ctx context.Context, userId uint, request *admin.AddCompetitionRequest) error {
	if c.CompetitionRepo.IsExist(ctx, request.Name) {
		return myErr.BadRequest("", "比赛已存在！")
	}

	startTime, _ := time.Parse(time.RFC3339, request.StartTime)
	deadline, _ := time.Parse(time.RFC3339, request.Deadline)
	return c.CompetitionRepo.Create(ctx, &model.Competition{
		Name:       request.Name,
		OwnerID:    userId,
		StartTime:  startTime,
		Deadline:   deadline,
		Category:   int(request.Category),
		Permission: int(request.Permission),
	})
}

func (c *CompetitionService) Update(ctx context.Context, request *admin.UpdateCompetitionRequest) error {
	startTime, _ := time.Parse(time.RFC3339, request.StartTime)
	deadline, _ := time.Parse(time.RFC3339, request.Deadline)

	_, err := c.CompetitionRepo.UpdateByID(ctx, uint(request.Id), map[string]any{
		"start_time": startTime,
		"deadline":   deadline,
		"category":   int(request.Category),
		"permission": int(request.Permission),
	})
	if err != nil {
		return myErr.BadRequest("", "修改失败！")
	}

	return nil
}

func (c *CompetitionService) Delete(ctx context.Context, request *admin.DeleteCompetitionRequest) error {
	db := c.CompetitionRepo.DB.WithContext(ctx)

	return db.Unscoped().Where("id = ?", request.Id).Delete(&model.Competition{}).Error
}
