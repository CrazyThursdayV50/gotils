package slice

import (
	"cmp"

	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/models"
)

func From[E any](sli []E) api.SliceAPI[E] {
	return models.FromSlice(sli)
}

func Collect[E any, T any](sli []E, collector func(E) T) api.SliceAPI[T] {
	src := models.FromSlice(sli)
	dst := models.MakeSlice[T](0, src.Len())
	src.IterFuncFully(func(e E) {
		dst.Append(collector(e))
	})
	return dst
}

func Map[E any, K cmp.Ordered, V any](sli []E, mapper func(E) (K, V)) *models.Map[K, V] {
	src := models.FromSlice(sli)
	dst := models.MakeMap[K, V](src.Len())
	src.IterFuncFully(func(e E) {
		k, v := mapper(e)
		dst.Add(k, v)
	})
	return dst
}
