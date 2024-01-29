package cron

import (
	"context"
	"time"

	"github.com/CrazyThursdayV50/gotils/pkg/async/goo"
	"github.com/CrazyThursdayV50/gotils/pkg/async/monitor"
)

type Cron struct {
	*monitor.Monitor
	runAfter     time.Duration
	waitAfterRun bool
	tick         time.Duration
	trigger      func()
}

type Option func(*Cron)

func defaultOptions() []Option {
	return []Option{
		WithContext(context.TODO()),
		WithTrigger(func() {}, time.Minute),
		WithRunAfterStart(-1),
		WithWaitAfterRun(false),
	}
}

func (c *Cron) start() {
	if c.runAfter < 0 {
		return
	}

	if c.runAfter == 0 {
		c.trigger()
		return
	}

	ticker := time.NewTicker(c.runAfter)
	select {
	case <-c.Done():
		return

	case <-ticker.C:
		c.trigger()
	}
}

func (c *Cron) tickRun() {
	ticker := time.NewTicker(c.tick)
	goo.Go(func() {
		for {
			select {
			case <-c.Done():
				return

			case <-ticker.C:
				select {
				case <-c.Done():
				default:
					c.trigger()
				}

			}
		}
	})
}

func (c *Cron) timerRun() {
	timer := time.NewTimer(c.tick)
	goo.Go(func() {
		for {
			select {
			case <-c.Done():
				return

			case <-timer.C:
				select {
				case <-c.Done():
				default:
					c.trigger()
					timer.Stop()
					timer.Reset(c.tick)
				}
			}
		}
	})
}
