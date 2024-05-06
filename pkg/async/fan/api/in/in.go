package in

import (
	fanapi "github.com/CrazyThursdayV50/gotils/pkg/async/fan/api"
	"github.com/CrazyThursdayV50/gotils/pkg/async/fan/models/in"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/models"
)

func From[element any, chanRead models.ChanRead[element]](handler func(element), from ...chanRead) fanapi.FanIn {
	return in.From(handler, from...)
}

func New[element any](handler func(element), from ...api.ChanAPIR[element]) fanapi.FanIn {
	return in.New(handler, from...)
}
