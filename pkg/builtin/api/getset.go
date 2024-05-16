package api

import "github.com/CrazyThursdayV50/gotils/pkg/wrapper"

type GetSeter[K any, V any] interface {
	Get(K) wrapper.UnWrapper[V]
	Set(k K, v V)
}
