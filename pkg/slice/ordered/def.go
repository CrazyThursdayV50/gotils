package ordered

import (
	"gotils/pkg/slice"
)

type Ordered[E slice.Element] interface {
	Element() E
	Equal(E) bool
	LessThan(E) bool
}
