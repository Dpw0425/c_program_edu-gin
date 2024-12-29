package v1

import (
	"c_program_edu-gin/internal/app/constants"
	"c_program_edu-gin/internal/app/schema"
	service "c_program_edu-gin/internal/app/service/web/v1"
	"c_program_edu-gin/internal/config"
	ctx "c_program_edu-gin/internal/pkg/context"
	myErr "c_program_edu-gin/pkg/error"
	"c_program_edu-gin/pkg/response"
	web "c_program_edu-gin/schema/genproto/web/v1/user"
)

type User struct {
	Config       *config.Config
	UserService  service.IUserService
	EmailService service.IEmailService
}

func (u *User) Register(ctx *ctx.Context) error {
	params := &web.UserRegisterRequest{}
	if err := ctx.Context.ShouldBind(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	if !u.EmailService.Verify(ctx.Ctx(), constants.EmailRegisterChannel, params.Email, params.VerifyCode) {
		return myErr.BadRequest("wrong_verification_code", "验证码错误！")
	}

	if err := u.UserService.Register(ctx.Ctx(), &schema.UserRegister{
		NickName: params.Nickname,
		Password: params.Password,
		Avatar:   params.Avatar,
		Email:    params.Email,
	}); err != nil {
		return myErr.BadRequest("", "用户注册失败: %v", err)
	}

	// 删除验证码缓存
	u.EmailService.Delete(ctx.Ctx(), constants.EmailRegisterChannel, params.Email)

	response.NorResponse(ctx.Context, &web.UserRegisterResponse{}, "注册成功！")
	return nil
}
