package models

import (
	"github.com/CrazyThursdayV50/gotils/pkg/async/goo"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api"
	"github.com/CrazyThursdayV50/gotils/pkg/wrapper"
	"github.com/CrazyThursdayV50/gotils/pkg/wrapper/wrap"
)

var _ api.SliceAPI[any] = (*Slice[any])(nil)

type Slice[E any] struct {
	slice    []E
	lessFunc func(E, E) bool
}

func (s *Slice[E]) Unwrap() []E {
	if s == nil {
		return nil
	}
	return s.slice
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
	return cap(s.Unwrap())
}

func (s *Slice[E]) Len() int {
	if s == nil {
		return 0
	}
	return len(s.Unwrap())
}

func (s *Slice[E]) Swap(i, j int) {
	if s == nil {
		return
	}
	s.slice[i], s.slice[j] = s.Unwrap()[j], s.Unwrap()[i]
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
	s.slice = append(s.Unwrap(), elements...)
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
		return wrap.Nil[E]()
	}
	if s.Len() < index+1 {
		return wrap.Nil[E]()
	}
	return wrap.Wrap(s.Unwrap()[index])
}

func (s *Slice[E]) Clear() {
	if s == nil {
		return
	}
	clear(s.Unwrap())
}

func (s *Slice[E]) Chunk(len int) <-chan []E {
	var ch = FromChan(make(chan []E))
	if s.Len() == 0 || len == 0 {
		goo.Go(ch.Close)
		return ch.Unwrap()
	}

	var count = (s.Len()-1)/len + 1
	goo.Go(func() {
		defer ch.Close()
		for i := range count {
			from := i * len
			end := from + len
			if end > s.Len() {
				end = s.Len()
			}

			ch.Send(s.slice[from:end])
		}
	})

	return ch.Unwrap()
}

func (s *Slice[E]) IterOkay(f func(index int, element E) bool) wrapper.UnWrapper[int] {
	if s == nil {
		return wrap.Wrap(-1)
	}
	for i, e := range s.Unwrap() {
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
	for i, e := range s.Unwrap() {
		err := f(i, e)
		if err != nil {
			return wrap.Wrap(i), err
		}
	}
	return wrap.Wrap(s.Len()), nil
}

func (s *Slice[E]) IterFully(f func(index int, element E) error) (err api.MapAPI[int, error]) {
	if s == nil {
		return
	}

	for i, e := range s.Unwrap() {
		er := f(i, e)
		if er != nil {
			if err == nil {
				err = MakeMap[int, error](0)
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
	for i, e := range s.Unwrap() {
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
	for i, e := range s.Unwrap() {
		err := f(i, e, s)
		if err != nil {
			return wrap.Wrap(i), err
		}
	}
	return wrap.Wrap(s.Len()), nil
}

func (s *Slice[E]) IterMutFully(f func(index int, element E, self api.GetSeter[int, E]) error) (err api.MapAPI[int, error]) {
	if s == nil {
		return
	}
	for i, e := range s.Unwrap() {
		er := f(i, e, s)
		if er != nil {
			if err == nil {
				err = MakeMap[int, error](0)
			}
			err.Set(i, er)
		}
	}
	return
}
