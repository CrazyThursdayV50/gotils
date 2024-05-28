package models

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/CrazyThursdayV50/gotils/pkg/async/goo"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api"
	"github.com/CrazyThursdayV50/gotils/pkg/wrapper"
	"github.com/CrazyThursdayV50/gotils/pkg/wrapper/wrap"
)

var _ api.ChanAPI[any] = (*ChanRW[any])(nil)

type ChanRW[E any] struct {
	l           *sync.Mutex
	done        bool
	recvTimeout time.Duration
	sendTimeout time.Duration
	c           chan E
	sendwg      sync.WaitGroup
	countR      int64
}

func FromChan[E any](c chan E) *ChanRW[E] {
	return &ChanRW[E]{
		l: &sync.Mutex{},
		c: c,
	}
}

func (c *ChanRW[E]) Len() int {
	if c == nil {
		return 0
	}
	return len(c.c)
}

func (c *ChanRW[E]) IsEmpty() bool { return c.Len() == 0 }

func (c *ChanRW[E]) Closed() bool {
	if c == nil {
		return true
	}
	return c.done
}

func (c *ChanRW[E]) closeSendChan() {
	c.sendwg.Wait()
	close(c.c)
}

func (c *ChanRW[E]) Close() {
	if c == nil {
		return
	}
	c.l.Lock()
	defer c.l.Unlock()
	if c.Closed() {
		return
	}

	c.done = true
	goo.Go(c.closeSendChan)
}

func (c *ChanRW[E]) Receive() (wrapper.UnWrapper[E], bool) {
	if c == nil {
		return wrap.Nil[E](), false
	}
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
	if c == nil {
		return
	}
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

func (c *ChanRW[E]) Unwrap() chan E {
	if c == nil {
		return nil
	}
	return c.c
}

func (c *ChanRW[E]) Renew(buff int) {
	if c == nil {
		return
	}
	if !c.Closed() {
		return
	}

	c.done = false
	c.c = make(chan E, buff)
}

func (c *ChanRW[E]) RenewForce(buff int) {
	if c == nil {
		return
	}
	c.c = make(chan E, buff)
	if c.Closed() {
		c.done = false
	}
}

func (c *ChanRW[E]) SendTimeout(send time.Duration) {
	if c == nil {
		return
	}
	c.sendTimeout = send
}

func (c *ChanRW[E]) RecvTimeout(recv time.Duration) {
	if c == nil {
		return
	}
	c.recvTimeout = recv
}

func (c *ChanRW[E]) IterOkay(f func(index int, element E) bool) wrapper.UnWrapper[int] {
	if c == nil {
		return wrap.Nil[int]()
	}

	for e := range c.c {
		atomic.AddInt64(&c.countR, 1)
		ok := f(int(c.countR), e)
		if !ok {
			return wrap.Wrap(int(c.countR))
		}
	}

	return nil
}

func (c *ChanRW[E]) IterError(f func(index int, element E) error) (wrapper.UnWrapper[int], error) {
	if c == nil {
		return wrap.Nil[int](), nil
	}

	for e := range c.c {
		atomic.AddInt64(&c.countR, 1)
		err := f(int(c.countR), e)
		if err != nil {
			return wrap.Wrap(int(c.countR)), err
		}
	}

	return nil, nil
}

func (c *ChanRW[E]) IterFully(f func(index int, element E) error) (err api.MapAPI[int, error, any]) {
	if c == nil {
		return
	}

	for e := range c.c {
		atomic.AddInt64(&c.countR, 1)
		er := f(int(c.countR), e)
		if er != nil {
			if err == nil {
				err = MakeMap[int, error, any](0)
			}
			err.Set(int(c.countR), er)
		}
	}

	return
}
