package service

import (
	"c_program_edu-gin/internal/app/schema"
	"c_program_edu-gin/internal/app/storage/model"
	"c_program_edu-gin/internal/app/storage/repo"
	myErr "c_program_edu-gin/pkg/error"
	web "c_program_edu-gin/schema/genproto/web/v1/question"
	"c_program_edu-gin/utils/judge"
	"context"
	"gorm.io/gorm"
)

var _ IQuestionService = (*QuestionService)(nil)

type IQuestionService interface {
	List(ctx context.Context, request *web.GetQuestionListRequest) ([]*schema.QuestionItem, int, error)
	Detail(ctx context.Context, request *web.GetQuestionDetailRequest) (*schema.QuestionItem, error)
	TestData(ctx context.Context, request *web.GetTestDataListRequest) ([]*schema.TestDataItem, error)
	CommitAnswer(ctx context.Context, request *web.CommitAnswerRequest, userId int64) (string, error)
}

type QuestionService struct {
	QuestionRepo *repo.QuestionRepo
}

func (q *QuestionService) List(ctx context.Context, request *web.GetQuestionListRequest) ([]*schema.QuestionItem, int, error) {
	db := q.QuestionRepo.DB.WithContext(ctx)
	var items []*schema.QuestionItem

	if request.Search == "热点题目" {
		err := db.Table("questions").Select("id", "title", "passing_rate").Order("passing_rate ASC").Limit(5).Scan(&items).Error
		if err != nil {
			return nil, 0, myErr.NotFound("", err.Error())
		}

		return items, 0, nil
	}

	err := db.Table("questions").Where("title LIKE ? AND status = 0", "%"+request.Search+"%").Limit(int(request.Number)).Offset(int((request.Page - 1) * request.Number)).Scan(&items).Error
	if err != nil {
		return nil, 0, myErr.NotFound("", err.Error())
	}

	var count int64
	db.Table("questions").Count(&count)

	return items, int(count), nil
}

func (q *QuestionService) Detail(ctx context.Context, request *web.GetQuestionDetailRequest) (*schema.QuestionItem, error) {
	db := q.QuestionRepo.DB.WithContext(ctx)
	var item *schema.QuestionItem

	err := db.Table("questions").Where("id = ?", request.Id).Scan(&item).Error
	if err != nil {
		return nil, myErr.NotFound("", err.Error())
	}

	return item, nil
}

func (q *QuestionService) TestData(ctx context.Context, request *web.GetTestDataListRequest) ([]*schema.TestDataItem, error) {
	db := q.QuestionRepo.DB.WithContext(ctx)
	var item []*schema.TestDataItem

	err := db.Table("test_data").Where("question_id = ?", request.Id).Scan(&item).Error
	if err != nil {
		return nil, myErr.NotFound("", err.Error())
	}

	return item, nil
}

func (q *QuestionService) CommitAnswer(ctx context.Context, request *web.CommitAnswerRequest, userId int64) (string, error) {
	newDB := q.QuestionRepo.DB.Session(&gorm.Session{NewDB: true})
	tx := newDB.WithContext(ctx).Begin()

	var eq *model.Question
	if err := tx.Table("questions").Where("id = ?", request.QuestionId).Scan(&eq).Error; err != nil {
		tx.Rollback()
		return "", myErr.BadRequest("", "题目不存在")
	}

	var result string
	var testDataList = make([]*model.TestData, 0)
	if err := tx.Table("test_data").Where("question_id = ?", request.QuestionId).Scan(&testDataList).Error; err != nil {
		tx.Rollback()
		return "", myErr.BadRequest("", "提交失败！")
	}

	for _, value := range testDataList {
		if flag := judge.Judge(request.Answer, value.Input, value.Output); flag != "Accepted" {
			result = flag
			break
		}
	}

	if result == "" {
		result = "Accepted"
	}

	if tx.Table("user_ques").Where("user_id = ? AND question_id = ?", userId, request.QuestionId).RowsAffected == 0 {
		if err := tx.Table("user_ques").Create(&model.UserQue{
			UserID:      userId,
			QuestionID:  uint(request.QuestionId),
			CommitTimes: 1,
			Status:      result,
		}).Error; err != nil {
			tx.Rollback()
			return "", myErr.BadRequest("", "提交失败！")
		}
	} else {
		if err := tx.Table("user_ques").Where("user_id = ? AND question_id = ?", userId, request.QuestionId).Updates(map[string]any{
			"commit_times": gorm.Expr("commit_times + ?", 1),
			"status":       result,
		}).Error; err != nil {
			tx.Rollback()
			return "", myErr.BadRequest("", "提交失败！")
		}
	}

	tx.Commit()
	return result, nil
}
