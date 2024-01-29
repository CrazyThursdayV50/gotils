package api

import (
	"cmp"

	"github.com/CrazyThursdayV50/gotils/pkg/builtin/models"
)

type MapAPI[K cmp.Ordered, V any] interface {
	Len() int
	Has(K) bool
	Add(K, V)
	AddSoft(K, V)
	Get(K) (V, bool)
	Del(K)
	Keys() *models.Slice[K]
	Values() *models.Slice[V]
	Map() map[K]V

	IterFunc(func(K, V) bool)
	IterFuncMut(func(K, V, *models.Map[K, V]) bool)
	IterFuncFully(func(K, V))
	IterFuncMutFully(func(K, V, *models.Map[K, V]))
}

var _ MapAPI[int, any] = (*models.Map[int, any])(nil)
