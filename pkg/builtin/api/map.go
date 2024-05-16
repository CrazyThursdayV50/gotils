package api

import (
	"cmp"
)

type MapAPI[K cmp.Ordered, V any] interface {
	Len() int
	Has(k K) bool
	AddSoft(k K, v V)
	Del(k K)
	Keys() SliceAPI[K]
	Values() SliceAPI[V]
	Map() map[K]V
	Clear()

	GetSeter[K, V]
	Iter[K, V]
	IterMut[K, V]
}
