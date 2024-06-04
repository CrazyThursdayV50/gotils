package api

import "github.com/CrazyThursdayV50/gotils/pkg/wrapper"

type SliceAPI[E any] interface {
	Cap() int
	Len() int
	Swap(i int, j int)
	Less(i int, j int) bool
	WithLessFunc(f func(a, b E) bool)
	Append(elements ...E)
	Clear()
	Chunk(len int) <-chan []E
	// Clone() []E
	// Compact(f func(a, b E) bool)
	// Contains(f func(element E) bool) bool
	// Index(f func(element E) bool) int
	// Del(index int)
	// DelFunc(f func(element E) bool)
	// Cut(from, to int) []E

	wrapper.UnWrapper[[]E]
	GetSeter[int, E]
	Iter[int, E, any]
	IterMut[int, E, any]
}
