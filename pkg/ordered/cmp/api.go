package cmp

import "cmp"

func New[X cmp.Ordered](x X) *Order[X] {
	return &Order[X]{x: x}
}

func Equal[X cmp.Ordered](a, b X) bool {
	return a == b
}
