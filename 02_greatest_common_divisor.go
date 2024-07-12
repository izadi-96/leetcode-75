package problems

func gcdOfStrings(str1, str2 string) string {
	l1, l2 := len(str1), len(str2)
	gcdLen := gcd(l1, l2)

	if str1+str2 != str2+str1 {
		return ""
	}
	return str1[:gcdLen]
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
