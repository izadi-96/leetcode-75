package problems

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	index := 0
	for i := 0; i < len(t); i++ {
		if s[index] == t[i] {
			index++
			if index >= len(s) {
				return true
			}
		}
	}
	return false
}
