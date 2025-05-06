package v1

import (
	adminservice "c_program_edu-gin/internal/app/service/admin/v1"
	ctx "c_program_edu-gin/internal/pkg/context"
	myErr "c_program_edu-gin/pkg/error"
	"c_program_edu-gin/pkg/response"
	admin "c_program_edu-gin/schema/genproto/admin/v1/question"
	"strings"
)

type Question struct {
	QuestionService adminservice.QuestionService
}

func (q *Question) Add(ctx *ctx.Context) error {
	return nil
}

func (q *Question) List(ctx *ctx.Context) error {
	params := &admin.QuestionListRequest{}
	if err := ctx.Context.ShouldBindQuery(params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	list, err := q.QuestionService.List(ctx.Ctx(), params)
	if err != nil {
		return err
	}

	items := make([]*admin.QuestionListResponse_QuestionItem, 0, len(list))
	for _, item := range list {
		items = append(items, &admin.QuestionListResponse_QuestionItem{
			Id:          int32(item.ID),
			Title:       item.Title,
			Tag:         strings.Split(item.Tag, ","),
			Degree:      int32(item.Degree),
			PassingRate: item.PassingRate,
			OwnerId:     int32(item.OwnerID),
		})
	}

	response.NorResponse(ctx.Context, &admin.QuestionListResponse{
		QuestionList: items,
		Total:        int32(len(items)),
	}, "查询成功！")
	return nil
}
