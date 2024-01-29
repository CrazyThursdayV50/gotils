package order

import (
	"cmp"
	"gotils/pkg/slice/ordered"
)

var _ ordered.Ordered[int] = New(0)

type Order[X cmp.Ordered] struct {
	x X
}
