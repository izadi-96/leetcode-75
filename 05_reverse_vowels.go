package problems

func reverseVowels(s string) string {
	vowels := map[rune]bool{
		'a': true,
		'e': true,
		'o': true,
		'i': true,
		'u': true,
		'A': true,
		'E': true,
		'O': true,
		'I': true,
		'U': true,
	}

	strArray := []rune(s)
	end := len(s) - 1

	for start, v := range strArray {
		if vowels[v] == true {
			for {
				if vowels[strArray[end]] == true {
					strArray[start], strArray[end] = strArray[end], strArray[start]
					end -= 1
					break
				}
				end -= 1
			}
		}
		if start >= end {
			break
		}
	}
	return string(strArray)
}
