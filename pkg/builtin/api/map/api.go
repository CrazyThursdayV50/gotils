package gmap

import (
	"cmp"

	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/models"
)

func From[K cmp.Ordered, V any](m map[K]V) api.MapAPI[K, V] {
	return models.FromMap(m)
}
