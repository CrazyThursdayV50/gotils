package worker

import (
	"github.com/CrazyThursdayV50/gotils/pkg/async/goo"
	"github.com/CrazyThursdayV50/gotils/pkg/async/monitor"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api"
)

type Worker[J any] struct {
	*monitor.Monitor
	do       func(J)
	count    int64
	trigger  api.ChanAPI[J]
	graceful bool
}

func (m *Worker[J]) onJob() {
	goo.Go(func() {
		for {
			select {
			case <-m.Done():
				if !m.graceful {
					return
				}

				if m.trigger.IsEmpty() {
					return
				}

			case job := <-m.trigger.Chan():
				m.do(job)
			}
		}
	})
}
