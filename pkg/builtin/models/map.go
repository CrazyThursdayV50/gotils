package models

import (
	"cmp"
	"sync"

	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api"
	"github.com/CrazyThursdayV50/gotils/pkg/wrapper"
	"github.com/CrazyThursdayV50/gotils/pkg/wrapper/wrap"
)

var _ api.MapAPI[*int, any, int] = (*Map[*int, any, int])(nil)

type Map[K cmp.Ordered | *T, V any, T any] struct {
	l *sync.RWMutex
	m map[K]V
}

func (m *Map[K, V, T]) Unwrap() map[K]V {
	if m == nil {
		return nil
	}
	return m.m
}

func MakeMap[K cmp.Ordered | *T, V any, T any](cap int) *Map[K, V, T] {
	return FromMap[K, V, T](make(map[K]V, cap))
}

func FromMap[K cmp.Ordered | *T, V any, T any](m map[K]V) *Map[K, V, T] {
	if m == nil {
		return MakeMap[K, V, T](0)
	}

	return &Map[K, V, T]{
		l: &sync.RWMutex{},
		m: m,
	}
}

func (m *Map[K, V, T]) Len() int {
	if m == nil {
		return 0
	}
	return len(m.m)
}

func (m *Map[K, V, T]) Set(k K, v V) {
	if m == nil {
		return
	}
	m.l.Lock()
	defer m.l.Unlock()
	m.m[k] = v
}

func (m *Map[K, V, T]) Has(k K) bool {
	if m == nil {
		return false
	}
	m.l.RLock()
	defer m.l.RUnlock()
	_, ok := m.m[k]
	return ok
}

// add if not exist
func (m *Map[K, V, T]) AddSoft(k K, v V) {
	if m == nil {
		return
	}

	if m.Has(k) {
		return
	}

	m.Set(k, v)
}

func (m *Map[K, V, T]) Get(k K) wrapper.UnWrapper[V] {
	if m == nil {
		return wrap.Nil[V]()
	}
	m.l.RLock()
	defer m.l.RUnlock()
	v, ok := m.m[k]
	if !ok {
		return wrap.Nil[V]()
	}
	return wrap.Wrap(v)
}

func (m *Map[K, V, T]) Del(k K) {
	if m == nil {
		return
	}
	m.l.Lock()
	defer m.l.Unlock()
	delete(m.m, k)
}

func (m *Map[K, V, T]) Keys() api.SliceAPI[K] {
	if m == nil {
		return nil
	}
	m.l.RLock()
	defer m.l.RUnlock()
	slice := MakeSlice[K](0, m.Len())
	for k := range m.m {
		slice.Append(k)
	}

	return slice
}

func (m *Map[K, V, T]) Values() api.SliceAPI[V] {
	if m == nil {
		return nil
	}
	m.l.RLock()
	defer m.l.RUnlock()
	slice := MakeSlice[V](0, m.Len())
	for _, v := range m.m {
		slice.Append(v)
	}

	return slice
}

func (m *Map[K, V, T]) Clear() {
	if m == nil {
		return
	}
	m.l.Lock()
	defer m.l.Unlock()
	clear(m.m)
}

func (m *Map[K, V, T]) IterOkay(f func(k K, v V) bool) wrapper.UnWrapper[K] {
	if m == nil {
		return wrap.Nil[K]()
	}
	m.l.RLock()
	defer m.l.RUnlock()
	for k, v := range m.m {
		ok := f(k, v)
		if !ok {
			return wrap.Wrap(k)
		}
	}
	return wrap.Nil[K]()
}

func (m *Map[K, V, T]) IterError(f func(k K, v V) error) (wrapper.UnWrapper[K], error) {
	if m == nil {
		return wrap.Nil[K](), nil
	}
	m.l.RLock()
	defer m.l.RUnlock()
	for k, v := range m.m {
		err := f(k, v)
		if err != nil {
			return wrap.Wrap(k), err
		}
	}
	return wrap.Nil[K](), nil
}

func (m *Map[K, V, T]) IterFully(f func(k K, v V) error) (err api.MapAPI[K, error, T]) {
	if m == nil {
		return
	}
	m.l.RLock()
	defer m.l.RUnlock()
	for k, v := range m.m {
		er := f(k, v)
		if er != nil {
			if err == nil {
				err = MakeMap[K, error, T](0)
			}
			err.Set(k, er)
		}
	}
	return
}

func (m *Map[K, V, T]) IterMutOkay(f func(k K, v V, self api.GetSeter[K, V]) bool) wrapper.UnWrapper[K] {
	if m == nil {
		return wrap.Nil[K]()
	}

	keys := m.Keys()
	index := keys.IterOkay(func(_ int, element K) bool {
		return f(element, m.Get(element).Unwrap(), m)
	})

	if index == nil {
		return wrap.Nil[K]()
	}

	return keys.Get(index.Unwrap())
}

func (m *Map[K, V, T]) IterMutError(f func(k K, v V, self api.GetSeter[K, V]) error) (wrapper.UnWrapper[K], error) {
	if m == nil {
		return wrap.Nil[K](), nil
	}

	keys := m.Keys()
	index, err := keys.IterError(func(_ int, element K) error {
		return f(element, m.Get(element).Unwrap(), m)
	})

	if err == nil {
		return wrap.Nil[K](), nil
	}

	return keys.Get(index.Unwrap()), err
}

func (m *Map[K, V, T]) IterMutFully(f func(k K, v V, self api.GetSeter[K, V]) error) (err api.MapAPI[K, error, T]) {
	if m == nil {
		return
	}

	keys := m.Keys()
	errs := keys.IterFully(func(_ int, element K) error {
		return f(element, m.Get(element).Unwrap(), m)
	})

	if errs == nil {
		return nil
	}

	err = MakeMap[K, error, T](0)
	_ = errs.IterFully(func(k int, v error) error {
		err.Set(keys.Get(k).Unwrap(), v)
		return nil
	})

	return
}
