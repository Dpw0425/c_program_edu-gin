package v1

import (
	service "c_program_edu-gin/internal/app/service/web/v1"
	ctx "c_program_edu-gin/internal/pkg/context"
	myErr "c_program_edu-gin/pkg/error"
	"c_program_edu-gin/pkg/response"
	web "c_program_edu-gin/schema/genproto/web/v1/question"
	"strings"
)

type Question struct {
	QuestionService service.QuestionService
}

func (q *Question) List(ctx *ctx.Context) error {
	params := &web.GetQuestionListRequest{}
	if err := ctx.Context.ShouldBindQuery(params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	list, err := q.QuestionService.List(ctx.Ctx(), params)
	if err != nil {
		return err
	}

	items := make([]*web.GetQuestionListResponse_QuestionItem, 0, len(list))
	for _, item := range list {
		items = append(items, &web.GetQuestionListResponse_QuestionItem{
			Id:          int32(item.ID),
			Title:       item.Title,
			Tag:         strings.Split(item.Tag, ","),
			Degree:      int32(item.Degree),
			PassingRate: item.PassingRate,
			OwnerId:     int32(item.OwnerID),
		})
	}

	response.NorResponse(ctx.Context, &web.GetQuestionListResponse{
		QuestionList: items,
		Total:        int32(len(items)),
	}, "查询成功！")
	return nil
}
