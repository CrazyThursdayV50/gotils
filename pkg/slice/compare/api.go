package compare

import (
	"cmp"
	"gotils/pkg/slice"
	"gotils/pkg/slice/ordered"
	"gotils/pkg/slice/ordered/order"
)

func From[E ordered.Ordered[T], T slice.Element](sli []E) Comparer[E, T] {
	return Comparer[E, T](slice.From(sli))
}

func FromFunc[E ordered.Ordered[T], T cmp.Ordered](sli []T, f func(T) E) Comparer[E, T] {
	s := slice.From(sli)
	var orderedSlice = slice.Make[E](0, s.Len())
	s.IterFunc(func(element T) bool {
		orderedSlice.Append(f(element))
		return true
	})
	return Comparer[E, T](orderedSlice)
}

func FromOrdered[T cmp.Ordered](sli []T) Comparer[*order.Order[T], T] {
	return FromFunc(sli, order.New)
}
