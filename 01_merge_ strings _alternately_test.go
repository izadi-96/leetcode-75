package problems

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Pair[T any] struct {
	w1, w2 T
}

func TestMerge(t *testing.T) {
	testCases := []struct {
		pair   Pair[string]
		result string
	}{
		{pair: Pair[string]{w1: "ab", w2: "cd"}, result: "acbd"},
		{pair: Pair[string]{w1: "ab", w2: "pqrs"}, result: "apbqrs"},
		{pair: Pair[string]{w1: "", w2: ""}, result: ""},
		{pair: Pair[string]{w1: "", w2: "abcd"}, result: "abcd"},
		{pair: Pair[string]{w1: "abcd", w2: ""}, result: "abcd"},
	}

	for _, v := range testCases {
		res := MergeAlternately(v.pair.w1, v.pair.w2)
		assert.Equal(t, v.result, res)
	}
}
