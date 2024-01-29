package monitor

import "github.com/CrazyThursdayV50/gotils/pkg/async/goo"

func (s *Monitor) Run() {
	if s.onStart != nil {
		s.onStart()
	}

	goo.Go(func() {
		<-s.ctx.Done()
		close(s.done)

		if s.onExit == nil {
			return
		}

		s.onExit()
	})
}

func (s *Monitor) Stop()                 { s.cancel() }
func (s *Monitor) Done() <-chan struct{} { return s.done }
func (s *Monitor) OnExit(f func())       { s.onExit = f }
func (s *Monitor) OnStart(f func())      { s.onStart = f }
