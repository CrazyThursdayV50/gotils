package api

import "github.com/CrazyThursdayV50/gotils/pkg/builtin/models"

type SliceAPI[E any] interface {
	Len() int
	Swap(i int, j int)
	Less(i int, j int) bool
	WithLessFunc(f func(i int, j int) bool)
	Append(elements ...E)
	Slice() []E

	IterFunc(func(element E) bool)
	IterFuncMut(func(element E, self *models.Slice[E]) bool)
	IterFuncFully(func(element E))
	IterFuncMutFully(func(element E, self *models.Slice[E]))
}

var _ SliceAPI[any] = (*models.Slice[any])(nil)
