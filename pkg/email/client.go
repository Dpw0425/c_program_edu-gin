package email

import (
	"c_program_edu-gin/pkg/logger"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

type Client struct {
	conf *ClientConf
}

type ClientConf struct {
	Host     string
	Smtp     string
	Addr     string // 邮箱号
	Name     string // 发件人
	Password string // 密码
}

func NewEmailClient(conf *ClientConf) *Client {
	return &Client{conf: conf}
}

type Option struct {
	To      []string // 收件人
	Subject string   // 邮件主题
	Content []byte   // 邮件内容
}

type OptionFunc func(msg *email.Email)

func (c *Client) SendEmail(send *Option, opt ...OptionFunc) bool {
	e := &email.Email{
		From:    fmt.Sprintf("%v <%v>", c.conf.Name, c.conf.Addr),
		To:      send.To,
		Subject: send.Subject,
		HTML:    send.Content,
	}

	for _, o := range opt {
		o(e)
	}

	return c.do(e)
}

func (c *Client) do(msg *email.Email) bool {
	err := msg.Send(c.conf.Smtp, smtp.PlainAuth("", c.conf.Addr, c.conf.Password, c.conf.Host))
	if err != nil {
		logger.Panicf("Email Client Error: %v!", err)
		return false
	}

	return true
}
