package models

import (
	"sync"
	"time"

	"github.com/CrazyThursdayV50/gotils/pkg/async/goo"
)

type (
	ChanWrite[E any] interface {
		chan E | chan<- E
	}

	ChanW[E any] struct {
		l           *sync.Mutex
		done        bool
		recvTimeout time.Duration
		sendTimeout time.Duration
		c           chan<- E
		sendwg      sync.WaitGroup
	}
)

func FromChanW[E any, C ChanWrite[E]](c C) *ChanW[E] {
	return &ChanW[E]{
		l: &sync.Mutex{},
		c: c,
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

func (c *ChanW[E]) closeSendChan() {
	c.sendwg.Wait()
	close(c.c)
}

func (c *ChanW[E]) Close() {
	c.l.Lock()
	defer c.l.Unlock()
	if c.Closed() {
		return
	}
	c.done = true
	goo.Go(c.closeSendChan)
}

func (c *ChanW[E]) Send(e E) {
	if c.done {
		return
	}

	c.sendwg.Add(1)
	defer c.sendwg.Done()

	if c.sendTimeout <= 0 {
		c.c <- e
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
