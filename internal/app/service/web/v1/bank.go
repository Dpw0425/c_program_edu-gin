package service

import (
	"c_program_edu-gin/internal/app/schema"
	"c_program_edu-gin/internal/app/storage/repo"
	myErr "c_program_edu-gin/pkg/error"
	web "c_program_edu-gin/schema/genproto/web/v1/bank"
	"context"
)

var _ IBankService = (*BankService)(nil)

type IBankService interface {
	List(ctx context.Context, request *web.GetAllBankRequest) ([]*schema.BankItem, error)
}

type BankService struct {
	BankRepo *repo.BankRepo
}

func (b *BankService) List(ctx context.Context, request *web.GetAllBankRequest) ([]*schema.BankItem, error) {
	db := b.BankRepo.DB.WithContext(ctx)
	var items []*schema.BankItem

	err := db.Table("question_banks").Where("FIND_IN_SET(?,open_grade)", request.Grade).Scan(&items).Error
	if err != nil {
		return nil, myErr.NotFound("", err.Error())
	}

	for _, item := range items {
		var count int64
		if err := db.Table("bank_ques").Where("bank_id = ?", item.ID).Count(&count).Error; err == nil {
			item.Count = int(count)
		} else {
			return nil, myErr.NotFound("", err.Error())
		}
	}

	return items, nil
}
