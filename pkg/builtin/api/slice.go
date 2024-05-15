package api

import (
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/models"
	"github.com/CrazyThursdayV50/gotils/pkg/wrapper"
)

type SliceAPI[E any] interface {
	Len() int
	Swap(i int, j int)
	Less(i int, j int) bool
	WithLessFunc(f func(i int, j int) bool)
	Append(elements ...E)
	Slice() []E
	Cut(from, to int) []E
	Index(element E, equal func(E, E) bool) int
	Get(index int) wrapper.UnWrapper[E]

	IterIndex(func(index int) bool) int
	IterFunc(func(element E) bool)
	IterError(func(element E) error) error
	IterFuncMut(func(element E, self *models.Slice[E]) bool)
	IterIndexFully(f func(int))
	IterFuncFully(func(element E))
	IterFuncMutFully(func(element E, self *models.Slice[E]))
	IterErrorFully(f func(E) error) (err *models.Slice[error])
}

var _ SliceAPI[any] = (*models.Slice[any])(nil)
