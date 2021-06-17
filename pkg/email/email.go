/**
 * @Author: Anpw
 * @Description:
 * @File:  email
 * @Version: 1.0.0
 * @Date: 2021/6/15 22:55
 */

package email

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type Email struct {
	*SMTPInfo
}

type SMTPInfo struct {
	Host     string
	Port     int
	isSSL    bool
	UserName string
	PassWord string
	From     string
}

func NewEmail(info *SMTPInfo) *Email {
	return &Email{SMTPInfo: info}
}

func (e *Email) SendMail(to []string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.From)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.PassWord)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.isSSL}
	return dialer.DialAndSend(m)
}