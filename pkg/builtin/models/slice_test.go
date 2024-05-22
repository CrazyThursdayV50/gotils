package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	slice := FromSlice([]int{1, 2, 3}...)
	assert.Equal(t, 3, slice.Cap())
	assert.Equal(t, 3, slice.Len())

	slice.Append(4)
	assert.True(t, slice.Equal(FromSlice([]int{1, 2, 3, 4}...), func(a int, b int) bool { return a == b }))
	assert.Equal(t, []int{2, 3}, slice.Cut(1, 3))
	assert.Equal(t, slice.Get(0).Unwrap(), 1)
	assert.Equal(t, 0, slice.Index(1, func(a, b int) bool { return a == b }).Get(0).Unwrap())

	slice.Del(0)
	assert.True(t, slice.Equal(FromSlice([]int{2, 3, 4}...), func(a int, b int) bool { return a == b }))

	slice.Clear()
	assert.Equal(t, []int{0, 0, 0}, slice.Unwrap())
}
