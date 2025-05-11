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
	params := &admin.AddQuestionRequest{}
	if err := ctx.Context.ShouldBind(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！%v", err)
	}

	if result := q.QuestionService.Add(ctx.Ctx(), params, ctx.UserID()); result != nil {
		return result
	}

	response.NorResponse(ctx.Context, &admin.AddQuestionResponse{}, "发布成功！")
	return nil
}

func (q *Question) List(ctx *ctx.Context) error {
	params := &admin.QuestionListRequest{}
	if err := ctx.Context.ShouldBindQuery(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	list, count, err := q.QuestionService.List(ctx.Ctx(), params)
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
			Content:     item.Content,
			Answer:      item.Answer,
			Status:      int32(item.Status),
		})
	}

	response.NorResponse(ctx.Context, &admin.QuestionListResponse{
		QuestionList: items,
		Total:        int32(count),
	}, "查询成功！")
	return nil
}

func (q *Question) Update(ctx *ctx.Context) error {
	params := &admin.UpdateQuestionRequest{}
	if err := ctx.Context.ShouldBind(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	if result := q.QuestionService.Update(ctx.Ctx(), params); result != nil {
		return result
	}

	response.NorResponse(ctx.Context, &admin.UpdateQuestionResponse{}, "修改成功！")
	return nil
}

func (q *Question) Delete(ctx *ctx.Context) error {
	params := &admin.DeleteQuestionRequest{}
	if err := ctx.Context.ShouldBindQuery(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	if result := q.QuestionService.Delete(ctx.Ctx(), params); result != nil {
		return result
	}

	response.NorResponse(ctx.Context, &admin.DeleteQuestionResponse{}, "删除成功！")
	return nil
}

func (q *Question) AddTestData(ctx *ctx.Context) error {
	params := &admin.AddTestDataRequest{}
	if err := ctx.Context.ShouldBind(&params); err != nil {
		return myErr.BadRequest("wrong_parameter", "请求参数错误！")
	}

	if result := q.QuestionService.AddTestData(ctx.Ctx(), params); result != nil {
		return result
	}

	response.NorResponse(ctx.Context, &admin.AddTestDataResponse{}, "添加成功！")
	return nil
}

func (q *Question) GetTestData(ctx *ctx.Context) error {
	params := &admin.GetTestDataRequest{}
	if err := ctx.Context.ShouldBindQuery(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	list, err := q.QuestionService.GetTestData(ctx.Ctx(), params)
	if err != nil {
		return err
	}

	items := make([]*admin.GetTestDataResponse_TestDataItem, 0, len(list))
	for _, item := range list {
		items = append(items, &admin.GetTestDataResponse_TestDataItem{
			Id:         int32(item.ID),
			Input:      item.Input,
			Output:     item.Output,
			QuestionId: int32(item.QuestionID),
		})
	}

	response.NorResponse(ctx.Context, &admin.GetTestDataResponse{TestDataList: items}, "查询成功！")
	return nil
}

func (q *Question) UpdateTestData(ctx *ctx.Context) error {
	params := &admin.UpdateTestDataRequest{}
	if err := ctx.Context.ShouldBind(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	if result := q.QuestionService.UpdateTestData(ctx.Ctx(), params); result != nil {
		return result
	}

	response.NorResponse(ctx.Context, &admin.UpdateTestDataResponse{}, "修改成功！")
	return nil
}

func (q *Question) DeleteTestData(ctx *ctx.Context) error {
	params := &admin.DeleteTestDataRequest{}
	if err := ctx.Context.ShouldBindQuery(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	if result := q.QuestionService.DeleteTestData(ctx.Ctx(), params); result != nil {
		return result
	}

	response.NorResponse(ctx.Context, &admin.DeleteTestDataResponse{}, "删除成功！")
	return nil
}
