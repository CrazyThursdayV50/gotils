package collector

import (
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api/slice"
	"github.com/CrazyThursdayV50/gotils/pkg/handlers"
)

func Slice[E any, T any](sli []E, collector handlers.Collector[E, T]) (list []T) {
	_ = slice.From(sli...).IterFully(func(_ int, v E) error {
		list = append(list, collector(v))
		return nil
	})
	return
}

func SliceError[E any, T any](sli []E, collector handlers.CollectorError[E, T]) (list []T, err error) {
	_, err = slice.From(sli...).IterError(func(_ int, v E) error {
		t, err := collector(v)
		if err != nil {
			return err
		}
		list = append(list, t)
		return nil
	})
	return
}

func SliceValid[E any, T any](sli []E, collector handlers.CollectorOkay[E, T]) (list []T) {
	_ = slice.From(sli...).IterOkay(func(_ int, v E) bool {
		t, ok := collector(v)
		if !ok {
			return true
		}
		list = append(list, t)
		return true
	})
	return
}

func SliceOkay[E any, T any](sli []E, collector handlers.CollectorOkay[E, T]) (list []T, ok bool) {
	_ = slice.From(sli...).IterOkay(func(_ int, v E) bool {
		var t T
		t, ok = collector(v)
		if !ok {
			return false
		}
		list = append(list, t)
		return true
	})
	return
}
