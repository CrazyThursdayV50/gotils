package slice

import (
	"cmp"

	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/models"
	"github.com/CrazyThursdayV50/gotils/pkg/wrapper"
	"github.com/CrazyThursdayV50/gotils/pkg/wrapper/wrap"
)

func From[E any](sli ...E) api.SliceAPI[E] {
	return models.FromSlice(sli...)
}

func Make[E any](len, cap int) api.SliceAPI[E] {
	return models.MakeSlice[E](len, cap)
}

func Empty(len int) api.SliceAPI[struct{}] {
	return models.MakeSlice[struct{}](len, len)
}

func Collect[E any, T any](sli []E, collector func(element E) T) api.SliceAPI[T] {
	src := models.FromSlice(sli...)
	dst := models.MakeSlice[T](0, src.Len())
	_ = src.IterOkay(func(_ int, e E) bool {
		dst.Append(collector(e))
		return true
	})
	return dst
}

func CollectValid[E any, T any](sli []E, collector func(E) (wrapper.UnWrapper[T], bool)) api.SliceAPI[T] {
	src := models.FromSlice(sli...)
	dst := models.MakeSlice[T](0, src.Len())
	_ = src.IterOkay(func(_ int, e E) bool {
		w, ok := collector(e)
		if !ok {
			return true
		}

		dst.Append(w.Unwrap())
		return true
	})
	return dst
}

func CollectCtrl[E any, T any](sli []E, collector func(E) (T, error)) (api.SliceAPI[T], error) {
	src := models.FromSlice(sli...)
	dst := models.MakeSlice[T](0, src.Len())
	_, err := src.IterError(func(_ int, e E) error {
		t, er := collector(e)
		if er != nil {
			return er
		}
		dst.Append(t)
		return nil
	})

	return dst, err
}

func Map[E any, K cmp.Ordered, V any](sli []E, mapper func(E) (K, V)) api.MapAPI[K, V, any] {
	src := models.FromSlice(sli...)
	dst := models.MakeMap[K, V, any](src.Len())
	src.IterFully(func(_ int, e E) error {
		k, v := mapper(e)
		dst.Set(k, v)
		return nil
	})
	return dst
}

func MapP[E any, K any, V any](sli []E, mapper func(E) (*K, V)) api.MapAPI[*K, V, K] {
	src := models.FromSlice(sli...)
	dst := models.MakeMap[*K, V, K](src.Len())
	src.IterFully(func(_ int, e E) error {
		k, v := mapper(e)
		dst.Set(k, v)
		return nil
	})
	return dst
}

func Group[E any, K cmp.Ordered, V any](sli []E, mapper func(E) (K, V)) api.MapAPI[K, api.SliceAPI[V], any] {
	src := models.FromSlice(sli...)
	dst := models.MakeMap[K, api.SliceAPI[V], any](src.Len())
	src.IterFully(func(_ int, e E) error {
		k, v := mapper(e)
		group := dst.Get(k)
		if group == nil {
			group = wrap.Wrap(api.SliceAPI[V](models.FromSlice[V]()))
			dst.Set(k, group.Unwrap())
		}
		group.Unwrap().Append(v)
		return nil
	})
	return dst
}

func GroupP[E any, K any, V any](sli []E, mapper func(E) (*K, V)) api.MapAPI[*K, api.SliceAPI[V], K] {
	src := models.FromSlice(sli...)
	dst := models.MakeMap[*K, api.SliceAPI[V], K](src.Len())
	src.IterFully(func(_ int, e E) error {
		k, v := mapper(e)
		group := dst.Get(k)
		if group == nil {
			group = wrap.Wrap(api.SliceAPI[V](models.FromSlice[V]()))
			dst.Set(k, group.Unwrap())
		}
		group.Unwrap().Append(v)
		return nil
	})
	return dst
}

func MapCtrl[E any, K cmp.Ordered, V any](sli []E, mapper func(E) (K, V, error)) (api.MapAPI[K, V, any], error) {
	src := models.FromSlice(sli...)
	dst := models.MakeMap[K, V, any](src.Len())
	_, err := src.IterError(func(_ int, e E) error {
		k, v, err := mapper(e)
		if err != nil {
			return err
		}
		dst.Set(k, v)
		return nil
	})
	return dst, err
}

func MapPCtrl[E any, K any, V any](sli []E, mapper func(E) (*K, V, error)) (api.MapAPI[*K, V, K], error) {
	src := models.FromSlice(sli...)
	dst := models.MakeMap[*K, V, K](src.Len())
	_, err := src.IterError(func(_ int, e E) error {
		k, v, err := mapper(e)
		if err != nil {
			return err
		}
		dst.Set(k, v)
		return nil
	})
	return dst, err
}

func GroupCtrl[E any, K cmp.Ordered, V any](sli []E, mapper func(E) (K, V, error)) (m api.MapAPI[K, api.SliceAPI[V], any], err error) {
	src := models.FromSlice(sli...)
	_, err = src.IterError(func(_ int, e E) error {
		k, v, err := mapper(e)
		if err != nil {
			return err
		}

		if m == nil {
			m = models.MakeMap[K, api.SliceAPI[V], any](src.Len())
		}

		group := m.Get(k)
		if group == nil {
			group = wrap.Wrap(api.SliceAPI[V](models.FromSlice[V]()))
			m.Set(k, group.Unwrap())
		}

		group.Unwrap().Append(v)
		return nil
	})

	return
}

func GroupPCtrl[E any, K any, V any](sli []E, mapper func(E) (*K, V, error)) (m api.MapAPI[*K, api.SliceAPI[V], K], err error) {
	src := models.FromSlice(sli...)
	_, err = src.IterError(func(_ int, e E) error {
		k, v, err := mapper(e)
		if err != nil {
			return err
		}

		if m == nil {
			m = models.MakeMap[*K, api.SliceAPI[V], K](src.Len())
		}

		group := m.Get(k)
		if group == nil {
			group = wrap.Wrap(api.SliceAPI[V](models.FromSlice[V]()))
			m.Set(k, group.Unwrap())
		}

		group.Unwrap().Append(v)
		return nil
	})

	return
}
