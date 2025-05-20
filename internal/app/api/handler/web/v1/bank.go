package v1

import (
	service "c_program_edu-gin/internal/app/service/web/v1"
	ctx "c_program_edu-gin/internal/pkg/context"
	myErr "c_program_edu-gin/pkg/error"
	"c_program_edu-gin/pkg/response"
	web "c_program_edu-gin/schema/genproto/web/v1/bank"
	"strings"
)

type Bank struct {
	BankService service.BankService
}

func (b *Bank) List(ctx *ctx.Context) error {
	params := &web.GetAllBankRequest{}
	if err1 := ctx.Context.ShouldBindQuery(&params); err1 != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	list, err := b.BankService.List(ctx.Ctx(), params)
	if err != nil {
		return err
	}

	items := make([]*web.GetAllBankResponse_BankItem, 0, len(list))
	for _, item := range list {
		items = append(items, &web.GetAllBankResponse_BankItem{
			Id:           int32(item.ID),
			Name:         item.Name,
			Content:      item.Content,
			OpenGrade:    strings.Split(item.OpenGrade, ","),
			Participants: int32(item.Participants),
			Completed:    int32(item.Completed),
			Count:        int32(item.Count),
		})
	}

	response.NorResponse(ctx.Context, &web.GetAllBankResponse{BankList: items}, "查询成功！")
	return nil
}
