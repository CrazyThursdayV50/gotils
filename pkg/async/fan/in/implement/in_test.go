package implement

import (
	"fmt"
	"testing"

	"github.com/CrazyThursdayV50/gotils/pkg/async/goo"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api"
	gchan "github.com/CrazyThursdayV50/gotils/pkg/builtin/api/chan"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api/slice"
)

func TestIn(t *testing.T) {
	chans := slice.Collect(slice.Empty(10).Unwrap(), func(struct{}) chan any {
		return make(chan any)
	})

	from := slice.Collect(chans.Unwrap(), func(element chan any) api.ChanAPI[any] {
		return gchan.From(element)
	})

	_ = from.IterFully(func(_ int, element api.ChanAPI[any]) error {
		goo.Go(func() {
			for i := range 10 {
				element.Send(i)
			}
		})
		return nil
	})

	fan := From[any](func(t any) {
		fmt.Printf("receive: %v\n", t)
	}, chans.Unwrap()...)
	defer fan.Close()

	<-make(chan struct{})
}
