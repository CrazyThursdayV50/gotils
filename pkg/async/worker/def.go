package worker

import (
	"github.com/CrazyThursdayV50/gotils/pkg/async/goo"
	"github.com/CrazyThursdayV50/gotils/pkg/async/monitor"
)

type Worker[J any] struct {
	*monitor.Monitor
	do       func(J)
	count    int64
	trigger  chan J
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

				if len(m.trigger) == 0 {
					return
				}

			case job := <-m.trigger:
				m.do(job)
			}
		}
	})
}
