package compare

import (
	"gotils/pkg/slice"
	"slices"
)

func (c Comparer[E, T]) Slice() slice.Slice[E] {
	return slice.Slice[E](c)
}

func (c Comparer[E, T]) Has(element E) (ok bool) {
	c.Slice().IterFunc(func(e E) bool {
		if e.Equal(element.Element()) {
			ok = true
			return false
		}
		return true
	})
	return
}

func (c Comparer[E, T]) Sort() {
	slices.SortFunc(c, func(a, b E) int {
		be := b.Element()

		if a.Equal(be) {
			return 0
		}

		if a.LessThan(be) {
			return -1
		}

		return 1
	})
}
