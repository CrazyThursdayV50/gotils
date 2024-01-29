package order

import "cmp"

func New[X cmp.Ordered](x X) *Order[X] {
	return &Order[X]{x: x}
}
