package greedy

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	var area int
	var preMin int
	for left < right {
		if height[left] > height[right] {
			if height[right] > preMin {
				area = maxInt(area, height[right]*(right-left))
				preMin = height[right]
			}
			right--
		} else {
			if height[left] > preMin {
				area = maxInt(area, height[left]*(right-left))
				preMin = height[left]
			}
			left++
		}
	}
	return area
}

func maxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}
