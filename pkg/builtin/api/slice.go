package api

import "github.com/CrazyThursdayV50/gotils/pkg/wrapper"

type SliceAPI[E any] interface {
	Cap() int
	Len() int
	Swap(i int, j int)
	Less(i int, j int) bool
	WithLessFunc(f func(a, b E) bool)
	Append(elements ...E)
	Cut(from, to int) []E
	Index(element E, equal func(E, E) bool) SliceAPI[int]
	Del(index int)
	Clear()

	wrapper.UnWrapper[[]E]
	GetSeter[int, E]
	Iter[int, E, any]
	IterMut[int, E, any]
}
