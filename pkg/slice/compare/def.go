package compare

import (
	"gotils/pkg/slice"
	"gotils/pkg/slice/order"
)

type Comparer[E order.Ordered[T], T slice.Element] slice.Slice[E]
