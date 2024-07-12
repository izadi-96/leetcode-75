package problems

import (
	"bytes"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)
	var buff bytes.Buffer
	for i := len(words) - 1; i >= 0; i-- {
		buff.WriteString(words[i])
		if i != 0 {
			buff.WriteString(" ")
		}
	}
	return buff.String()
}
