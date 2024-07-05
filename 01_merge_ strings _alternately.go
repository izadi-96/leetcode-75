package problems

import "bytes"

func MergeAlternately(word1 string, word2 string) string {
	len1 := len(word1)
	len2 := len(word2)

	var buffer bytes.Buffer
	buffer.Grow(len1 + len2)

	for i := 0; i < len1 || i < len2; i++ {
		if i < len1 {
			buffer.WriteString(string(word1[i]))
		}
		if i < len2 {
			buffer.WriteString(string(word2[i]))
		}
	}
	return buffer.String()
}
