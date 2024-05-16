package gmap

import (
	"cmp"

	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/models"
)

func From[K cmp.Ordered, V any](m map[K]V) api.MapAPI[K, V, any] {
	return models.FromMap[K, V, any](m)
}

func Make[K cmp.Ordered, V any](cap int) api.MapAPI[K, V, any] {
	return From(make(map[K]V, cap))
}

func FromP[K any, V any](m map[*K]V) api.MapAPI[*K, V, K] {
	return models.FromMap[*K, V, K](m)
}

func MakeP[K any, V any](cap int) api.MapAPI[*K, V, K] {
	return FromP(make(map[*K]V, cap))
}
