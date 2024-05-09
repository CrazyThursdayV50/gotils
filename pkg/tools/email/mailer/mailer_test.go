package mailer

import (
	"testing"

	"github.com/CrazyThursdayV50/gotils/pkg/tools/email/mail"
)

func TestMailer(t *testing.T) {
	var username = "mailer tester"
	var sender = "xxx@qq.com"
	var receiver = "yyy@qq.com"
	var password = ""
	var endpoint = "smtp.qq.com"
	var port = 465
	var mailer, err = New(
		WithUsername(username),
		WithSender(sender, password),
		WithSMTP(endpoint, port),
	)
	if err != nil {
		t.Fatalf("new mailer failed: %v", err)
	}

	var to = []string{
		sender,
		receiver,
	}
	var cc = []string{
		sender,
		receiver,
	}

	var mail = mail.New(
		mail.WithFrom(sender),
		mail.WithTo(to...),
		mail.WithCc(cc...),
		mail.WithSubject("MailerTest"),
		mail.WithBody("mailer test"),
	)

	t.Logf("mail body: %s", mail.Message())
	err = mailer.Send(mail)
	if err != nil {
		t.Fatalf("send mail failed: %v", err)
	}
}
