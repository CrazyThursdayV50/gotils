package models

import (
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api"
	"github.com/CrazyThursdayV50/gotils/pkg/wrapper"
	"github.com/CrazyThursdayV50/gotils/pkg/wrapper/wrap"
)

var _ api.SliceAPI[any] = (*Slice[any])(nil)

type Slice[E any] struct {
	slice    []E
	lessFunc func(E, E) bool
}

func (s *Slice[E]) Inner() []E {
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

func (s *Slice[E]) Index(element E, equal func(E, E) bool) wrapper.UnWrapper[int] {
	if s == nil {
		return nil
	}
	if equal == nil {
		return nil
	}

	return s.IterOkay(func(_ int, e E) bool {
		return !equal(e, element)
	})
}

func FromSlice[E any](slice ...E) *Slice[E] {
	return &Slice[E]{
		slice: slice,
	}
}

func MakeSlice[E any](len, cap int) *Slice[E] {
	return FromSlice(make([]E, len, cap)...)
}

func (s *Slice[E]) Cap() int {
	if s == nil {
		return 0
	}
	return cap(s.slice)
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
	ie := s.Get(i)
	je := s.Get(j)
	if ie == nil || je == nil {
		return false
	}
	return s.lessFunc(ie.Unwrap(), je.Unwrap())
}

func (s *Slice[E]) WithLessFunc(f func(a, b E) bool) {
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

func (s *Slice[E]) Set(index int, element E) {
	if s == nil {
		return
	}
	if s.Len() < index+1 {
		return
	}
	s.slice[index] = element
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

func (s *Slice[E]) Clear() {
	if s == nil {
		return
	}
	clear(s.slice)
}

func (s *Slice[E]) IterOkay(f func(index int, element E) bool) wrapper.UnWrapper[int] {
	if s == nil {
		return wrap.Wrap(-1)
	}
	for i, e := range s.slice {
		ok := f(i, e)
		if !ok {
			return wrap.Wrap(i)
		}
	}
	return wrap.Wrap(s.Len())
}

func (s *Slice[E]) IterError(f func(index int, element E) error) (wrapper.UnWrapper[int], error) {
	if s == nil {
		return wrap.Wrap(-1), nil
	}
	for i, e := range s.slice {
		err := f(i, e)
		if err != nil {
			return wrap.Wrap(i), err
		}
	}
	return wrap.Wrap(s.Len()), nil
}

func (s *Slice[E]) IterFully(f func(index int, element E) error) (err api.MapAPI[int, error, any]) {
	if s == nil {
		return
	}

	for i, e := range s.slice {
		er := f(i, e)
		if er != nil {
			if err == nil {
				err = MakeMap[int, error, any](0)
			}
			err.Set(i, er)
		}
	}
	return
}

func (s *Slice[E]) IterMutOkay(f func(index int, element E, self api.GetSeter[int, E]) bool) wrapper.UnWrapper[int] {
	if s == nil {
		return wrap.Wrap(-1)
	}
	for i, e := range s.slice {
		ok := f(i, e, s)
		if !ok {
			return wrap.Wrap(i)
		}
	}
	return wrap.Wrap(s.Len())
}

func (s *Slice[E]) IterMutError(f func(index int, element E, self api.GetSeter[int, E]) error) (wrapper.UnWrapper[int], error) {
	if s == nil {
		return wrap.Wrap(-1), nil
	}
	for i, e := range s.slice {
		err := f(i, e, s)
		if err != nil {
			return wrap.Wrap(i), err
		}
	}
	return wrap.Wrap(s.Len()), nil
}

func (s *Slice[E]) IterMutFully(f func(index int, element E, self api.GetSeter[int, E]) error) (err api.MapAPI[int, error, any]) {
	if s == nil {
		return
	}
	for i, e := range s.slice {
		er := f(i, e, s)
		if er != nil {
			if err == nil {
				err = MakeMap[int, error, any](0)
			}
			err.Set(i, er)
		}
	}
	return
}
