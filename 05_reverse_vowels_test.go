package problems

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseVowel(t *testing.T) {
	testCases := []struct {
		str    string
		result string
	}{
		{str: "leetcode", result: "leotcede"},
		{str: "hello", result: "holle"},
		{str: "ielllei", result: "ielllei"},
		{str: "iellualei", result: "iellaulei"},
		{str: "iieel", result: "eeiil"},
		{str: "iellllei", result: "iellllei"},
	}

	for _, v := range testCases {
		result := reverseVowels(v.str)
		assert.Equal(t, v.result, result)
	}
}
