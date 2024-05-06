package models

import (
	"sync"
	"time"

	"github.com/CrazyThursdayV50/gotils/pkg/wrapper"
	"github.com/CrazyThursdayV50/gotils/pkg/wrapper/wrap"
)

type (
	chanRead[E any] interface {
		chan E | <-chan E
	}

	ChanR[E any] struct {
		l           *sync.Mutex
		done        chan struct{}
		recvTimeout time.Duration
		sendTimeout time.Duration
		c           <-chan E
	}
)

func FromChanR[E any, C chanRead[E]](c C) *ChanR[E] {
	return &ChanR[E]{
		l:    &sync.Mutex{},
		done: make(chan struct{}),
		c:    c,
	}
}

func (c *ChanR[E]) Len() int {
	return len(c.c)
}

func (c *ChanR[E]) IsEmpty() bool { return c.Len() == 0 }

func (c *ChanR[E]) Closed() bool {
	_, ok := <-c.c
	if !ok {
		return true
	}
	return false
}

func (c *ChanR[E]) ChanR() <-chan E {
	return c.c
}

func (c *ChanR[E]) RecvTimeout(recv time.Duration) {
	c.recvTimeout = recv
}

func (c *ChanR[E]) Receive() (wrapper.UnWrapper[E], bool) {
	if c.recvTimeout <= 0 {
		element, ok := <-c.c
		return wrap.Wrap(element), ok
	}

	timer := time.NewTimer(c.recvTimeout)
	select {
	case element, ok := <-c.c:
		return wrap.Wrap(element), ok

	case <-timer.C:
		return nil, false
	}
}

func (c *ChanR[E]) IterFunc(f func(E) bool) {
	for e := range c.c {
		ok := f(e)
		if !ok {
			return
		}
	}
}

func (c *ChanR[E]) IterFuncFully(f func(E)) {
	for e := range c.c {
		f(e)
	}
}
