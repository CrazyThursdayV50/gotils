package order

import (
	"cmp"
	"gotils/pkg/ordered"
)

var _ ordered.Ordered[int] = New(0)

type Order[X cmp.Ordered] struct {
	x X
}
