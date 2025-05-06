package service

import (
	"c_program_edu-gin/internal/app/schema"
	"c_program_edu-gin/internal/app/storage/repo"
	myErr "c_program_edu-gin/pkg/error"
	web "c_program_edu-gin/schema/genproto/web/v1/question"
	"context"
)

var _ IQuestionService = (*QuestionService)(nil)

type IQuestionService interface {
	List(ctx context.Context, request *web.GetQuestionListRequest) ([]*schema.QuestionItem, error)
}

type QuestionService struct {
	QuestionRepo *repo.QuestionRepo
}

func (q *QuestionService) List(ctx context.Context, request *web.GetQuestionListRequest) ([]*schema.QuestionItem, error) {
	db := q.QuestionRepo.DB.WithContext(ctx)
	var items []*schema.QuestionItem

	if request.Search == "热点题目" {
		err := db.Table("questions").Select("id", "title", "passing_rate").Order("passing_rate ASC").Limit(5).Scan(&items).Error
		if err != nil {
			return nil, myErr.NotFound("", err.Error())
		}

		return items, nil
	}

	err := db.Table("questions").Select("id", "title", "tag", "degree", "owner_id", "passing_rate").Where("title LIKE ?", "%"+request.Search+"%").Limit(int(request.Number)).Offset(int((request.Page - 1) * request.Number)).Scan(&items).Error
	if err != nil {
		return nil, myErr.NotFound("", err.Error())
	}

	return items, nil
}
