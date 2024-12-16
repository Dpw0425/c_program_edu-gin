package v1

import (
	"c_program_edu-gin/internal/app/service/web/v1"
	"c_program_edu-gin/internal/pkg/context"
	myErr "c_program_edu-gin/pkg/error"
	"c_program_edu-gin/pkg/response"
	web "c_program_edu-gin/schema/genproto/web/v1/common"
)

type Common struct {
	EmailService service.EmailService
}

func (c *Common) SendEmailCode(ctx *ctx.Context) error {
	params := &web.SendEmailCodeRequest{}
	if err := ctx.Context.ShouldBind(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	err := c.EmailService.Send(ctx.Ctx(), params.Channel, params.Email)
	if err != nil {
		return err
	}

	response.NorResponse(ctx.Context, &web.SendEmailCodeResponse{}, "发送成功！")

	return nil
}
