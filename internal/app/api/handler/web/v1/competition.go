package v1

import (
	service "c_program_edu-gin/internal/app/service/web/v1"
	ctx "c_program_edu-gin/internal/pkg/context"
	myErr "c_program_edu-gin/pkg/error"
	"c_program_edu-gin/pkg/response"
	web "c_program_edu-gin/schema/genproto/web/v1/competition"
	web2 "c_program_edu-gin/schema/genproto/web/v1/question"
	"strings"
)

type Competition struct {
	CompetitionService *service.CompetitionService
}

func (c *Competition) List(ctx *ctx.Context) error {
	list, count, err := c.CompetitionService.List(ctx.Ctx())
	if err != nil {
		return err
	}

	var items = make([]*web.GetCompetitionListResponse_CompetitionItem, 0, len(list))
	for _, item := range list {
		items = append(items, &web.GetCompetitionListResponse_CompetitionItem{
			Id:         int32(item.ID),
			Name:       item.Name,
			Contestant: strings.Split(item.Contestant, ","),
			OwnerId:    int32(item.OwnerID),
			StartTime:  item.StartTime.Format("2006-01-02 15:04:05"),
			Deadline:   item.Deadline.Format("2006-01-02 15:04:05"),
			Status:     int32(item.Status),
			Category:   int32(item.Category),
			Permission: int32(item.Permission),
			Quantity:   int32(item.Quantity),
		})
	}

	response.NorResponse(ctx.Context, &web.GetCompetitionListResponse{
		CompetitionList: items,
		Total:           int32(count),
	}, "查询成功！")
	return nil
}

func (c *Competition) Detail(ctx *ctx.Context) error {
	params := &web.GetCompetitionDetailRequest{}
	if err := ctx.Context.ShouldBindQuery(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	result, err := c.CompetitionService.Detail(ctx.Ctx(), params)
	if err != nil {
		return err
	}

	response.NorResponse(ctx.Context, &web.GetCompetitionDetailResponse{CompetitionItem: &web.GetCompetitionDetailResponse_Competition{
		Id:         int32(result.ID),
		Name:       result.Name,
		Contestant: strings.Split(result.Contestant, ","),
		OwnerId:    int32(result.OwnerID),
		StartTime:  result.StartTime.Format("2006-01-02 15:04:05"),
		Deadline:   result.Deadline.Format("2006-01-02 15:04:05"),
		Status:     int32(result.Status),
		Category:   int32(result.Category),
		Permission: int32(result.Permission),
		Quantity:   int32(result.Quantity),
	}}, "查询成功！")
	return nil
}

func (c *Competition) GetQuestionList(ctx *ctx.Context) error {
	params := &web2.GetQuestionInCompetitionRequest{}
	if err := ctx.Context.ShouldBindQuery(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	list, err := c.CompetitionService.GetQuestionList(ctx.Ctx(), params)
	if err != nil {
		return err
	}

	var items = make([]*web2.GetQuestionInCompetitionResponse_QuestionItem, 0, len(list))
	for _, item := range list {
		items = append(items, &web2.GetQuestionInCompetitionResponse_QuestionItem{
			Id:          int32(item.ID),
			Title:       item.Title,
			Tag:         strings.Split(item.Tag, ","),
			Degree:      int32(item.Degree),
			PassingRate: item.PassingRate,
			OwnerId:     int32(item.OwnerID),
		})
	}

	response.NorResponse(ctx.Context, &web2.GetQuestionInCompetitionResponse{
		QuestionList: items,
	}, "查询成功！")
	return nil
}

func (c *Competition) Ranking(ctx *ctx.Context) error {
	params := &web.GetRankingRequest{}
	if err := ctx.Context.ShouldBindQuery(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	list, err := c.CompetitionService.Ranking(ctx.Ctx(), params)
	if err != nil {
		return nil
	}

	var items = make([]*web.GetRankingResponse_RankItem, 0, len(list))
	for _, item := range list {
		items = append(items, &web.GetRankingResponse_RankItem{
			UserName:    item.UserName,
			Score:       int32(item.Score),
			TotalCommit: int32(item.TotalCommit),
		})
	}

	response.NorResponse(ctx.Context, &web.GetRankingResponse{UserList: items}, "查询成功！")
	return nil
}
