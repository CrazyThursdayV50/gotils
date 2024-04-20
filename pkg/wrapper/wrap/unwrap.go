package wrap

type unwrapper[T any] struct {
	t T
}

func Wrap[T any](t T) *unwrapper[T] {
	return &unwrapper[T]{t: t}
}

func (w *unwrapper[T]) Unwrap() T { return w.t }
