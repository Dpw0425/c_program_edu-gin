package service

import (
	"c_program_edu-gin/internal/app/schema"
	"c_program_edu-gin/internal/app/storage/repo"
	myErr "c_program_edu-gin/pkg/error"
	web "c_program_edu-gin/schema/genproto/web/v1/competition"
	web2 "c_program_edu-gin/schema/genproto/web/v1/question"
	"context"
	"strconv"
	"strings"
)

var _ ICompetitionService = (*CompetitionService)(nil)

type ICompetitionService interface {
	List(ctx context.Context) ([]*schema.WebCompetitionItem, int, error)
	Detail(ctx context.Context, request *web.GetCompetitionDetailRequest) (*schema.WebCompetitionItem, error)
	GetQuestionList(ctx context.Context, request *web2.GetQuestionInCompetitionRequest) ([]*schema.QuestionItem, error)
	Ranking(ctx context.Context, request *web.GetRankingRequest) ([]*schema.Rank, error)
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

func (c *CompetitionService) Detail(ctx context.Context, request *web.GetCompetitionDetailRequest) (*schema.WebCompetitionItem, error) {
	db := c.CompetitionRepo.DB.WithContext(ctx)

	var item *schema.WebCompetitionItem
	err := db.Table("competitions").Where("id = ?", request.Id).Scan(&item).Error
	if err != nil {
		return nil, myErr.NotFound("", err.Error())
	}

	var quantity int64
	db.Table("cpt_ques").Where("competition_id = ?", item.ID).Count(&quantity)
	item.Quantity = int(quantity)

	return item, nil
}

func (c *CompetitionService) GetQuestionList(ctx context.Context, request *web2.GetQuestionInCompetitionRequest) ([]*schema.QuestionItem, error) {
	db := c.CompetitionRepo.DB.WithContext(ctx)
	var items []*schema.QuestionItem
	var questionIds []uint

	err1 := db.Table("cpt_ques").Select("question_id").Where("competition_id = ?", request.Id).Scan(&questionIds).Error
	if err1 != nil {
		return nil, myErr.BadRequest("", "查询失败！")
	}

	if len(questionIds) > 0 {
		if err := db.Table("questions").Where("id IN ?", questionIds).Scan(&items).Error; err != nil {
			return nil, myErr.BadRequest("", "查询失败！")
		}
	}

	return items, nil
}

func (c *CompetitionService) Ranking(ctx context.Context, request *web.GetRankingRequest) ([]*schema.Rank, error) {
	db := c.CompetitionRepo.DB.WithContext(ctx)

	var contestant string
	db.Table("competitions").Select("contestant").Where("id = ?", request.Id).Scan(&contestant)
	users := strings.Split(contestant, ",")
	userIds := make([]int64, 0, len(users))
	for _, strID := range users {
		if id, err := strconv.ParseInt(strID, 10, 64); err == nil {
			userIds = append(userIds, id)
		}
	}

	var questionIds []uint
	db.Table("cpt_ques").Select("question_id").Where("competition_id = ?", request.Id).Scan(&questionIds)

	var topUsers []*schema.Rank
	subQuery := db.Table("users").
		Select("users.user_id, users.user_name").
		Where("users.user_id IN (?)", userIds)

	db.Table("(?) as u", subQuery).
		Select("u.user_name, "+
			"COUNT(DISTINCT CASE WHEN uq.status = 'Accepted' THEN uq.question_id END) AS score, "+
			"COALESCE(SUM(uq.commit_times), 0) AS total_commit").
		Joins("LEFT JOIN user_ques uq ON u.user_id = uq.user_id AND uq.question_id IN (?)", questionIds).
		Group("u.user_id, u.user_name").
		Order("score DESC, total_commit ASC").
		Limit(10).Scan(&topUsers)

	return topUsers, nil
}
