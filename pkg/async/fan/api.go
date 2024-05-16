package fan

import (
	"github.com/CrazyThursdayV50/gotils/pkg/async/fan/in"
	fanin "github.com/CrazyThursdayV50/gotils/pkg/async/fan/in/implement"
	"github.com/CrazyThursdayV50/gotils/pkg/async/fan/out"
	fanout "github.com/CrazyThursdayV50/gotils/pkg/async/fan/out/implement"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/models"
)

func NewIn[element any](handler func(element), from ...api.ChanAPIR[element]) in.Fan {
	return fanin.New(handler, from...)
}

func NewInFrom[element any, chanRead models.ChanRead[element]](handler func(element), from ...chanRead) in.Fan {
	return fanin.From(handler, from...)
}

func NewOut[element any](count, buffer int, handler func(element)) out.Fan[element] {
	return fanout.New[element](count, buffer, handler)
}
