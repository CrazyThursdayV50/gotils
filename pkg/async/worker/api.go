package worker

import (
	"context"

	gchan "github.com/CrazyThursdayV50/gotils/pkg/builtin/api/chan"
)

func New[J any](do func(job J)) (*Worker[J], func(J)) {
	var m Worker[J]
	m.do = func(j J) {
		do(j)
		m.count++
	}
	trigger := gchan.Make[J](0)
	m.WithContext(context.TODO())
	m.WithTrigger(trigger)
	return &m, func(j J) {
		trigger.Send(j)
	}
}
