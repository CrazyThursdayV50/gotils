package monitor

import "context"

func New(ctx context.Context) *Monitor {
	var s Monitor
	s.ctx, s.cancel = context.WithCancel(ctx)
	s.done = make(chan struct{})
	return &s
}
