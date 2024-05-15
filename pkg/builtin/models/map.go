package models

import (
	"cmp"
	"sync"

	"github.com/CrazyThursdayV50/gotils/pkg/wrapper"
	"github.com/CrazyThursdayV50/gotils/pkg/wrapper/wrap"
)

type Map[K cmp.Ordered, V any] struct {
	l *sync.RWMutex
	m map[K]V
}

func (m *Map[K, V]) Map() map[K]V {
	if m == nil {
		return nil
	}
	return m.m
}

func FromMap[K cmp.Ordered, V any](m map[K]V) *Map[K, V] {
	if m == nil {
		return MakeMap[K, V](0)
	}

	return &Map[K, V]{
		l: &sync.RWMutex{},
		m: m,
	}
}

func MakeMap[K cmp.Ordered, V any](cap int) *Map[K, V] {
	return FromMap(make(map[K]V, cap))
}

func (m *Map[K, V]) Len() int {
	if m == nil {
		return 0
	}
	return len(m.m)
}

func (m *Map[K, V]) Add(k K, v V) {
	if m == nil {
		return
	}
	m.l.Lock()
	defer m.l.Unlock()
	m.m[k] = v
}

func (m *Map[K, V]) Has(k K) bool {
	if m == nil {
		return false
	}
	m.l.RLock()
	defer m.l.RUnlock()
	_, ok := m.m[k]
	return ok
}

// add if not exist
func (m *Map[K, V]) AddSoft(k K, v V) {
	if m == nil {
		return
	}

	if m.Has(k) {
		return
	}

	m.Add(k, v)
}

func (m *Map[K, V]) Get(k K) wrapper.UnWrapper[V] {
	if m == nil {
		return nil
	}
	m.l.RLock()
	defer m.l.RUnlock()
	v, ok := m.m[k]
	if !ok {
		return nil
	}
	return wrap.Wrap(v)
}

func (m *Map[K, V]) Del(k K) {
	if m == nil {
		return
	}
	m.l.Lock()
	defer m.l.Unlock()
	delete(m.m, k)
}

func (m *Map[K, V]) IterFunc(f func(k K, v V) bool) {
	if m == nil {
		return
	}
	m.l.RLock()
	defer m.l.RUnlock()
	for k, v := range m.m {
		ok := f(k, v)
		if !ok {
			return
		}
	}
}

func (m *Map[K, V]) IterError(f func(k K, v V) error) error {
	if m == nil {
		return nil
	}
	m.l.RLock()
	defer m.l.RUnlock()
	for k, v := range m.m {
		err := f(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Map[K, V]) IterFuncMut(f func(k K, v V, m *Map[K, V]) bool) {
	if m == nil {
		return
	}
	m.l.Lock()
	defer m.l.Unlock()
	for k, v := range m.m {
		ok := f(k, v, m)
		if !ok {
			return
		}
	}
}

func (m *Map[K, V]) IterFuncFully(f func(k K, v V)) {
	if m == nil {
		return
	}
	m.l.RLock()
	defer m.l.RUnlock()
	for k, v := range m.m {
		f(k, v)
	}
}

func (m *Map[K, V]) IterFuncMutFully(f func(k K, v V, m *Map[K, V])) {
	if m == nil {
		return
	}
	m.l.Lock()
	defer m.l.Unlock()
	for k, v := range m.m {
		f(k, v, m)
	}
}

func (m *Map[K, V]) IterErrorFully(f func(k K, v V) error) (err *Map[K, error]) {
	if m == nil {
		return
	}
	m.l.RLock()
	defer m.l.RUnlock()
	for k, v := range m.m {
		er := f(k, v)
		if er == nil {
			continue
		}
		if err == nil {
			err = FromMap(err.Map())
		}
		err.Add(k, er)
	}

	return
}

func (m *Map[K, V]) Keys() *Slice[K] {
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

func (m *Map[K, V]) Values() *Slice[V] {
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
