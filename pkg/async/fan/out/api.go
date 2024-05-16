package out

import "github.com/CrazyThursdayV50/gotils/pkg/builtin/api"

type Fan[T any] interface {
	To() api.ChanAPI[T]
	Close()
	Send(element T)
}
