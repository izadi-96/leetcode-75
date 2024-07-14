package problems

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	testCases := []struct {
		input  []int
		result []int
	}{
		{input: []int{1, 2, 3, 4}, result: []int{24, 12, 8, 6}},
		{input: []int{2, 2, 2, 2}, result: []int{8, 8, 8, 8}},
		{input: []int{1, -1, 1, -1, 1}, result: []int{1, -1, 1, -1, 1}},
	}

	for _, v := range testCases {
		result := productExceptSelfOptimize(v.input)
		assert.Equal(t, v.result, result)
	}
}
