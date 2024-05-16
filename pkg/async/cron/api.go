package cron

import (
	"context"
	"time"

	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api/slice"
)

func New(opts ...Option) *Cron {
	var c Cron
	opts = append(defaultOptions(), opts...)
	_ = slice.From(opts...).IterFully(func(_ int, opt Option) error { opt(&c); return nil })
	return &c
}

func WithContext(ctx context.Context) Option {
	return func(c *Cron) {
		c.ctx = ctx
	}
}

func WithJob(job func(), tick time.Duration) Option {
	return func(c *Cron) {
		c.job = job
		c.tick = tick
	}
}

func WithRunAfterStart(duration time.Duration) Option {
	return func(c *Cron) {
		c.runAfter = duration
	}
}

func WithWaitAfterRun(ok bool) Option {
	return func(c *Cron) {
		c.waitAfterRun = ok
	}
}
