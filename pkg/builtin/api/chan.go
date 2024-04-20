package api

import (
	"time"

	"github.com/CrazyThursdayV50/gotils/pkg/builtin/models"
	"github.com/CrazyThursdayV50/gotils/pkg/wrapper"
)

type ChanAPI[E any] interface {
	Len() int
	IsEmpty() bool
	Closed() bool
	Close()
	Receive() (wrapper.UnWrapper[E], bool)
	Send(E)
	Chan() chan E
	Timeout(send, recv time.Duration)
	Renew(int)
	RenewForce(int)

	IterFunc(func(E) bool)
	IterFuncMut(func(E, *models.Chan[E]) bool)
	IterFuncFully(func(E))
	IterFuncMutFully(func(E, *models.Chan[E]))
}

var _ ChanAPI[any] = (*models.Chan[any])(nil)
