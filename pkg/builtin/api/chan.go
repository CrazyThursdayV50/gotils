package api

import (
	"time"

	"github.com/CrazyThursdayV50/gotils/pkg/wrapper"
)

type (
	baseCommonChanAPI[E any] interface {
		Len() int
		IsEmpty() bool
		Closed() bool
	}

	baseReadWriteChanAPI[E any] interface {
		Inner() chan E
		Renew(buffer int)
		RenewForce(buffer int)
	}

	baseWriteChanAPI[E any] interface {
		InnerW() chan<- E
		Close()
		Send(element E)
		SendTimeout(send time.Duration)
	}

	baseReadChanAPI[E any] interface {
		InnerR() <-chan E
		Receive() (wrapper.UnWrapper[E], bool)
		RecvTimeout(recv time.Duration)
		Iter[int, E, any]
	}

	ChanAPIR[E any] interface {
		baseCommonChanAPI[E]
		baseReadChanAPI[E]
	}

	ChanAPIW[E any] interface {
		baseCommonChanAPI[E]
		baseWriteChanAPI[E]
	}

	ChanAPI[E any] interface {
		baseCommonChanAPI[E]
		baseWriteChanAPI[E]
		baseReadChanAPI[E]
		baseReadWriteChanAPI[E]
	}
)

// var (
// 	_ ChanAPI[any]  = (*models.ChanRW[any])(nil)
// 	_ ChanAPIR[any] = (*models.ChanR[any])(nil)
// 	_ ChanAPIR[any] = (*models.ChanRW[any])(nil)
// 	_ ChanAPIW[any] = (*models.ChanW[any])(nil)
// 	_ ChanAPIW[any] = (*models.ChanRW[any])(nil)
// )
