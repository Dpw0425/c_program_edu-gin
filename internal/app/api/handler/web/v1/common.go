package v1

import (
	"c_program_edu-gin/internal/app/service/web/v1"
	"c_program_edu-gin/internal/pkg/context"
)

type Common struct {
	EmailService service.EmailService
}

func (c *Common) SendEmailCode(ctx *ctx.Context) error {
	return nil
}
