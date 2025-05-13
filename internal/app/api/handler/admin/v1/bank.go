package v1

import (
	adminservice "c_program_edu-gin/internal/app/service/admin/v1"
	ctx "c_program_edu-gin/internal/pkg/context"
	myErr "c_program_edu-gin/pkg/error"
	"c_program_edu-gin/pkg/response"
	admin "c_program_edu-gin/schema/genproto/admin/v1/bank"
	"strings"
)

type Bank struct {
	BankService adminservice.BankService
}

func (b *Bank) List(ctx *ctx.Context) error {
	params := &admin.BankListRequest{}
	if err := ctx.Context.ShouldBindQuery(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	list, count, err := b.BankService.List(ctx.Ctx(), params)
	if err != nil {
		return err
	}

	items := make([]*admin.BankListResponse_BankItem, 0, len(list))
	for _, item := range list {
		items = append(items, &admin.BankListResponse_BankItem{
			Id:           int32(item.ID),
			Name:         item.Name,
			Content:      item.Content,
			OpenGrade:    strings.Split(item.OpenGrade, ","),
			Participants: int32(item.Participants),
			Completed:    int32(item.Completed),
			Count:        int32(item.Count),
		})
	}

	response.NorResponse(ctx.Context, &admin.BankListResponse{
		BankList: items,
		Total:    int32(count),
	}, "查询成功！")
	return nil
}

func (b *Bank) Add(ctx *ctx.Context) error {
	params := &admin.AddBankRequest{}
	if err := ctx.Context.ShouldBind(&params); err != nil {
		return err
	}

	if result := b.BankService.Add(ctx.Ctx(), params); result != nil {
		return result
	}

	response.NorResponse(ctx.Context, &admin.AddBankResponse{}, "创建成功！")
	return nil
}

func (b *Bank) Update(ctx *ctx.Context) error {
	params := &admin.UpdateBankRequest{}
	if err := ctx.Context.ShouldBind(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	if result := b.BankService.Update(ctx.Ctx(), params); result != nil {
		return result
	}

	response.NorResponse(ctx.Context, &admin.UpdateBankResponse{}, "修改成功！")
	return nil
}

func (b *Bank) Delete(ctx *ctx.Context) error {
	params := &admin.DeleteBankRequest{}
	if err := ctx.Context.ShouldBindQuery(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	if result := b.BankService.Delete(ctx.Ctx(), params); result != nil {
		return result
	}

	response.NorResponse(ctx.Context, &admin.DeleteBankResponse{}, "删除成功！")
	return nil
}

func (b *Bank) GetQuestionInBank(ctx *ctx.Context) error {
	params := &admin.GetQuestionInBankRequest{}
	if err := ctx.Context.ShouldBindQuery(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	list, count, err := b.BankService.GetQuestionInBank(ctx.Ctx(), params)
	if err != nil {
		return err
	}

	items := make([]*admin.GetQuestionInBankResponse_QuestionItem, 0, len(list))
	for _, item := range list {
		items = append(items, &admin.GetQuestionInBankResponse_QuestionItem{
			Id:          int32(item.ID),
			Title:       item.Title,
			Tag:         strings.Split(item.Tag, ","),
			Degree:      int32(item.Degree),
			PassingRate: item.PassingRate,
			OwnerId:     int32(item.OwnerID),
			Content:     item.Content,
			Answer:      item.Answer,
			Status:      int32(item.Status),
		})
	}

	response.NorResponse(ctx.Context, &admin.GetQuestionInBankResponse{
		QuestionList: items,
		Total:        int32(count),
	}, "查询成功！")
	return nil
}

func (b *Bank) GetQuestionBesideBank(ctx *ctx.Context) error {
	params := &admin.GetQuestionBesideBankRequest{}
	if err := ctx.Context.ShouldBindQuery(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	list, count, err := b.BankService.GetQuestionBesideBank(ctx.Ctx(), params)
	if err != nil {
		return err
	}

	items := make([]*admin.GetQuestionBesideBankResponse_QuestionItem, 0, len(list))
	for _, item := range list {
		items = append(items, &admin.GetQuestionBesideBankResponse_QuestionItem{
			Id:          int32(item.ID),
			Title:       item.Title,
			Tag:         strings.Split(item.Tag, ","),
			Degree:      int32(item.Degree),
			PassingRate: item.PassingRate,
			OwnerId:     int32(item.OwnerID),
			Content:     item.Content,
			Answer:      item.Answer,
			Status:      int32(item.Status),
		})
	}

	response.NorResponse(ctx.Context, &admin.GetQuestionBesideBankResponse{
		QuestionList: items,
		Total:        int32(count),
	}, "查询成功！")
	return nil
}

func (b *Bank) ExcludeQuestion(ctx *ctx.Context) error {
	params := &admin.ExcludeQuestionRequest{}
	if err := ctx.Context.ShouldBindQuery(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	if result := b.BankService.ExcludeQuestion(ctx.Ctx(), params); result != nil {
		return result
	}

	response.NorResponse(ctx.Context, &admin.ExcludeQuestionResponse{}, "删除成功！")
	return nil
}
