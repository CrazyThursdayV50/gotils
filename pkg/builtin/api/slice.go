package api

import "github.com/CrazyThursdayV50/gotils/pkg/wrapper"

type SliceAPI[E any] interface {
	Len() int
	Swap(i int, j int)
	Less(i int, j int) bool
	WithLessFunc(f func(a, b E) bool)
	Append(elements ...E)
	Slice() []E
	Cut(from, to int) []E
	Index(element E, equal func(E, E) bool) wrapper.UnWrapper[int]
	Clear()

	GetSeter[int, E]
	Iter[int, E]
	IterMut[int, E]
}

// var _ SliceAPI[any] = (*models.Slice[any])(nil)
