package wrapper

type UnWrapper[T any] interface {
	Unwrap() T
}
