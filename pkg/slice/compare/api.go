package compare

import (
	"cmp"
	"gotils/pkg/ordered"
	"gotils/pkg/ordered/anyorder"
	"gotils/pkg/ordered/order"
	"gotils/pkg/slice"
)

func From[E ordered.Ordered[T], T slice.Element](sli []E) Comparer[E, T] {
	return Comparer[E, T](slice.From(sli))
}

func FromFunc[E ordered.Ordered[T], T slice.Element](sli []T, f func(T) E) Comparer[E, T] {
	s := slice.From(sli)
	var orderedSlice = slice.Make[E](0, s.Len())
	s.IterFuncFully(func(element T) { orderedSlice.Append(f(element)) })
	return From(orderedSlice)
}

func FromOrdered[T cmp.Ordered](sli []T) Comparer[*order.Order[T], T] {
	return FromFunc(sli, order.New)
}

func FromAny[T slice.Element](sli []T, cmp func(a, b T) int) Comparer[*anyorder.Order[T], T] {
	return FromFunc(sli, func(a T) *anyorder.Order[T] {
		return anyorder.New(a, func(b T) int { return cmp(a, b) })
	})
}
