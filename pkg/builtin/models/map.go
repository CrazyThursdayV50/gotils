package models

import (
	"cmp"
	"sync"
)

type Map[K cmp.Ordered, V any] struct {
	l *sync.RWMutex
	m map[K]V
}

func (m *Map[K, V]) Map() map[K]V {
	return m.m
}

func FromMap[K cmp.Ordered, V any](m map[K]V) *Map[K, V] {
	return &Map[K, V]{
		l: &sync.RWMutex{},
		m: m,
	}
}

func MakeMap[K cmp.Ordered, V any](cap int) *Map[K, V] {
	return FromMap(make(map[K]V, cap))
}

func (m *Map[K, V]) Len() int {
	return len(m.m)
}

func (m *Map[K, V]) Add(k K, v V) {
	m.l.Lock()
	defer m.l.Unlock()
	m.m[k] = v
}

func (m *Map[K, V]) Has(k K) bool {
	m.l.RLock()
	defer m.l.RUnlock()
	_, ok := m.m[k]
	return ok
}

// add if not exist
func (m *Map[K, V]) AddSoft(k K, v V) {
	if m.Has(k) {
		return
	}

	m.Add(k, v)
}

func (m *Map[K, V]) Get(k K) (V, bool) {
	m.l.RLock()
	defer m.l.RUnlock()
	v, ok := m.m[k]
	return v, ok
}

func (m *Map[K, V]) Del(k K) {
	m.l.Lock()
	defer m.l.Unlock()
	delete(m.m, k)
}

func (m *Map[K, V]) IterFunc(f func(k K, v V) bool) {
	for k, v := range m.m {
		ok := f(k, v)
		if !ok {
			return
		}
	}
}

func (m *Map[K, V]) IterFuncMut(f func(k K, v V, m *Map[K, V]) bool) {
	for k, v := range m.m {
		ok := f(k, v, m)
		if !ok {
			return
		}
	}
}

func (m *Map[K, V]) IterFuncFully(f func(k K, v V)) {
	for k, v := range m.m {
		f(k, v)
	}
}

func (m *Map[K, V]) IterFuncMutFully(f func(k K, v V, m *Map[K, V])) {
	for k, v := range m.m {
		f(k, v, m)
	}
}

func (m *Map[K, V]) Keys() *Slice[K] {
	m.l.RLock()
	defer m.l.RUnlock()
	slice := MakeSlice[K](0, m.Len())
	for k := range m.m {
		slice.Append(k)
	}

	return slice
}

func (m *Map[K, V]) Values() *Slice[V] {
	m.l.RLock()
	defer m.l.RUnlock()
	slice := MakeSlice[V](0, m.Len())
	for _, v := range m.m {
		slice.Append(v)
	}

	return slice
}
