package collector

import "gotils/pkg/slice"

func From[F CollectFunc[E, T], E slice.Element, T any](sli []E, f func(E) T) Collector[F, E, T] {
	return Collector[F, E, T]{
		Slice:   slice.From(sli),
		collect: f,
	}
}
