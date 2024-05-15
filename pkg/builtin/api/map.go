package api

import (
	"cmp"

	"github.com/CrazyThursdayV50/gotils/pkg/builtin/models"
	"github.com/CrazyThursdayV50/gotils/pkg/wrapper"
)

type MapAPI[K cmp.Ordered, V any] interface {
	Len() int
	Has(k K) bool
	Add(k K, v V)
	AddSoft(k K, v V)
	Get(k K) wrapper.UnWrapper[V]
	Del(k K)
	Keys() *models.Slice[K]
	Values() *models.Slice[V]
	Map() map[K]V

	IterFunc(func(k K, v V) bool)
	IterError(func(k K, v V) error) error
	IterFuncMut(func(k K, v V, m *models.Map[K, V]) bool)
	IterFuncFully(func(k K, v V))
	IterFuncMutFully(func(k K, v V, m *models.Map[K, V]))
	IterErrorFully(func(k K, v V) error) *models.Map[K, error]
}

var _ MapAPI[int, any] = (*models.Map[int, any])(nil)
