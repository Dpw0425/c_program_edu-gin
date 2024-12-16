package constants

type EmailSendChannel string

const (
	EmailRegisterChannel = "register"
	EmailLoginChannel    = "login"
	EmailForgetChannel   = "forget_password"
	EmailChangeChannel   = "change_password"
)
