package service

import (
	"c_program_edu-gin/internal/app/constants"
	"c_program_edu-gin/internal/app/storage/cache"
	"c_program_edu-gin/internal/app/storage/repo"
	"c_program_edu-gin/pkg/email"
	myErr "c_program_edu-gin/pkg/error"
	"c_program_edu-gin/utils/generator"
	"context"
	"time"
)

var _ IEmailService = (*EmailService)(nil)

type IEmailService interface {
	Send(ctx context.Context, channel string, email string) error
	Verify(ctx context.Context, channel string, email string, code string) bool
	Delete(ctx context.Context, channel string, email string)
}

type EmailService struct {
	Storage  *cache.EmailStorage
	UserRepo *repo.UserRepo
	Template *TemplateService
	Client   *email.Client
}

func (e *EmailService) Send(ctx context.Context, channel string, to string) error {
	// 请求类型判断
	switch channel {
	case constants.EmailLoginChannel, constants.EmailForgetChannel, constants.EmailChangeChannel:
		if !e.UserRepo.IsExist(ctx, to) {
			return myErr.BadRequest("business_error", "账号不存在！")
		}
	case constants.EmailRegisterChannel:
		if e.UserRepo.IsExist(ctx, to) {
			return myErr.BadRequest("business_error", "当前邮箱号已被注册！")
		}
	default:
		return myErr.BadRequest("", "请求异常！")
	}

	// 生成验证码
	code := generator.Random(6)
	if err := e.Storage.Set(ctx, channel, to, code, 15*time.Minute); err != nil {
		return myErr.BadRequest("", "ERROR: %s", err.Error())
	}

	data := make(map[string]string)
	data["channel"] = channel
	data["code"] = code

	// 按照模板生成邮件内容
	msg, err := e.Template.LoadTemplate(data)
	if err != nil {
		return myErr.InternalServerError("", "生成邮件失败！")
	}

	// TODO: ADD SUBJECT NAME
	// 构建邮件发送对象
	send := &email.Option{
		To:      []string{to},
		Subject: "欢迎您注册CProgramEdu",
		Content: msg,
	}

	// 执行邮件发送
	if !e.Client.SendEmail(send) {
		return myErr.InternalServerError("", "邮件发送失败！")
	}

	return nil
}

func (e *EmailService) Verify(ctx context.Context, channel string, email string, code string) bool {
	return e.Storage.Verify(ctx, channel, email, code)
}

func (e *EmailService) Delete(ctx context.Context, channel string, email string) {
	_ = e.Storage.Del(ctx, channel, email)
}
