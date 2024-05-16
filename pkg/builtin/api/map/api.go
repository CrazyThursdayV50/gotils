package gmap

import (
	"cmp"

	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/models"
)

func From[K cmp.Ordered | *T, V any, T any](m map[K]V) api.MapAPI[K, V, T] {
	return models.FromMap[K, V, T](m)
}

func Make[K cmp.Ordered | *T, V any, T any](cap int) api.MapAPI[K, V, T] {
	return From[K, V, T](make(map[K]V, cap))
}
