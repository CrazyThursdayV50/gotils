package leader

import (
	"context"
	"testing"
	"time"

	"github.com/CrazyThursdayV50/gotils/pkg/async/worker"
)

func TestLeader(t *testing.T) {
	leader := New[int](context.TODO(), time.Millisecond, time.Millisecond)
	handler1 := func(id int) {
		t.Logf("id[1]: %d\n", id)
	}
	handler2 := func(id int) {
		t.Logf("id[2]: %d\n", id)
	}

	leader.AddWorker(worker.New[int](handler1))
	leader.AddWorker(worker.New[int](handler2))
	for id := range make([]int, 100) {
		leader.Do(id)
	}

	<-make(chan struct{})
}
