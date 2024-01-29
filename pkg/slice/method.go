package slice

func (s Slice[E]) Len() int {
	return len(s)
}

func (s Slice[E]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *Slice[E]) Append(elements ...E) {
	if s == nil {
		return
	}
	*s = append(*s, elements...)
}

func (s Slice[E]) IterFunc(f func(E) bool) {
	for _, e := range s {
		ok := f(e)
		if !ok {
			return
		}
	}
}

func (s Slice[E]) IterFuncFully(f func(E)) {
	for _, e := range s {
		f(e)
	}
}
