package problems

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	testCases := []struct {
		pattern string
		input   string
		result  bool
	}{
		//{pattern: "abc", input: "abc", result: true},
		//{pattern: "abc", input: "ahbgdc", result: true},
		//{pattern: "axc", input: "ahbgdc", result: false},
		{pattern: "b", input: "abc", result: true},
	}

	for _, v := range testCases {
		result := isSubsequence(v.pattern, v.input)
		assert.Equal(t, v.result, result)
	}
}
