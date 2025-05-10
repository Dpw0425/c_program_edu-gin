package admin_service

import (
	"c_program_edu-gin/internal/app/schema"
	"c_program_edu-gin/internal/app/storage/model"
	"c_program_edu-gin/internal/app/storage/repo"
	myErr "c_program_edu-gin/pkg/error"
	admin "c_program_edu-gin/schema/genproto/admin/v1/question"
	"context"
	"strings"
)

var _ IQuestionService = (*QuestionService)(nil)

type IQuestionService interface {
	Add(ctx context.Context, request *admin.AddQuestionRequest, userId int64) error
	List(ctx context.Context, request *admin.QuestionListRequest) ([]*schema.QuestionItem, int, error)
	Update(ctx context.Context, request *admin.UpdateQuestionRequest) error
	Delete(ctx context.Context, request *admin.DeleteQuestionRequest) error
}

type QuestionService struct {
	QuestionRepo *repo.QuestionRepo
}

func (q *QuestionService) Add(ctx context.Context, request *admin.AddQuestionRequest, userId int64) error {
	if q.QuestionRepo.IsExist(ctx, request.Title) {
		return myErr.BadRequest("", "题目已存在！")
	}

	return q.QuestionRepo.Create(ctx, &model.Question{
		Title:   request.Title,
		Content: request.Content,
		Answer:  request.Answer,
		Tag:     strings.Join(request.Tag, ","),
		Degree:  int(request.Degree),
		OwnerID: userId,
		Status:  int(request.Status),
	})
}

func (q *QuestionService) List(ctx context.Context, request *admin.QuestionListRequest) ([]*schema.QuestionItem, int, error) {
	db := q.QuestionRepo.DB.WithContext(ctx)
	var items []*schema.QuestionItem

	err := db.Table("questions").Where("title LIKE ?", "%"+request.Search+"%").Limit(int(request.Number)).Offset(int((request.Page - 1) * request.Number)).Scan(&items).Error
	if err != nil {
		return nil, 0, myErr.NotFound("", err.Error())
	}

	var count int64
	db.Table("questions").Count(&count)

	return items, int(count), nil
}

func (q *QuestionService) Update(ctx context.Context, request *admin.UpdateQuestionRequest) error {
	_, err := q.QuestionRepo.UpdateByID(ctx, uint(request.Id), map[string]any{
		"title":   request.Title,
		"content": request.Content,
		"answer":  request.Answer,
		"tag":     strings.Join(request.Tag, ","),
		"degree":  request.Degree,
		"status":  request.Status,
	})
	if err != nil {
		return myErr.BadRequest("", "修改失败！")
	}

	return nil
}

func (q *QuestionService) Delete(ctx context.Context, request *admin.DeleteQuestionRequest) error {
	db := q.QuestionRepo.DB.WithContext(ctx)

	return db.Unscoped().Where("id = ?", request.Id).Delete(&model.Question{}).Error
}
