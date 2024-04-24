package gmap

import (
	"cmp"

	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/models"
)

func From[K cmp.Ordered, V any](m map[K]V) api.MapAPI[K, V] {
	return models.FromMap(m)
}

func Make[K cmp.Ordered, V any](cap int) api.MapAPI[K, V] {
	return From(make(map[K]V, cap))
}
