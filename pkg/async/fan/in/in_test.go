package in

import (
	"fmt"
	"testing"

	"github.com/CrazyThursdayV50/gotils/pkg/async/goo"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api"
	gchan "github.com/CrazyThursdayV50/gotils/pkg/builtin/api/chan"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api/slice"
)

func TestIn(t *testing.T) {
	chans := slice.Collect(slice.Make[struct{}](10, 10).Slice(), func(struct{}) chan any {
		return make(chan any)
	})

	from := slice.Collect(chans.Slice(), func(element chan any) api.ChanAPI[any] {
		return gchan.From(element)
	})

	from.IterFuncFully(func(element api.ChanAPI[any]) {
		goo.Go(func() {
			for i := range 10 {
				element.Send(i)
			}
		})
	})

	fan := From[any](func(t any) {
		fmt.Printf("receive: %v\n", t)
	}, chans.Slice()...)
	defer fan.Close()

	<-make(chan struct{})
}
