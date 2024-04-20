package worker

import (
	"testing"
	"time"
)

func TestWork(t *testing.T) {
	t.Logf("start\n")
	worker := New(func(id int) {
		t.Logf("id: %d\n", id)
	})

	worker.Run()
	time.Sleep(time.Millisecond * 10)
	for i := range make([]int, 100) {
		worker.Delivery(i)
	}

	<-make(chan struct{})
	worker.Stop()
}
