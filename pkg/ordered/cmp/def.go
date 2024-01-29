package cmp

import (
	"cmp"

	"github.com/CrazyThursdayV50/gotils/pkg/ordered"
)

var _ ordered.Ordered[int] = New(0)

type Order[X cmp.Ordered] struct {
	x X
}
