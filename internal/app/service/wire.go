package service

import (
	adminservice "c_program_edu-gin/internal/app/service/admin/v1"
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

	wire.Struct(new(service.QuestionService), "*"),
	wire.Bind(new(service.IQuestionService), new(*service.QuestionService)),
)

var AdminProviderSet = wire.NewSet(
	wire.Struct(new(adminservice.AdminService), "*"),
	wire.Bind(new(adminservice.IAdminService), new(*adminservice.AdminService)),

	wire.Struct(new(adminservice.TagService), "*"),
	wire.Bind(new(adminservice.ITagService), new(*adminservice.TagService)),

	wire.Struct(new(adminservice.QuestionService), "*"),
	wire.Bind(new(adminservice.IQuestionService), new(*adminservice.QuestionService)),

	wire.Struct(new(adminservice.AuthService), "*"),
	wire.Bind(new(adminservice.IAuthService), new(*adminservice.AuthService)),

	wire.Struct(new(adminservice.BankService), "*"),
	wire.Bind(new(adminservice.IBankService), new(*adminservice.BankService)),

	wire.Struct(new(adminservice.CompetitionService), "*"),
	wire.Bind(new(adminservice.ICompetitionService), new(*adminservice.CompetitionService)),
)
