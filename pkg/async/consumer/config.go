package consumer

import (
	"context"
	"fmt"

	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api"
	"github.com/CrazyThursdayV50/gotils/pkg/logger"
	defaultLogger "github.com/CrazyThursdayV50/gotils/pkg/logger/default"
)

type Config[T any] struct {
	name       string
	ctx        context.Context
	ch         api.ChanAPI[T]
	logger     logger.Logger
	handler    consumerHandler[T]
	errHandler func(error)
}

func NewConfig[T any](name string) *Config[T] {
	var conf Config[T]
	conf.name = name
	conf.ctx = context.TODO()
	conf.handler = func(e T, l logger.Logger) {
		if l != nil {
			l.Debug("default handler, event: %v", e)
			return
		}

		fmt.Printf("default consumer handler, event: %v\n", e)
	}
	conf.logger = defaultLogger.DefaultLogger()
	conf.errHandler = func(err error) {
		if err == nil {
			return
		}
		conf.logger.Error("%v", err)
	}
	return &conf
}

func (c *Config[T]) SetContext(ctx context.Context) *Config[T] {
	c.ctx = ctx
	return c
}

func (c *Config[T]) SetChannel(ch api.ChanAPI[T]) *Config[T] {
	c.ch = ch
	return c
}

func (c *Config[T]) SetLogger(logger logger.Logger) *Config[T] {
	c.logger = logger
	return c
}

func (c *Config[T]) SetHandler(handler consumerHandler[T]) *Config[T] {
	c.handler = handler
	return c
}

func (c *Config[T]) SetErrHandler(handler func(error)) *Config[T] {
	c.errHandler = handler
	return c
}

func (c *Config[T]) Build() *Consumer[T] {
	var consumer Consumer[T]
	consumer.done = make(chan struct{})
	consumer.ctx, consumer.cancel = context.WithCancel(c.ctx)
	consumer.ch = c.ch
	consumer.logger = c.logger
	consumer.handle = c.handler
	consumer.errHandler = c.errHandler
	return &consumer
}
