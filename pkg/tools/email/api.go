package email

type MailMessager interface {
	From() string
	To() []string
	Message() []byte
}
type MailSender interface {
	Send(MailMessager) error
}
