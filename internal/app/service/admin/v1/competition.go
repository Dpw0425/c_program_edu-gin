package admin_service

import (
	"c_program_edu-gin/internal/app/schema"
	"c_program_edu-gin/internal/app/storage/model"
	"c_program_edu-gin/internal/app/storage/repo"
	myErr "c_program_edu-gin/pkg/error"
	admin "c_program_edu-gin/schema/genproto/admin/v1/competition"
	"context"
	"gorm.io/gorm"
	"time"
)

var _ ICompetitionService = (*CompetitionService)(nil)

type ICompetitionService interface {
	List(ctx context.Context, request *admin.CompetitionListRequest) ([]*schema.CompetitionItem, int, error)
	Add(ctx context.Context, userId uint, request *admin.AddCompetitionRequest) error
	Update(ctx context.Context, request *admin.UpdateCompetitionRequest) error
	Delete(ctx context.Context, request *admin.DeleteCompetitionRequest) error
	GetQuestionInCpt(ctx context.Context, request *admin.GetQuestionInCptRequest) ([]*schema.QuestionItem, int, error)
	GetQuestionBesideCpt(ctx context.Context, request *admin.GetQuestionBesideCptRequest) ([]*schema.QuestionItem, int, error)
	AddQuestionToCpt(ctx context.Context, request *admin.AddQuestionToCptRequest) error
	ExcludeQuestionFromCpt(ctx context.Context, request *admin.ExcludeQuestionFromCptRequest) error
}

type CompetitionService struct {
	CompetitionRepo *repo.CompetitionRepo
	QuestionRepo    *repo.QuestionRepo
	CptQueRepo      *repo.CptQueRepo
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

func (c *CompetitionService) GetQuestionInCpt(ctx context.Context, request *admin.GetQuestionInCptRequest) ([]*schema.QuestionItem, int, error) {
	db := c.QuestionRepo.DB.WithContext(ctx)
	var items []*schema.QuestionItem
	var questionIds []uint

	err1 := db.Table("cpt_ques").Select("question_id").Where("competition_id = ?", request.Id).Limit(int(request.Number)).Offset(int((request.Page - 1) * request.Number)).Scan(&questionIds).Error
	if err1 != nil {
		return nil, 0, myErr.BadRequest("", "查询失败！")
	}

	var count int64
	if err := db.Table("cpt_ques").Where("competition_id = ?", request.Id).Count(&count).Error; err != nil {
		return nil, 0, myErr.BadRequest("", "查询失败！")
	}

	if len(questionIds) > 0 {
		if err := db.Table("questions").Where("id IN ?", questionIds).Scan(&items).Error; err != nil {
			return nil, 0, myErr.BadRequest("", "查询失败！")
		}
	}

	return items, int(count), nil
}

func (c *CompetitionService) GetQuestionBesideCpt(ctx context.Context, request *admin.GetQuestionBesideCptRequest) ([]*schema.QuestionItem, int, error) {
	db := c.QuestionRepo.DB.WithContext(ctx)
	var items []*schema.QuestionItem
	var questionIds []uint

	if err := db.Table("cpt_ques").Select("question_id").Where("competition_id = ?", request.Id).Scan(&questionIds).Error; err != nil {
		return nil, 0, myErr.BadRequest("", "查询失败！")
	}

	var count int64

	if len(questionIds) != 0 {
		if err := db.Table("questions").Where("id NOT IN ?", questionIds).Count(&count).Error; err != nil {
			return nil, 0, myErr.BadRequest("", "查询失败！")
		}

		if err := db.Table("questions").Where("id NOT IN ?", questionIds).Limit(int(request.Number)).Offset(int((request.Page - 1) * request.Number)).Scan(&items).Error; err != nil {
			return nil, 0, myErr.BadRequest("", "查询失败！")
		}
	} else {
		if err := db.Table("questions").Count(&count).Error; err != nil {
			return nil, 0, myErr.BadRequest("", "查询失败！")
		}

		if err := db.Table("questions").Limit(int(request.Number)).Offset(int((request.Page - 1) * request.Number)).Scan(&items).Error; err != nil {
			return nil, 0, myErr.BadRequest("", "查询失败！")
		}
	}

	return items, int(count), nil
}

func (c *CompetitionService) AddQuestionToCpt(ctx context.Context, request *admin.AddQuestionToCptRequest) error {
	newDB := c.CptQueRepo.DB.Session(&gorm.Session{NewDB: true})
	tx := newDB.WithContext(ctx).Begin()

	for _, questionId := range request.Ids {
		err := tx.Create(&model.CptQue{
			CompetitionID: uint(request.CompetitionId),
			QuestionID:    uint(questionId),
		}).Error
		if err != nil {
			tx.Rollback()
			return myErr.BadRequest("", err.Error())
		}
	}

	tx.Commit()
	return nil
}

func (c *CompetitionService) ExcludeQuestionFromCpt(ctx context.Context, request *admin.ExcludeQuestionFromCptRequest) error {
	db := c.CptQueRepo.DB.WithContext(ctx)

	return db.Unscoped().Where("competition_id = ? AND question_id = ?", request.CompetitionId, request.QuestionId).Delete(&model.CptQue{}).Error
}
