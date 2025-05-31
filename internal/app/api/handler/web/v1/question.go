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

	list, count, err := q.QuestionService.List(ctx.Ctx(), params)
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
		Total:        int32(count),
	}, "查询成功！")
	return nil
}

func (q *Question) Detail(ctx *ctx.Context) error {
	params := &web.GetQuestionDetailRequest{}
	if err := ctx.Context.ShouldBindQuery(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	item, err := q.QuestionService.Detail(ctx.Ctx(), params)
	if err != nil {
		return err
	}

	result := &web.GetQuestionDetailResponse_QuestionItem{
		Id:          int32(item.ID),
		Title:       item.Title,
		Tag:         strings.Split(item.Tag, ","),
		Degree:      int32(item.Degree),
		PassingRate: item.PassingRate,
		OwnerId:     int32(item.OwnerID),
		Content:     item.Content,
	}

	response.NorResponse(ctx.Context, &web.GetQuestionDetailResponse{QuestionItem: result}, "查询成功！")
	return nil
}

func (q *Question) TestData(ctx *ctx.Context) error {
	params := &web.GetTestDataListRequest{}
	if err := ctx.Context.ShouldBindQuery(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	list, err := q.QuestionService.TestData(ctx.Ctx(), params)
	if err != nil {
		return err
	}

	items := make([]*web.GetTestDataListResponse_TestData, 0, len(list))
	for _, item := range list {
		items = append(items, &web.GetTestDataListResponse_TestData{
			Input:  item.Input,
			Output: item.Output,
		})
	}

	response.NorResponse(ctx.Context, &web.GetTestDataListResponse{TestData: items}, "查询成功！")
	return nil
}

func (q *Question) CommitAnswer(ctx *ctx.Context) error {
	params := &web.CommitAnswerRequest{}
	if err := ctx.Context.ShouldBind(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	result, err := q.QuestionService.CommitAnswer(ctx.Ctx(), params, ctx.UserID())
	if err != nil {
		return err
	}

	response.NorResponse(ctx.Context, &web.CommitAnswerResponse{Result: result}, "提交成功！")
	return nil
}
