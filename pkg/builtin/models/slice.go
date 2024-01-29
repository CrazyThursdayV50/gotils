package models

type Slice[E any] struct {
	slice    []E
	lessFunc func(int, int) bool
}

func (s *Slice[E]) Slice() []E { return s.slice }

func FromSlice[E any](slice []E) *Slice[E] {
	return &Slice[E]{
		slice: slice,
	}
}

func MakeSlice[E any](len, cap int) *Slice[E] {
	return FromSlice(make([]E, len, cap))
}

func (s *Slice[E]) Len() int {
	return len(s.slice)
}

func (s *Slice[E]) Swap(i, j int) {
	s.slice[i], s.slice[j] = s.slice[j], s.slice[i]
}

func (s *Slice[E]) Less(i, j int) bool {
	if s.lessFunc == nil {
		return false
	}
	return s.lessFunc(i, j)
}

func (s *Slice[E]) WithLessFunc(f func(int, int) bool) {
	s.lessFunc = f
}

func (s *Slice[E]) Append(elements ...E) {
	if s == nil {
		return
	}
	s.slice = append(s.slice, elements...)
}

func (s *Slice[E]) IterFunc(f func(E) bool) {
	for _, e := range s.slice {
		ok := f(e)
		if !ok {
			return
		}
	}
}

func (s *Slice[E]) IterFuncMut(f func(E, *Slice[E]) bool) {
	for _, e := range s.slice {
		ok := f(e, s)
		if !ok {
			return
		}
	}
}

func (s *Slice[E]) IterFuncFully(f func(E)) {
	for _, e := range s.slice {
		f(e)
	}
}

func (s *Slice[E]) IterFuncMutFully(f func(E, *Slice[E])) {
	for _, e := range s.slice {
		f(e, s)
	}
}
