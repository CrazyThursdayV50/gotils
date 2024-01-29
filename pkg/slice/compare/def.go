package compare

import (
	"gotils/pkg/ordered"
	"gotils/pkg/slice"
)

type Comparer[E ordered.Ordered[T], T slice.Element] slice.Slice[E]
