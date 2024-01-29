package compare

import (
	"gotils/pkg/slice"
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
