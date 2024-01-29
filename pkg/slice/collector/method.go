package collector

import "gotils/pkg/slice"

func (c *Collector[F, E, T]) Collections() slice.Slice[T] {
	var collections = slice.Make[T](0, 0)
	c.IterFunc(func(e E) bool {
		collections.Append(c.collect(e))
		return true
	})
	return collections
}
