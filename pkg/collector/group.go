package collector

import (
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api/slice"
	"github.com/CrazyThursdayV50/gotils/pkg/handlers"
)

func Group[E any, K comparable, V any](sli []E, mapper handlers.CollectorKV[E, K, V]) (m map[K][]V) {
	slice.From(sli...).IterFully(func(_ int, v E) error {
		key, val := mapper(v)
		if m == nil {
			m = make(map[K][]V)
		}
		m[key] = append(m[key], val)
		return nil
	})
	return
}

func GroupValid[E any, K comparable, V any](sli []E, mapper handlers.CollectorKVOkay[E, K, V]) (m map[K][]V) {
	_ = slice.From(sli...).IterOkay(func(_ int, v E) bool {
		key, val, ok := mapper(v)
		if !ok {
			return true
		}

		if m == nil {
			m = make(map[K][]V)
		}

		m[key] = append(m[key], val)
		return true
	})
	return
}

func GroupOkay[E any, K comparable, V any](sli []E, mapper handlers.CollectorKVOkay[E, K, V]) (m map[K][]V, ok bool) {
	_ = slice.From(sli...).IterOkay(func(_ int, v E) bool {
		var key K
		var val V
		key, val, ok = mapper(v)
		if !ok {
			return false
		}

		if m == nil {
			m = make(map[K][]V)
		}

		m[key] = append(m[key], val)
		return true
	})
	return
}

func GroupError[E any, K comparable, V any](sli []E, mapper handlers.CollectorKVError[E, K, V]) (m map[K][]V, err error) {
	_, err = slice.From(sli...).IterError(func(_ int, v E) error {
		key, val, err := mapper(v)
		if err != nil {
			return err
		}
		if m == nil {
			m = make(map[K][]V)
		}
		m[key] = append(m[key], val)
		return nil
	})
	return
}
