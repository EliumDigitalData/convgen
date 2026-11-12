package box

type Box[T any] struct {
	V T
}

type Wrap[T any] struct {
	V T
}

func Convert[T any](b Box[T]) Wrap[T] {
	return Wrap[T]{V: b.V}
}
