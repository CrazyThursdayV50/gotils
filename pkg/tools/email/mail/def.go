package mail

import (
	"fmt"
	"strings"

	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api/slice"
)

type Mail struct {
	from    string
	to      []string
	cc      []string
	subject string
	body    string
}

const (
	template = `From: %s
To: %s
Cc: %s
Subject: %s

%s`
)

type option func(*Mail)

func WithFrom(from string) option {
	return func(m *Mail) {
		m.from = from
	}
}

func WithTo(to ...string) option {
	return func(m *Mail) {
		m.to = to
	}
}

func WithCc(cc ...string) option {
	return func(m *Mail) {
		m.cc = cc
	}
}

func WithSubject(subject string) option {
	return func(m *Mail) {
		m.subject = subject
	}
}

func WithBody(body string) option {
	return func(m *Mail) {
		m.body = body
	}
}

func New(opts ...option) *Mail {
	var m Mail
	slice.From(opts).IterFuncFully(func(opt option) {
		opt(&m)
	})
	return &m
}

func (m *Mail) From() string { return m.from }
func (m *Mail) To() []string { return m.to }

func (m *Mail) Message() []byte {
	return []byte(fmt.Sprintf(template, m.from, strings.Join(m.to, ","), strings.Join(m.cc, ","), m.subject, m.body))
}
