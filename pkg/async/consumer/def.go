package consumer

import (
	"context"

	"github.com/CrazyThursdayV50/gotils/pkg/async/goo"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api"
	"github.com/CrazyThursdayV50/gotils/pkg/logger"
)

type consumerHandler[T any] func(T, logger.Logger)

type Consumer[T any] struct {
	ctx        context.Context
	cancel     context.CancelFunc
	done       chan struct{}
	ch         api.ChanAPI[T]
	logger     logger.Logger
	handle     consumerHandler[T]
	errHandler func(error)
}

func (c *Consumer[T]) Run() {
	ch := c.ch.Inner()

	goo.Goo(func() {
		for {
			select {
			case event := <-ch:
				c.handle(event, c.logger)

			case <-c.done:
				if c.ch.IsEmpty() {
					c.logger.Warn("consumer exit")
					return
				}

				c.logger.Debug("channel is not empty")
			}
		}
	}, c.errHandler)

	goo.Goo(func() {
		<-c.ctx.Done()
		c.logger.Debug("close consumer channel")
		close(c.done)
	}, c.errHandler)
}

func (c *Consumer[T]) Stop() {
	c.cancel()
}
