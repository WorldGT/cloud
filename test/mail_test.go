package test

import (
	"crypto/tls"
	"net/smtp"
	"testing"

	"github.com/jordan-wright/email"
)

func TestSendMail(t *testing.T) {
	e := email.NewEmail()
	e.From = "Get <gtong404@163.com>"
	e.To = []string{"346614213@qq.com"}
	e.Subject = "验证码测试"
	e.HTML = []byte("你的验证码为：<h1>132456</h1>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "gtong404@163.com", "HBEJKTLRBQMZSKCM", "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})

	if err != nil {
		t.Fatal(err)
	}

}

//HBEJKTLRBQMZSKCM
