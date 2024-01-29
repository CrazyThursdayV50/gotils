package ordered

type Ordered[E any] interface {
	Element() E
	Equal(E) bool
	LessThan(E) bool
}
