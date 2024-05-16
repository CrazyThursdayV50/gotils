package mailer

import (
	"crypto/tls"
	"fmt"
	"net/smtp"

	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api/slice"
	"github.com/CrazyThursdayV50/gotils/pkg/tools/email"
)

type Mailer struct {
	username     string
	from         string
	password     string
	smtpEndpoint string
	smtpPort     int
	smtpAuth     smtp.Auth
	client       *smtp.Client
}

type Option func(*Mailer)

func WithSender(account, password string) Option {
	return func(m *Mailer) {
		m.from, m.password = account, password
	}
}

func WithSMTP(endpoint string, port int) Option {
	return func(m *Mailer) {
		m.smtpEndpoint, m.smtpPort = endpoint, port
	}
}

func WithUsername(username string) Option {
	return func(m *Mailer) {
		m.username = username
	}
}

func (m *Mailer) smtpServer() string {
	return fmt.Sprintf("%s:%d", m.smtpEndpoint, m.smtpPort)
}

func New(opts ...Option) (*Mailer, error) {
	var m Mailer
	_ = slice.From(opts...).IterFully(func(_ int, opt Option) error {
		opt(&m)
		return nil
	})

	m.smtpAuth = smtp.PlainAuth(m.username, m.from, m.password, m.smtpEndpoint)

	tlsConfig := tls.Config{
		InsecureSkipVerify: true,
		ServerName:         m.smtpEndpoint,
	}

	conn, err := tls.Dial("tcp", m.smtpServer(), &tlsConfig)
	if err != nil {
		return nil, err
	}

	client, err := smtp.NewClient(conn, m.smtpEndpoint)
	if err != nil {
		return nil, err
	}
	m.client = client

	err = m.client.Auth(m.smtpAuth)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (m *Mailer) From() string { return m.from }

func (m *Mailer) Send(mail email.MailMessager) error {
	err := m.client.Mail(m.from)
	if err != nil {
		fmt.Printf("mail\n")
		return err
	}
	for _, to := range mail.To() {
		err = m.client.Rcpt(to)
		if err != nil {
			fmt.Printf("rcpt\n")
			return err
		}
	}
	w, err := m.client.Data()
	if err != nil {
		fmt.Printf("data\n")
		return err
	}
	_, err = w.Write(mail.Message())
	if err != nil {
		fmt.Printf("write\n")
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return nil
}

func (m *Mailer) Close() {
	m.client.Quit()
}
