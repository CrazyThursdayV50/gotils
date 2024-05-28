package models

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api"
	"github.com/CrazyThursdayV50/gotils/pkg/wrapper"
	"github.com/CrazyThursdayV50/gotils/pkg/wrapper/wrap"
)

var _ api.ChanAPIR[any] = (*ChanR[any])(nil)

type (
	ChanRead[E any] interface {
		chan E | <-chan E
	}

	ChanR[E any] struct {
		l           *sync.Mutex
		done        chan struct{}
		recvTimeout time.Duration
		sendTimeout time.Duration
		c           <-chan E
		count       int64
	}
)

func FromChanR[E any, C ChanRead[E]](c C) *ChanR[E] {
	return &ChanR[E]{
		l:    &sync.Mutex{},
		done: make(chan struct{}),
		c:    c,
	}
}

func (c *ChanR[E]) Len() int {
	if c == nil {
		return 0
	}
	return len(c.c)
}

func (c *ChanR[E]) IsEmpty() bool { return c.Len() == 0 }

func (c *ChanR[E]) Closed() bool {
	if c == nil {
		return true
	}

	_, ok := <-c.c
	if !ok {
		return true
	}
	return false
}

func (c *ChanR[E]) Unwrap() <-chan E {
	if c == nil {
		return nil
	}
	return c.c
}

func (c *ChanR[E]) RecvTimeout(recv time.Duration) {
	if c == nil {
		return
	}
	c.recvTimeout = recv
}

func (c *ChanR[E]) Receive() (wrapper.UnWrapper[E], bool) {
	if c == nil {
		return wrap.Nil[E](), false
	}

	if c.recvTimeout <= 0 {
		element, ok := <-c.c
		if ok {
			atomic.AddInt64(&c.count, 1)
		}
		return wrap.Wrap(element), ok
	}

	timer := time.NewTimer(c.recvTimeout)
	select {
	case element, ok := <-c.c:
		if ok {
			atomic.AddInt64(&c.count, 1)
		}
		return wrap.Wrap(element), ok

	case <-timer.C:
		return wrap.Nil[E](), false
	}
}

func (c *ChanR[E]) IterOkay(f func(index int, element E) bool) wrapper.UnWrapper[int] {
	if c == nil {
		return wrap.Nil[int]()
	}

	for e := range c.c {
		atomic.AddInt64(&c.count, 1)
		ok := f(int(c.count), e)
		if !ok {
			return wrap.Wrap(int(c.count))
		}
	}

	return wrap.Nil[int]()
}

func (c *ChanR[E]) IterError(f func(index int, element E) error) (wrapper.UnWrapper[int], error) {
	if c == nil {
		return wrap.Nil[int](), nil
	}

	for e := range c.c {
		atomic.AddInt64(&c.count, 1)
		err := f(int(c.count), e)
		if err != nil {
			return wrap.Wrap(int(c.count)), err
		}
	}

	return wrap.Nil[int](), nil
}

func (c *ChanR[E]) IterFully(f func(index int, element E) error) (err api.MapAPI[int, error, any]) {
	if c == nil {
		return
	}

	for e := range c.c {
		atomic.AddInt64(&c.count, 1)
		er := f(int(c.count), e)
		if er != nil {
			if err == nil {
				err = MakeMap[int, error, any](0)
			}
			err.Set(int(c.count), er)
		}
	}

	return
}
