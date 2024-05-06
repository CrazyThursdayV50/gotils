package api

import "github.com/CrazyThursdayV50/gotils/pkg/builtin/api"

type (
	FanIn interface {
		Close()
	}

	FanOut[T any] interface {
		To() api.ChanAPI[T]
		Close()
		Send(element T)
	}
)
