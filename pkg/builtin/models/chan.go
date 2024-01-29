package models

import "sync"

type Chan[E any] struct {
	l    *sync.Mutex
	done chan struct{}
	c    chan E
}

func FromChan[E any](c chan E) *Chan[E] {
	return &Chan[E]{
		l:    &sync.Mutex{},
		done: make(chan struct{}),
		c:    c,
	}
}

func (c *Chan[E]) Len() int {
	return len(c.c)
}

func (c *Chan[E]) IsEmpty() bool { return c.Len() == 0 }

func (c *Chan[E]) Closed() bool {
	select {
	case <-c.done:
		return true
	default:
		return false
	}
}

func (c *Chan[E]) Close() {
	c.l.Lock()
	defer c.l.Unlock()
	if c.Closed() {
		return
	}
	close(c.done)
}

func (c *Chan[E]) Receive() E {
	return <-c.c
}

func (c *Chan[E]) Send(e E) {
	if c.Closed() {
		return
	}
	c.c <- e
}

func (c *Chan[E]) Chan() chan E {
	return c.c
}

func (c *Chan[E]) IterFunc(f func(E) bool) {
	for e := range c.c {
		ok := f(e)
		if !ok {
			return
		}
	}
}

func (c *Chan[E]) IterFuncFully(f func(E)) {
	for e := range c.c {
		f(e)
	}
}

func (c *Chan[E]) IterFuncMut(f func(E, *Chan[E]) bool) {
	for e := range c.c {
		ok := f(e, c)
		if !ok {
			return
		}
	}
}

func (c *Chan[E]) IterFuncMutFully(f func(E, *Chan[E])) {
	for e := range c.c {
		f(e, c)
	}
}

func (c *Chan[E]) Renew(buff int) {
	if !c.Closed() {
		return
	}

	c.done = make(chan struct{})
	c.c = make(chan E, buff)
}

func (c *Chan[E]) RenewForce(buff int) {
	c.c = make(chan E, buff)
	if c.Closed() {
		c.done = make(chan struct{})
	}
}
