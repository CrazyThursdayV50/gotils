package collector

import "gotils/pkg/slice"

type CollectFunc[E slice.Element, T any] func(E) T

type Collector[F CollectFunc[E, T], E slice.Element, T any] struct {
	slice.Slice[E]
	collect F
}
