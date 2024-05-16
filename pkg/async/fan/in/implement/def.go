package implement

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

	return New(handler, from.Inner()...)
}

func New[element any](handler func(t element), from ...api.ChanAPIR[element]) *Fan[element] {
	var fan Fan[element]
	worker, delivery := worker.New(handler)
	worker.WithGraceful(true)
	worker.Run()
	_ = slice.From(from...).IterFully(func(_ int, ch api.ChanAPIR[element]) error {
		goo.Go(func() {
			_ = ch.IterFully(func(_ int, element element) error {
				delivery(element)
				return nil
			})
		})
		return nil
	})
	fan.close = worker.Stop
	return &fan
}

func (f *Fan[T]) Close() { f.close() }
