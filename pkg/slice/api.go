package slice

func From[E Element](slice []E) Slice[E] {
	return Slice[E](slice)
}

func Make[E Element](len, cap int) Slice[E] {
	return make([]E, len, cap)
}
