package problems

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	testCases := []struct {
		input  []int
		result []int
	}{
		{input: []int{1, 0, 4, 4}, result: []int{1, 4, 4, 0}},
		{input: []int{0}, result: []int{0}},
		{input: []int{1}, result: []int{1}},
		{input: []int{0, 1}, result: []int{1, 0}},
	}

	for _, v := range testCases {
		moveZeroes(v.input)
		assert.Equal(t, v.result, v.input)
	}
}
