package admin_service

import (
	"c_program_edu-gin/internal/app/schema"
	"c_program_edu-gin/internal/app/storage/model"
	"c_program_edu-gin/internal/app/storage/repo"
	myErr "c_program_edu-gin/pkg/error"
	admin "c_program_edu-gin/schema/genproto/admin/v1/bank"
	"context"
	"gorm.io/gorm"
	"strings"
)

var _ IBankService = (*BankService)(nil)

type IBankService interface {
	List(ctx context.Context, request *admin.BankListRequest) ([]*schema.BankItem, int, error)
	Add(ctx context.Context, request *admin.AddBankRequest) error
	Update(ctx context.Context, request *admin.UpdateBankRequest) error
	Delete(ctx context.Context, request *admin.DeleteBankRequest) error
	GetQuestionInBank(ctx context.Context, request *admin.GetQuestionInBankRequest) ([]*schema.QuestionItem, int, error)
	GetQuestionBesideBank(ctx context.Context, request *admin.GetQuestionBesideBankRequest) ([]*schema.QuestionItem, int, error)
	ExcludeQuestion(ctx context.Context, request *admin.ExcludeQuestionRequest) error
}

type BankService struct {
	QuestionRepo *repo.QuestionRepo
	BankRepo     *repo.BankRepo
	BankQueRepo  *repo.BankQueRepo
}

func (b *BankService) List(ctx context.Context, request *admin.BankListRequest) ([]*schema.BankItem, int, error) {
	db := b.BankRepo.DB.WithContext(ctx)
	var items []*schema.BankItem

	err := db.Table("question_banks").Where("name LIKE ?", "%"+request.Search+"%").Limit(int(request.Number)).Offset(int((request.Page - 1) * request.Number)).Scan(&items).Error
	if err != nil {
		return nil, 0, myErr.NotFound("", err.Error())
	}

	var count int64
	db.Table("question_banks").Count(&count)

	return items, int(count), nil
}

func (b *BankService) Add(ctx context.Context, request *admin.AddBankRequest) error {
	newDB := b.BankRepo.DB.Session(&gorm.Session{NewDB: true})
	tx := newDB.WithContext(ctx).Begin()

	if b.BankRepo.IsExist(ctx, request.Name) {
		return myErr.BadRequest("", "题目已存在！")
	}

	if err := tx.Create(&model.QuestionBank{
		Name:      request.Name,
		Content:   request.Content,
		OpenGrade: strings.Join(request.OpenGrade, ","),
	}).Error; err != nil {
		tx.Rollback()
		return myErr.BadRequest("", err.Error())
	}

	if request.Questions == nil {
		tx.Commit()
		return nil
	} else {
		newDB1 := b.BankQueRepo.DB.Session(&gorm.Session{NewDB: true})
		tx1 := newDB1.WithContext(ctx).Begin()
		var id uint
		err := tx.Table("question_banks").Select("id").Where("name = ?", request.Name).Scan(&id).Error
		if err != nil {
			tx.Rollback()
			return myErr.BadRequest("", err.Error())
		}

		for _, questionId := range request.Questions {
			err := tx1.Create(&model.BankQue{
				BankID:     id,
				QuestionID: uint(questionId),
			}).Error
			if err != nil {
				tx.Rollback()
				tx1.Rollback()
				return myErr.BadRequest("", err.Error())
			}
		}

		tx.Commit()
		tx1.Commit()
		return nil
	}
}

func (b *BankService) Update(ctx context.Context, request *admin.UpdateBankRequest) error {
	newDB := b.BankQueRepo.DB.Session(&gorm.Session{NewDB: true})
	tx := newDB.WithContext(ctx).Begin()

	_, err := b.BankRepo.UpdateByID(ctx, uint(request.Id), map[string]any{
		"content":    request.Content,
		"open_grade": strings.Join(request.OpenGrade, ","),
	})
	if err != nil {
		tx.Rollback()
		return myErr.BadRequest("", "修改失败！")
	}

	if request.QuestionId != nil {
		for _, id := range request.QuestionId {
			if err := tx.Create(&model.BankQue{
				BankID:     uint(request.Id),
				QuestionID: uint(id),
			}).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	tx.Commit()

	return nil
}

func (b *BankService) Delete(ctx context.Context, request *admin.DeleteBankRequest) error {
	newDB := b.BankRepo.DB.Session(&gorm.Session{NewDB: true})
	tx := newDB.WithContext(ctx).Begin()
	newDB1 := b.BankQueRepo.DB.Session(&gorm.Session{NewDB: true})
	tx1 := newDB1.WithContext(ctx).Begin()

	if err := tx.Unscoped().Where("id = ?", request.Id).Delete(&model.QuestionBank{}).Error; err != nil {
		tx.Rollback()
		tx1.Rollback()
		return myErr.BadRequest("", err.Error())
	}

	if err := tx.Unscoped().Where("question_id = ?", request.Id).Delete(&model.BankQue{}).Error; err != nil {
		tx.Rollback()
		tx1.Rollback()
		return myErr.BadRequest("", err.Error())
	}

	tx.Commit()
	tx1.Commit()
	return nil
}

func (b *BankService) GetQuestionInBank(ctx context.Context, request *admin.GetQuestionInBankRequest) ([]*schema.QuestionItem, int, error) {
	db := b.QuestionRepo.DB.WithContext(ctx)
	var items []*schema.QuestionItem
	var questionIds []uint

	err1 := db.Table("bank_ques").Select("question_id").Where("bank_id = ?", request.Id).Limit(int(request.Number)).Offset(int((request.Page - 1) * request.Number)).Scan(&questionIds).Error
	if err1 != nil {
		return nil, 0, myErr.BadRequest("", "查询失败！")
	}

	var count int64
	if err := db.Table("bank_ques").Where("bank_id = ?", request.Id).Count(&count).Error; err != nil {
		return nil, 0, myErr.BadRequest("", "查询失败！")
	}

	if len(questionIds) > 0 {
		if err := db.Table("questions").Where("id IN ?", questionIds).Scan(&items).Error; err != nil {
			return nil, 0, myErr.BadRequest("", "查询失败！")
		}
	}

	return items, int(count), nil
}

func (b *BankService) GetQuestionBesideBank(ctx context.Context, request *admin.GetQuestionBesideBankRequest) ([]*schema.QuestionItem, int, error) {
	db := b.QuestionRepo.DB.WithContext(ctx)
	var items []*schema.QuestionItem
	var questionIds []uint

	if err := db.Table("bank_ques").Select("question_id").Where("bank_id = ?", request.Id).Scan(&questionIds).Error; err != nil {
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

func (b *BankService) ExcludeQuestion(ctx context.Context, request *admin.ExcludeQuestionRequest) error {
	db := b.BankQueRepo.DB.WithContext(ctx)

	return db.Unscoped().Where("bank_id = ? AND question_id = ?", request.BankId, request.QuestionId).Delete(&model.BankQue{}).Error
}
