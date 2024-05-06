package gchan

import (
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/models"
)

func From[E any](c chan E) api.ChanAPI[E] {
	return models.FromChan(c)
}

func FromRead[E any](c <-chan E) api.ChanAPIR[E] {
	return models.FromChanR(c)
}

func FromWrite[E any](c chan<- E) api.ChanAPIW[E] {
	return models.FromChanW(c)
}

func Make[E any](buff int) api.ChanAPI[E] {
	return From(make(chan E, buff))
}

func MakeRead[E any](buff int) api.ChanAPIR[E] {
	return FromRead(make(chan E, buff))
}

func MakeWrite[E any](buff int) api.ChanAPIW[E] {
	return FromWrite(make(chan E, buff))
}
