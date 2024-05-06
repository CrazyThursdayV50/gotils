package worker

import (
	"context"

	"github.com/CrazyThursdayV50/gotils/pkg/async/monitor"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api"
	gchan "github.com/CrazyThursdayV50/gotils/pkg/builtin/api/chan"
)

func (w *Worker[J]) WithContext(ctx context.Context) {
	w.Monitor = monitor.New(ctx)
}

func (w *Worker[J]) WithGraceful(ok bool) {
	w.graceful = ok
}

func (w *Worker[J]) WithBuffer(buffer int) {
	w.trigger = gchan.Make[J](buffer)
}

func (w *Worker[J]) WithTrigger(trigger api.ChanAPIR[J]) {
	w.trigger = trigger
}

func (m *Worker[J]) Run() {
	m.run()
	m.Monitor.Run()
}

func (m *Worker[J]) Stop() { m.Monitor.Stop() }

func (m *Worker[J]) Count() int64 { return m.count }
