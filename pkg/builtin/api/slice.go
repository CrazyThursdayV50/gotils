package api

import "github.com/CrazyThursdayV50/gotils/pkg/builtin/models"

type SliceAPI[E any] interface {
	Len() int
	Swap(int, int)
	Less(int, int) bool
	WithLessFunc(f func(int, int) bool)
	Append(...E)
	Slice() []E

	IterFunc(func(E) bool)
	IterFuncMut(func(E, *models.Slice[E]) bool)
	IterFuncFully(func(E))
	IterFuncMutFully(func(E, *models.Slice[E]))
}

var _ SliceAPI[any] = (*models.Slice[any])(nil)
