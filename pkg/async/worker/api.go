package worker

import (
	"context"

	gchan "github.com/CrazyThursdayV50/gotils/pkg/builtin/api/chan"
)

func New[J any](do func(J)) *Worker[J] {
	var m Worker[J]
	m.do = func(j J) {
		do(j)
		m.count++
	}
	m.trigger = gchan.Make[J](0)
	m.WithContext(context.TODO())
	return &m
}
