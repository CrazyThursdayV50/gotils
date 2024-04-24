package cron

import (
	"context"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	var cron = New(
		WithContext(context.TODO()),
		WithJob(func() {
			t.Logf("[NOW]: %s\n", time.Now())
		}, time.Second),
		WithRunAfterStart(time.Second*3),
		WithWaitAfterRun(false),
	)

	cron.Run()
	<-make(chan struct{})
}
