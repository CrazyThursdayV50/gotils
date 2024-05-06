package models

import (
	"sync"
	"time"
)

type (
	chanWrite[E any] interface {
		chan E | chan<- E
	}

	ChanW[E any] struct {
		l           *sync.Mutex
		done        chan struct{}
		recvTimeout time.Duration
		sendTimeout time.Duration
		c           chan<- E
	}
)

func FromChanW[E any, C chanWrite[E]](c C) *ChanW[E] {
	return &ChanW[E]{
		l:    &sync.Mutex{},
		done: make(chan struct{}),
		c:    c,
	}
}

func (c *ChanW[E]) Len() int {
	return len(c.c)
}

func (c *ChanW[E]) IsEmpty() bool { return c.Len() == 0 }

func (c *ChanW[E]) Closed() bool {
	if len(c.c) == 0 && cap(c.c) == 0 {
		return true
	}

	return false
}

func (c *ChanW[E]) Close() {
	c.l.Lock()
	defer c.l.Unlock()
	if c.Closed() {
		return
	}
	close(c.done)
	close(c.c)
}

func (c *ChanW[E]) Send(e E) {
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

func (c *ChanW[E]) ChanW() chan<- E {
	return c.c
}

func (c *ChanW[E]) SendTimeout(send time.Duration) {
	c.sendTimeout = send
}
