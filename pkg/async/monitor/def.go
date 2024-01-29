package monitor

import "context"

type Monitor struct {
	ctx     context.Context
	cancel  context.CancelFunc
	done    chan struct{}
	onStart func()
	onExit  func()
}
