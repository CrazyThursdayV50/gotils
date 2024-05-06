package out

import (
	fanapi "github.com/CrazyThursdayV50/gotils/pkg/async/fan/api"
	"github.com/CrazyThursdayV50/gotils/pkg/async/fan/models/out"
)

func New[element any](count, buffer int, handler func(element)) fanapi.FanOut[element] {
	return out.New[element](count, buffer, handler)
}
