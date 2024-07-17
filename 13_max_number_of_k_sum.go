package problems

func maxOperations(nums []int, k int) int {
	if len(nums) < 2 {
		return 0
	}
	pairCount := 0
	hashmap := make(map[int]int)
	for _, num := range nums {
		complement := k - num
		if count, found := hashmap[complement]; found && count > 0 {
			pairCount++
			hashmap[complement]--
		} else {
			hashmap[num]++
		}
	}
	return pairCount
}
