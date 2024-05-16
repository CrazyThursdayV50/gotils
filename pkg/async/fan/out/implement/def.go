package implement

import (
	"github.com/CrazyThursdayV50/gotils/pkg/async/worker"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api"
	gchan "github.com/CrazyThursdayV50/gotils/pkg/builtin/api/chan"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api/slice"
)

type Fan[T any] struct {
	in    api.ChanAPI[T]
	close func()
}

func New[T any](count, buffer int, handler func(t T)) *Fan[T] {
	var fan Fan[T]
	fan.in = gchan.Make[T](buffer)
	closeFuncs := slice.From(fan.in.Close)

	slice.Make[struct{}](count, count).IterFully(func(int, struct{}) error {
		worker, _ := worker.New(handler)
		worker.WithGraceful(true)
		worker.WithTrigger(fan.in)
		worker.Run()
		closeFuncs.Append(worker.Stop)
		return nil
	})

	fan.close = func() {
		_ = closeFuncs.IterFully(func(_ int, f func()) error {
			f()
			return nil
		})
	}
	return &fan
}

func (f *Fan[T]) To() api.ChanAPI[T] { return f.in }
func (f *Fan[T]) Close()             { f.close() }
func (f *Fan[T]) Send(t T)           { f.in.Send(t) }
