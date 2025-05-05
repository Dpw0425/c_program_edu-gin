package v1

import (
	adminservice "c_program_edu-gin/internal/app/service/admin/v1"
	ctx "c_program_edu-gin/internal/pkg/context"
	myErr "c_program_edu-gin/pkg/error"
	"c_program_edu-gin/pkg/response"
	admin "c_program_edu-gin/schema/genproto/admin/v1/tag"
)

type Tag struct {
	TagService adminservice.TagService
}

func (t *Tag) Add(ctx *ctx.Context) error {
	params := &admin.AddTagRequest{}
	if err := ctx.Context.ShouldBind(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	if err := t.TagService.Add(ctx.Ctx(), params.Name); err != nil {
		return err
	}

	response.NorResponse(ctx.Context, &admin.AddTagResponse{}, "添加成功！")
	return nil
}

func (t *Tag) List(ctx *ctx.Context) error {
	params := &admin.TagListRequest{}
	if err := ctx.Context.ShouldBind(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	list, err := t.TagService.List(ctx.Ctx(), params.Search)
	if err != nil {
		return err
	}

	items := make([]*admin.TagListResponse_TagItem, 0, len(list))
	for _, item := range list {
		items = append(items, &admin.TagListResponse_TagItem{
			Name:  item.Name,
			Count: int32(item.Count),
		})
	}

	response.NorResponse(ctx.Context, &admin.TagListResponse{
		TagList: items,
		Total:   int32(len(items)),
	}, "查询成功！")
	return nil
}
