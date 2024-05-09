package email

type MailMessager interface {
	To() []string
	Message() []byte
}

type MailSender interface {
	From() string
	Send(MailMessager) error
}
