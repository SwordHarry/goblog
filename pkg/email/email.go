package email

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

// 对发送电子邮件的行为进行封装

type Email struct {
	*SMTPInfo
}

type SMTPInfo struct {
	Host     string
	Port     int
	IsSSL    bool
	UserName string
	PassWord string
	From     string // 发件人
}

func NewEmail(info *SMTPInfo) *Email {
	return &Email{SMTPInfo: info}
}

// 发送邮件
func (e *Email) SendMail(subject, body string, to []string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.From)
	m.SetHeader("To", to...)        // 收件人
	m.SetHeader("Subject", subject) // 主题
	m.SetBody("text/html", body)    // 正文
	// 实例化 拨号器
	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.PassWord)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.IsSSL}
	return dialer.DialAndSend(m)
}
