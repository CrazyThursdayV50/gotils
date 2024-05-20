package api

import (
	"cmp"

	"github.com/CrazyThursdayV50/gotils/pkg/wrapper"
)

type MapAPI[K cmp.Ordered | *T, V any, T any] interface {
	Len() int
	Has(k K) bool
	AddSoft(k K, v V)
	Del(k K)
	Keys() SliceAPI[K]
	Values() SliceAPI[V]
	Clear()

	wrapper.UnWrapper[map[K]V]
	GetSeter[K, V]
	Iter[K, V, T]
	IterMut[K, V, T]
}
