package gmap

import (
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/models"
)

func From[K comparable, V any](m map[K]V) api.MapAPI[K, V] {
	return models.FromMap[K, V](m)
}

func Make[K comparable, V any](cap int) api.MapAPI[K, V] {
	return From(make(map[K]V, cap))
}
