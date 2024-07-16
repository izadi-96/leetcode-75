package problems

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeString(t *testing.T) {
	type Pair[T any] struct {
		w1, w2 T
	}

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

func TestGreatestDivisor(t *testing.T) {
	testCases := []struct {
		w1     string
		w2     string
		result string
	}{
		{w1: "CXTXNCXTXNCXTXNCXTXNCXTXN", w2: "CXTXNCXTXNCXTXNCXTXNCXTXNCXTXNCXTXNCXTXNCXTXNCXTXNCXTXNCXTXNCXTXN", result: "CXTXN"},
		{w1: "TAUXXTAUXXTAUXXTAUXXTAUXX", w2: "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX", result: "TAUXX"},
		{w1: "ABCABC", w2: "ABC", result: "ABC"},
		{w1: "TEST", w2: "MAKE", result: ""},
		{w1: "ABCABC", w2: "ABCABC", result: "ABCABC"},
		{w1: "LEET", w2: "CODE", result: ""},
		{w1: "ABABAB", w2: "AB", result: "AB"},
	}

	for _, v := range testCases {
		result := gcdOfStrings(v.w1, v.w2)
		assert.Equal(t, v.result, result)
	}
}

// 05_reverse_vowels.go
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

// 06_reverse_words.go
func TestReverseWords(t *testing.T) {
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

// 07_product_array_except_self.go
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

// 10_move_zeroes.go
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

// 11_is_subsequence.go
func TestIsSubsequence(t *testing.T) {
	testCases := []struct {
		pattern string
		input   string
		result  bool
	}{
		{pattern: "abc", input: "abc", result: true},
		{pattern: "abc", input: "ahbgdc", result: true},
		{pattern: "axc", input: "ahbgdc", result: false},
		{pattern: "b", input: "abc", result: true},
	}

	for _, v := range testCases {
		result := isSubsequence(v.pattern, v.input)
		assert.Equal(t, v.result, result)
	}
}

// 12_container_with_most_water.go
func TestContainerWithMostWater(t *testing.T) {
	testCases := []struct {
		input  []int
		result int
	}{
		{input: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, result: 49},
		{input: []int{1, 8, 6, 2, 5, 4, 8, 25, 7}, result: 49},
		{input: []int{2, 2, 2, 2, 0, 0, 0, 0, 0, 1}, result: 9},
		{input: []int{1, 1}, result: 1},
		{input: []int{1, 2, 4, 3}, result: 4},
		{input: []int{2, 3, 10, 5, 7, 8, 9}, result: 36},
		{input: []int{1, 3, 2, 5, 25, 24, 5}, result: 24},
	}

	for _, v := range testCases {
		result := maxArea(v.input)
		assert.Equal(t, v.result, result)
	}
}
