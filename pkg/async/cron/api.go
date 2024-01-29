package cron

import (
	"context"
	"time"

	"github.com/CrazyThursdayV50/gotils/pkg/async/monitor"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api/slice"
)

func New(opts ...Option) *Cron {
	opts = append(defaultOptions(), opts...)
	var c Cron
	slice.From(opts).IterFuncFully(func(opt Option) { opt(&c) })
	return &c
}

func WithContext(ctx context.Context) Option {
	return func(c *Cron) {
		c.Monitor = monitor.New(ctx)
	}
}

func WithTrigger(t func(), tick time.Duration) Option {
	return func(c *Cron) {
		c.trigger = t
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
