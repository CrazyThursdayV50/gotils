package cron

import (
	"context"
	"time"

	"github.com/CrazyThursdayV50/gotils/pkg/async/goo"
	"github.com/CrazyThursdayV50/gotils/pkg/async/worker"
)

type Cron struct {
	ctx          context.Context
	job          func()
	runAfter     time.Duration
	waitAfterRun bool
	tick         time.Duration
	worker       *worker.Worker[time.Time]
	delivery     func(time.Time)
	runOnStart   func()
}

type Option func(*Cron)

func defaultOptions() []Option {
	return []Option{
		WithContext(context.TODO()),
		WithJob(func() {}, time.Minute),
		WithRunAfterStart(-1),
		WithWaitAfterRun(false),
	}
}

func timerDo(duration time.Duration, done <-chan struct{}, do func()) {
	timer := time.NewTimer(duration)
	select {
	case <-done:
		return
	case <-timer.C:
		do()
	}
}

func (c *Cron) init() {
	c.worker, c.delivery = worker.New(func(time.Time) { c.job() })
	c.worker.WithContext(c.ctx)
	c.worker.WithGraceful(false)

	if c.runAfter < 0 {
		return
	}

	do := func() { c.delivery(time.Now()) }
	if c.runAfter == 0 {
		c.runOnStart = do
		return
	}

	c.runOnStart = func() { timerDo(c.runAfter, c.worker.Done(), do) }
}

func (c *Cron) tickRun() {
	ticker := time.NewTicker(c.tick)
	goo.Go(func() {
		for {
			select {
			case <-c.worker.Done():
				return

			case t := <-ticker.C:
				select {
				case <-c.worker.Done():
				default:
					c.delivery(t)
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
			case <-c.worker.Done():
				return

			case t := <-timer.C:
				select {
				case <-c.worker.Done():
				default:
					c.delivery(t)
					timer.Stop()
					timer.Reset(c.tick)
				}
			}
		}
	})
}
