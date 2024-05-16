package ordered

type Ordered[E any] interface {
	Element() E
	Equal(E) bool
	LessThan(E) bool
}

type CmpOrdered[T any] interface {
	Less(T) bool
}
