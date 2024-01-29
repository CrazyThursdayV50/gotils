package worker

import "context"

func New[J any](do func(J)) *Worker[J] {
	var m Worker[J]
	m.do = func(j J) {
		do(j)
		m.count++
	}
	m.trigger = make(chan J)
	m.WithContext(context.TODO())
	return &m
}
