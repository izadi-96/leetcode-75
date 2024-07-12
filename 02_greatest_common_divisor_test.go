package problems

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDivisor(t *testing.T) {
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
