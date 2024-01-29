package gchan

import (
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/models"
)

func From[E any](c chan E) api.ChanAPI[E] {
	return models.FromChan(c)
}

func New[E any](buff int) api.ChanAPI[E] {
	return From(make(chan E, buff))
}
