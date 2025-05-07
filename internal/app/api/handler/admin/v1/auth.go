package v1

import (
	adminservice "c_program_edu-gin/internal/app/service/admin/v1"
	ctx "c_program_edu-gin/internal/pkg/context"
	myErr "c_program_edu-gin/pkg/error"
	"c_program_edu-gin/pkg/response"
	admin "c_program_edu-gin/schema/genproto/admin/v1/auth"
)

type Auth struct {
	AuthService adminservice.AuthService
}

func (a *Auth) UserList(ctx *ctx.Context) error {
	params := &admin.UserListRequest{}
	if err := ctx.Context.ShouldBindQuery(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	list, count, err := a.AuthService.UserList(ctx.Ctx(), params)
	if err != nil {
		return err
	}

	items := make([]*admin.UserListResponse_UserItem, 0, len(list))
	for _, item := range list {
		items = append(items, &admin.UserListResponse_UserItem{
			UserId:    item.UserID,
			UserName:  item.UserName,
			StudentId: item.StudentID,
			Grade:     int32(item.Grade),
			Status:    int32(item.Status),
		})
	}

	response.NorResponse(ctx.Context, &admin.UserListResponse{
		UserList: items,
		Total:    int32(count),
	}, "查询成功！")
	return nil
}

func (a *Auth) AdminList(ctx *ctx.Context) error {
	params := &admin.AdminListRequest{}
	if err := ctx.Context.ShouldBindQuery(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	list, count, err := a.AuthService.AdminList(ctx.Ctx(), params)
	if err != nil {
		return err
	}

	items := make([]*admin.AdminListResponse_AdminItem, 0, len(list))
	for _, item := range list {
		items = append(items, &admin.AdminListResponse_AdminItem{
			Id:         int32(item.ID),
			UserName:   item.UserName,
			TeacherId:  item.TeacherID,
			Permission: item.Permission,
			Status:     int32(item.Status),
		})
	}

	response.NorResponse(ctx.Context, &admin.AdminListResponse{
		AdminList: items,
		Total:     int32(count),
	}, "查询成功！")
	return nil
}
