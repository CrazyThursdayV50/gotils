package email

import (
	"github.com/CrazyThursdayV50/gotils/pkg/tools/email"
	"github.com/CrazyThursdayV50/gotils/pkg/tools/email/mailer"
)

func NewSender(opts ...mailer.Option) (email.MailSender, error) {
	return mailer.New(opts...)
}
