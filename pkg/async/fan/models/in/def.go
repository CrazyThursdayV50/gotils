package in

import (
	"github.com/CrazyThursdayV50/gotils/pkg/async/goo"
	"github.com/CrazyThursdayV50/gotils/pkg/async/worker"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api"
	gchan "github.com/CrazyThursdayV50/gotils/pkg/builtin/api/chan"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api/slice"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/models"
)

type Fan[T any] struct {
	close func()
}

func From[element any, C models.ChanRead[element]](handler func(t element), chans ...C) *Fan[element] {
	from := slice.Collect(chans, func(element C) api.ChanAPIR[element] {
		return gchan.FromRead(element)
	})

	return New(handler, from.Slice()...)
}

func New[element any](handler func(t element), from ...api.ChanAPIR[element]) *Fan[element] {
	var fan Fan[element]
	worker, delivery := worker.New(handler)
	worker.WithGraceful(true)
	worker.Run()
	slice.From(from).IterFuncFully(func(ch api.ChanAPIR[element]) {
		goo.Go(func() {
			ch.IterFunc(func(element element) bool {
				delivery(element)
				return true
			})
		})
	})
	fan.close = worker.Stop
	return &fan
}

func (f *Fan[T]) Close() { f.close() }
