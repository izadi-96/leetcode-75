package problems

func maxArea(height []int) int {
	if len(height) == 1 {
		return 0
	}
	start := 0
	end := len(height) - 1

	area := 0
	for start < end {
		tmpArea := min(height[start], height[end]) * (end - start)
		if tmpArea > area {
			area = tmpArea
		}
		if height[start] > height[end] {
			end -= 1
		} else {
			start += 1
		}
	}
	return area
}
