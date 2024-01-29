package cron

import (
	"context"
	"gotils/pkg/async/monitor"
)

func (c *Cron) WithContext(ctx context.Context) {
	c.Monitor = monitor.New(ctx)
}

func (c *Cron) Run() {
	c.Monitor.Run()
	c.start()
	if c.waitAfterRun {
		c.timerRun()
	} else {
		c.tickRun()
	}
}

func (c *Cron) Stop() {
	c.Monitor.Stop()
}
