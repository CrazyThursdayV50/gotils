package out

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
	closeFuncs := []func(){
		fan.in.Close,
	}

	slice.Make[struct{}](count, count).IterFuncFully(func(struct{}) {
		worker, _ := worker.New(handler)
		worker.WithGraceful(true)
		worker.WithTrigger(fan.in)
		worker.Run()
		closeFuncs = append(closeFuncs, worker.Stop)
	})

	fan.close = func() {
		slice.From(closeFuncs).IterFuncFully(func(f func()) {
			f()
		})
	}
	return &fan
}

func (f *Fan[T]) In() api.ChanAPI[T] { return f.in }
func (f *Fan[T]) Close()             { f.close() }
func (f *Fan[T]) Send(t T)           { f.in.Send(t) }
