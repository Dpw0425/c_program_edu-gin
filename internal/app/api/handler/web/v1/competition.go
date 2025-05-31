package v1

import (
	service "c_program_edu-gin/internal/app/service/web/v1"
	ctx "c_program_edu-gin/internal/pkg/context"
	"c_program_edu-gin/pkg/response"
	web "c_program_edu-gin/schema/genproto/web/v1/competition"
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
