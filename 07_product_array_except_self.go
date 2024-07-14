package problems

// v1
func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	prefixProduct := make([]int, n)
	prefixProduct[0] = 1

	postfixProduct := make([]int, n)
	postfixProduct[n-1] = 1

	for i := 1; i < n; i++ {
		prefixProduct[i] = prefixProduct[i-1] * nums[i-1]
	}

	for i := n - 2; i >= 0; i-- {
		postfixProduct[i] = postfixProduct[i+1] * nums[i+1]
	}
	//
	for i := 0; i < n; i++ {
		result[i] = prefixProduct[i] * postfixProduct[i]
	}

	return result
}

func productExceptSelfOptimize(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	result[0] = 1

	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	postfixProduct := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= postfixProduct
		postfixProduct *= nums[i]
	}

	return result
}
