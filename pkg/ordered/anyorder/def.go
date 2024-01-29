package anyorder

import (
	"fmt"

	"github.com/CrazyThursdayV50/gotils/pkg/ordered"
)

var _ ordered.Ordered[int] = New(0, func(i int) int {
	if i == 0 {
		return 0
	}
	if i > 0 {
		return 1
	}
	return -1
})

type Order[T any] struct {
	x   T
	cmp func(T) int
}

func (a *Order[T]) Element() T {
	return a.x
}

func (a *Order[T]) Equal(x T) bool {
	return a.cmp(x) == 0
}

func (a *Order[T]) LessThan(x T) bool {
	return a.cmp(x) < 0
}

func (a *Order[T]) String() string {
	return fmt.Sprintf("%v", a.Element())
}

func New[T any](x T, cmp func(T) int) *Order[T] {
	return &Order[T]{x, cmp}
}
