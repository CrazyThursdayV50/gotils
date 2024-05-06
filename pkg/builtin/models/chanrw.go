package models

import (
	"sync"
	"time"

	"github.com/CrazyThursdayV50/gotils/pkg/wrapper"
	"github.com/CrazyThursdayV50/gotils/pkg/wrapper/wrap"
)

type ChanRW[E any] struct {
	l           *sync.Mutex
	done        chan struct{}
	recvTimeout time.Duration
	sendTimeout time.Duration
	c           chan E
}

func FromChan[E any](c chan E) *ChanRW[E] {
	return &ChanRW[E]{
		l:    &sync.Mutex{},
		done: make(chan struct{}),
		c:    c,
	}
}

func (c *ChanRW[E]) Len() int {
	return len(c.c)
}

func (c *ChanRW[E]) IsEmpty() bool { return c.Len() == 0 }

func (c *ChanRW[E]) Closed() bool {
	select {
	case <-c.done:
		return true
	default:
		return false
	}
}

func (c *ChanRW[E]) Close() {
	c.l.Lock()
	defer c.l.Unlock()
	if c.Closed() {
		return
	}
	close(c.done)
	close(c.c)
}

func (c *ChanRW[E]) Receive() (wrapper.UnWrapper[E], bool) {
	if c.recvTimeout <= 0 {
		element := <-c.c
		return wrap.Wrap(element), true
	}

	timer := time.NewTimer(c.recvTimeout)
	select {
	case element := <-c.c:
		return wrap.Wrap(element), true

	case <-timer.C:
		return nil, false
	}
}

func (c *ChanRW[E]) Send(e E) {
	if c.Closed() {
		return
	}

	if c.sendTimeout <= 0 {
		c.c <- e
		return
	}

	timer := time.NewTimer(c.sendTimeout)
	select {
	case <-timer.C:
	case c.c <- e:
	}
}

func (c *ChanRW[E]) Chan() chan E {
	return c.c
}

func (c *ChanRW[E]) ChanR() <-chan E {
	return c.c
}

func (c *ChanRW[E]) ChanW() chan<- E {
	return c.c
}

func (c *ChanRW[E]) IterFunc(f func(E) bool) {
	for e := range c.c {
		ok := f(e)
		if !ok {
			return
		}
	}
}

func (c *ChanRW[E]) IterFuncFully(f func(E)) {
	for e := range c.c {
		f(e)
	}
}

func (c *ChanRW[E]) IterFuncMut(f func(E, *ChanRW[E]) bool) {
	for e := range c.c {
		ok := f(e, c)
		if !ok {
			return
		}
	}
}

func (c *ChanRW[E]) IterFuncMutFully(f func(E, *ChanRW[E])) {
	for e := range c.c {
		f(e, c)
	}
}

func (c *ChanRW[E]) Renew(buff int) {
	if !c.Closed() {
		return
	}

	c.done = make(chan struct{})
	c.c = make(chan E, buff)
}

func (c *ChanRW[E]) RenewForce(buff int) {
	c.c = make(chan E, buff)
	if c.Closed() {
		c.done = make(chan struct{})
	}
}

func (c *ChanRW[E]) SendTimeout(send time.Duration) {
	c.sendTimeout = send
}

func (c *ChanRW[E]) RecvTimeout(recv time.Duration) {
	c.recvTimeout = recv
}
