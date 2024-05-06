package out

import (
	"github.com/CrazyThursdayV50/gotils/pkg/async/worker"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api"
	gchan "github.com/CrazyThursdayV50/gotils/pkg/builtin/api/chan"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api/slice"
)

type Fan[element any] struct {
	to    api.ChanAPI[element]
	close func()
}

func New[element any](count, buffer int, handler func(t element)) *Fan[element] {
	var fan Fan[element]
	fan.to = gchan.Make[element](buffer)
	closeFuncs := []func(){
		fan.to.Close,
	}

	slice.Empty(count).IterFuncFully(func(struct{}) {
		worker, _ := worker.New(handler)
		worker.WithGraceful(true)
		worker.WithTrigger(fan.to)
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

func (f *Fan[T]) To() api.ChanAPI[T] { return f.to }
func (f *Fan[T]) Close()             { f.close() }
func (f *Fan[T]) Send(t T)           { f.to.Send(t) }
