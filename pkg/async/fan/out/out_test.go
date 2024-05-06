package out

import (
	"fmt"
	"testing"
)

func TestOut(t *testing.T) {
	fan := New[any](10, 1, func(t any) {
		fmt.Printf("receive %v\n", t)
	})

	for i := range int(1e2) {
		fan.Send(i)
	}
}
