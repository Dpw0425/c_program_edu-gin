package service

import (
	service "c_program_edu-gin/internal/app/service/web/v1"
	"github.com/google/wire"
)

var WebProviderSet = wire.NewSet(
	wire.Struct(new(service.EmailService), "*"),
	wire.Bind(new(service.IEmailService), new(*service.EmailService)),

	wire.Struct(new(service.TemplateService), "*"),
	wire.Bind(new(service.ITemplateService), new(*service.TemplateService)),

	wire.Struct(new(service.UserService), "*"),
	wire.Bind(new(service.IUserService), new(*service.UserService)),
)
