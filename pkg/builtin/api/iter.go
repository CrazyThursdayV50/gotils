package api

import (
	"cmp"

	"github.com/CrazyThursdayV50/gotils/pkg/wrapper"
)

type Iter[K cmp.Ordered | *T, V any, T any] interface {
	IterOkay(func(k K, v V) bool) wrapper.UnWrapper[K]
	IterError(func(k K, v V) error) (wrapper.UnWrapper[K], error)
	IterFully(func(k K, v V) error) MapAPI[K, error, T]
}

type IterMut[K cmp.Ordered | *T, V any, T any] interface {
	IterMutOkay(func(k K, v V, s GetSeter[K, V]) bool) wrapper.UnWrapper[K]
	IterMutError(func(k K, v V, s GetSeter[K, V]) error) (wrapper.UnWrapper[K], error)
	IterMutFully(func(k K, v V, s GetSeter[K, V]) error) MapAPI[K, error, T]
}
