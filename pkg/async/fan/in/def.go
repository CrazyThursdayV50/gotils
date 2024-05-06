package in

import (
	"github.com/CrazyThursdayV50/gotils/pkg/async/goo"
	"github.com/CrazyThursdayV50/gotils/pkg/async/worker"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api"
	gchan "github.com/CrazyThursdayV50/gotils/pkg/builtin/api/chan"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api/slice"
)

type Fan[T any] struct {
	close func()
}

type c[T any] interface {
	chan T | <-chan T
}

func From[T any, C c[T]](handler func(t T), chans ...C) *Fan[T] {
	from := slice.Collect(chans, func(element C) api.ChanAPIR[T] {
		return gchan.FromRead(element)
	})

	return New(handler, from.Slice()...)
}

func New[T any](handler func(t T), from ...api.ChanAPIR[T]) *Fan[T] {
	var fan Fan[T]
	worker, delivery := worker.New(handler)
	worker.WithGraceful(true)
	worker.Run()
	slice.From(from).IterFuncFully(func(ch api.ChanAPIR[T]) {
		goo.Go(func() {
			ch.IterFunc(func(element T) bool {
				delivery(element)
				return true
			})
		})
	})
	fan.close = worker.Stop
	return &fan
}

func (f *Fan[T]) Close() { f.close() }
