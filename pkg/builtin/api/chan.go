package api

import "github.com/CrazyThursdayV50/gotils/pkg/builtin/models"

type ChanAPI[E any] interface {
	Len() int
	IsEmpty() bool
	Closed() bool
	Close()
	Receive() E
	Send(E)
	Chan() chan E
	Renew(int)
	RenewForce(int)

	IterFunc(func(E) bool)
	IterFuncMut(func(E, *models.Chan[E]) bool)
	IterFuncFully(func(E))
	IterFuncMutFully(func(E, *models.Chan[E]))
}

var _ ChanAPI[any] = (*models.Chan[any])(nil)
