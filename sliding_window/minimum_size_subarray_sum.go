package sliding_window

import "math"

func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	ans := math.MaxInt
	curSum := 0
	for right < len(nums) {
		curSum += nums[right]
		for curSum >= target {
			ans = minInt(ans, right-left+1)
			curSum -= nums[left]
			left++
		}
		right++
	}
	if ans == math.MaxInt {
		return 0
	}
	return ans
}

func minInt(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
