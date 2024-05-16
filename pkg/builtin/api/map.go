package api

import (
	"cmp"
)

type MapAPI[K cmp.Ordered | *T, V any, T any] interface {
	Len() int
	Has(k K) bool
	AddSoft(k K, v V)
	Del(k K)
	Keys() SliceAPI[K]
	Values() SliceAPI[V]
	Inner() map[K]V
	Clear()

	GetSeter[K, V]
	Iter[K, V, T]
	IterMut[K, V, T]
}
