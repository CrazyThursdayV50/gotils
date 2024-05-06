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
	trigger  api.ChanAPIR[J]
	graceful bool
}

func (m *Worker[J]) run() {
	switch m.graceful {
	case true:
		goo.Go(func() {
			m.trigger.IterFuncFully(func(element J) {
				m.do(element)
			})
		})

	default:
		goo.Go(func() {
			m.trigger.IterFunc(func(element J) bool {
				select {
				case <-m.Done():
					return false

				default:
					m.do(element)
					return true
				}
			})
		})
	}
}
