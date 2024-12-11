package service

import (
	service "c_program_edu-gin/internal/app/service/web/v1"
	"github.com/google/wire"
)

var WebProviderSet = wire.NewSet(
	wire.Struct(new(service.EmailService), "*"),
	wire.Bind(new(service.IEmailService), new(*service.EmailService)),
)