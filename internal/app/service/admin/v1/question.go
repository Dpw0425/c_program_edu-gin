package admin_service

import (
	"c_program_edu-gin/internal/app/schema"
	"c_program_edu-gin/internal/app/storage/model"
	"c_program_edu-gin/internal/app/storage/repo"
	myErr "c_program_edu-gin/pkg/error"
	admin "c_program_edu-gin/schema/genproto/admin/v1/question"
	"context"
)

var _ IQuestionService = (*QuestionService)(nil)

type IQuestionService interface {
	Add(ctx context.Context, question *model.Question) error
	List(ctx context.Context, request *admin.QuestionListRequest) ([]*schema.QuestionItem, error)
}

type QuestionService struct {
	QuestionRepo *repo.QuestionRepo
}

func (q *QuestionService) Add(ctx context.Context, question *model.Question) error {

	return nil
}

func (q *QuestionService) List(ctx context.Context, request *admin.QuestionListRequest) ([]*schema.QuestionItem, error) {
	db := q.QuestionRepo.DB.WithContext(ctx)
	var items []*schema.QuestionItem

	err := db.Table("questions").Select("id", "title", "tag", "degree", "owner_id", "passing_rate").Where("title LIKE ?", "%"+request.Search+"%").Limit(int(request.Number)).Offset(int((request.Page - 1) * request.Number)).Scan(&items).Error
	if err != nil {
		return nil, myErr.NotFound("", err.Error())
	}

	return items, nil
}
