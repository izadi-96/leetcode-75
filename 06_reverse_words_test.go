package problems

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		str    string
		result string
	}{
		{str: "a ba c", result: "c ba a"},
		{str: " make me free", result: "free me make"},
		{str: " make me free ", result: "free me make"},
		{str: " make  me      free ", result: "free me make"},
		{str: "   ", result: ""},
	}

	for _, v := range testCases {
		result := reverseWords(v.str)
		assert.Equal(t, v.result, result)
	}
}
