package service

var _ IEmailService = (*EmailService)(nil)

type IEmailService interface {
}

type EmailService struct {
}
