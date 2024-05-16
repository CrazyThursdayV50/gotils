package mail

import (
	"bytes"
	"fmt"
	"net/mail"
	"net/textproto"
	"strings"

	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api/slice"
)

const (
	ContentType = "Content-Type"
	ContentHTML = "text/html; charset=UTF-8"
)

type Mail struct {
	body   string
	header textproto.MIMEHeader
}

type option func(*Mail)

func WithFrom(from string) option {
	return func(m *Mail) {
		m.header.Set("From", from)
	}
}

func WithTo(to ...string) option {
	return func(m *Mail) {
		for _, to := range to {
			m.header.Add("To", to)
		}
	}
}

func WithCc(cc ...string) option {
	return func(m *Mail) {
		for _, cc := range cc {
			m.header.Add("Cc", cc)
		}
	}
}

func WithSubject(subject string) option {
	return func(m *Mail) {
		m.header.Set("Subject", subject)
	}
}

func WithBody(body string) option {
	return func(m *Mail) {
		m.body = body
	}
}

func WithHeader(k, v string) option {
	return func(m *Mail) {
		m.header.Add(k, v)
	}
}

func WithHTML() option {
	return func(m *Mail) {
		m.header.Set(ContentType, ContentHTML)
	}
}

func New(opts ...option) *Mail {
	var m Mail
	m.header = make(textproto.MIMEHeader)
	_ = slice.From(opts...).IterFully(func(_ int, opt option) error {
		opt(&m)
		return nil
	})
	return &m
}

func (m *Mail) To() []string { return m.header["To"] }

func (m *Mail) Message() []byte {
	var message mail.Message
	message.Header = mail.Header(m.header)
	message.Body = bytes.NewBufferString(m.body)
	var msg strings.Builder
	for k, v := range message.Header {
		msg.WriteString(fmt.Sprintf("%s: %s\n", k, strings.Join(v, ",")))
	}
	msg.WriteString("\n")
	msg.WriteString(m.body)
	return []byte(msg.String())
}
