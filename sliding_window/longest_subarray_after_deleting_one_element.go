package sliding_window

func longestSubarray(nums []int) int {
	var left, right, zeros int
	for right < len(nums) {
		if nums[right] == 0 {
			zeros++
		}
		right++
		if zeros > 1 {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}
	}
	return right - left - 1
}
