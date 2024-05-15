package models

import (
	"github.com/CrazyThursdayV50/gotils/pkg/wrapper"
	"github.com/CrazyThursdayV50/gotils/pkg/wrapper/wrap"
)

type Slice[E any] struct {
	slice    []E
	lessFunc func(int, int) bool
}

func (s *Slice[E]) Slice() []E {
	if s == nil {
		return nil
	}
	return s.slice
}

func (s *Slice[E]) Cut(from, to int) []E {
	if s == nil {
		return nil
	}

	if from > s.Len()-1 || from < 0 {
		return nil
	}

	if to > s.Len()-1 || to < 0 {
		return nil
	}

	if from < to {
		return nil
	}

	return s.slice[from:to]
}

func (s *Slice[E]) Index(element E, equal func(E, E) bool) int {
	if s == nil {
		return -1
	}
	if equal == nil {
		return -1
	}

	var index = -1
	s.IterIndex(func(i int) bool {
		ok := equal(s.Get(i).Unwrap(), element)
		if ok {
			index = i
			return false
		}
		return true
	})
	return index
}

func FromSlice[E any](slice ...E) *Slice[E] {
	return &Slice[E]{
		slice: slice,
	}
}

func MakeSlice[E any](len, cap int) *Slice[E] {
	return FromSlice(make([]E, len, cap)...)
}

func (s *Slice[E]) Len() int {
	if s == nil {
		return 0
	}
	return len(s.slice)
}

func (s *Slice[E]) Swap(i, j int) {
	if s == nil {
		return
	}
	s.slice[i], s.slice[j] = s.slice[j], s.slice[i]
}

func (s *Slice[E]) Less(i, j int) bool {
	if s == nil {
		return false
	}
	if s.lessFunc == nil {
		return false
	}
	return s.lessFunc(i, j)
}

func (s *Slice[E]) WithLessFunc(f func(int, int) bool) {
	if s == nil {
		return
	}
	s.lessFunc = f
}

func (s *Slice[E]) Append(elements ...E) {
	if s == nil {
		return
	}
	s.slice = append(s.slice, elements...)
}

func (s *Slice[E]) Get(index int) wrapper.UnWrapper[E] {
	if s == nil {
		return nil
	}
	if s.Len() < index+1 {
		return nil
	}
	return wrap.Wrap(s.slice[index])
}

func (s *Slice[E]) IterIndex(f func(int) bool) int {
	if s == nil {
		return -1
	}
	for i := range s.slice {
		ok := f(i)
		if !ok {
			return i
		}
	}
	return s.Len()
}

func (s *Slice[E]) IterFunc(f func(E) bool) {
	if s == nil {
		return
	}
	for _, e := range s.slice {
		ok := f(e)
		if !ok {
			return
		}
	}
}

func (s *Slice[E]) IterFuncMut(f func(E, *Slice[E]) bool) {
	if s == nil {
		return
	}
	for _, e := range s.slice {
		ok := f(e, s)
		if !ok {
			return
		}
	}
}

func (s *Slice[E]) IterError(f func(E) error) error {
	if s == nil {
		return nil
	}
	for _, e := range s.slice {
		err := f(e)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Slice[E]) IterIndexFully(f func(int)) {
	if s == nil {
		return
	}
	for i := range s.slice {
		f(i)
	}
}

func (s *Slice[E]) IterFuncFully(f func(E)) {
	if s == nil {
		return
	}
	for _, e := range s.slice {
		f(e)
	}
}

func (s *Slice[E]) IterFuncMutFully(f func(E, *Slice[E])) {
	if s == nil {
		return
	}
	for _, e := range s.slice {
		f(e, s)
	}
}

func (s *Slice[E]) IterErrorFully(f func(E) error) (err *Slice[error]) {
	if s == nil {
		return nil
	}
	for _, e := range s.slice {
		er := f(e)
		if er != nil {
			if err == nil {
				err = FromSlice(err.Slice()...)
			}
			err.Append(er)
		}
	}
	return
}
