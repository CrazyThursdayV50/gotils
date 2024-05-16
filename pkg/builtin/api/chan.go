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
		Chan() chan E
		Renew(buffer int)
		RenewForce(buffer int)
	}

	baseWriteChanAPI[E any] interface {
		ChanW() chan<- E
		Close()
		Send(element E)
		SendTimeout(send time.Duration)
	}

	baseReadChanAPI[E any] interface {
		ChanR() <-chan E
		Receive() (wrapper.UnWrapper[E], bool)
		RecvTimeout(recv time.Duration)
		Iter[int, E]
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
