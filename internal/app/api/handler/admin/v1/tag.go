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
	if err := ctx.Context.ShouldBindQuery(params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	list, count, err := t.TagService.List(ctx.Ctx(), params)
	if err != nil {
		return err
	}

	items := make([]*admin.TagListResponse_TagItem, 0, len(list))
	for _, item := range list {
		items = append(items, &admin.TagListResponse_TagItem{
			Id:    int32(item.ID),
			Name:  item.Name,
			Count: int32(item.Count),
		})
	}

	response.NorResponse(ctx.Context, &admin.TagListResponse{
		TagList: items,
		Total:   int32(count),
	}, "查询成功！")
	return nil
}

func (t *Tag) Update(ctx *ctx.Context) error {
	params := &admin.TagUpdateRequest{}
	if err := ctx.Context.ShouldBind(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	if result := t.TagService.Update(ctx.Ctx(), params); result != nil {
		return result
	}

	response.NorResponse(ctx.Context, &admin.TagUpdateResponse{}, "修改成功！")
	return nil
}

func (t *Tag) Delete(ctx *ctx.Context) error {
	params := &admin.DeleteTagRequest{}
	if err := ctx.Context.ShouldBindQuery(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	if result := t.TagService.Delete(ctx.Ctx(), params); result != nil {
		return result
	}

	response.NorResponse(ctx.Context, &admin.DeleteTagResponse{}, "删除成功！")
	return nil
}

func (t *Tag) GetAll(ctx *ctx.Context) error {
	list, err := t.TagService.GetAll(ctx.Ctx())
	if err != nil {
		return err
	}

	items := make([]*admin.TagListResponse_TagItem, 0, len(list))
	for _, item := range list {
		items = append(items, &admin.TagListResponse_TagItem{
			Id:   int32(item.ID),
			Name: item.Name,
		})
	}

	response.NorResponse(ctx.Context, &admin.TagListResponse{
		TagList: items,
	}, "查询成功！")
	return nil
}
