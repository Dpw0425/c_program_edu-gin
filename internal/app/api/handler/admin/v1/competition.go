package v1

import (
	adminservice "c_program_edu-gin/internal/app/service/admin/v1"
	ctx "c_program_edu-gin/internal/pkg/context"
	myErr "c_program_edu-gin/pkg/error"
	"c_program_edu-gin/pkg/response"
	admin "c_program_edu-gin/schema/genproto/admin/v1/competition"
	"strings"
)

type Competition struct {
	CompetitionService adminservice.CompetitionService
}

func (c *Competition) List(ctx *ctx.Context) error {
	params := &admin.CompetitionListRequest{}
	if err := ctx.Context.ShouldBindQuery(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	list, count, err := c.CompetitionService.List(ctx.Ctx(), params)
	if err != nil {
		return err
	}

	items := make([]*admin.CompetitionListResponse_CompetitionItem, 0, len(list))
	for _, item := range list {
		items = append(items, &admin.CompetitionListResponse_CompetitionItem{
			Id:         int32(item.ID),
			Name:       item.Name,
			Contestant: strings.Split(item.Contestant, ","),
			OwnerId:    int32(item.OwnerID),
			StartTime:  item.StartTime.Format("2006-01-02 15:04:05"),
			Deadline:   item.Deadline.Format("2006-01-02 15:04:05"),
			Status:     int32(item.Status),
			Category:   int32(item.Category),
			Permission: int32(item.Permission),
		})
	}

	response.NorResponse(ctx.Context, &admin.CompetitionListResponse{
		CompetitionList: items,
		Total:           int32(count),
	}, "查询成功！")
	return nil
}

func (c *Competition) Add(ctx *ctx.Context) error {
	params := &admin.AddCompetitionRequest{}
	if err := ctx.Context.ShouldBind(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	if result := c.CompetitionService.Add(ctx.Ctx(), uint(ctx.UserID()), params); result != nil {
		return result
	}

	response.NorResponse(ctx.Context, &admin.AddCompetitionResponse{}, "添加成功！")
	return nil
}

func (c *Competition) Update(ctx *ctx.Context) error {
	params := &admin.UpdateCompetitionRequest{}
	if err := ctx.Context.ShouldBind(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	if result := c.CompetitionService.Update(ctx.Ctx(), params); result != nil {
		return result
	}

	response.NorResponse(ctx.Context, &admin.UpdateCompetitionResponse{}, "修改成功！")
	return nil
}

func (c *Competition) Delete(ctx *ctx.Context) error {
	params := &admin.DeleteCompetitionRequest{}
	if err := ctx.Context.ShouldBindQuery(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	if result := c.CompetitionService.Delete(ctx.Ctx(), params); result != nil {
		return result
	}

	response.NorResponse(ctx.Context, &admin.DeleteCompetitionResponse{}, "删除成功！")
	return nil
}
