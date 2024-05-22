package models

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	slice := FromSlice([]int{1, 2, 3}...)
	assert.Equal(t, 3, slice.Cap())
	assert.Equal(t, 3, slice.Len())
	assert.Equal(t, slice.Get(0).Unwrap(), 1)

	slice.Append(4)
	assert.True(t, slices.EqualFunc([]int{1, 2, 3, 4}, slice.Unwrap(), func(a int, b int) bool { return a == b }))
	assert.Equal(t, 0, slices.IndexFunc(slice.Unwrap(), func(element int) bool { return element == 1 }))

	slice.Clear()
	assert.Equal(t, []int{0, 0, 0, 0}, slice.Unwrap())
}
