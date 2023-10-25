package sliding_window

func longestOnes(nums []int, k int) int {
	var left, right, zeros int
	for right < len(nums) {
		if nums[right] == 0 {
			zeros++
		}
		right++
		if zeros > k {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}
	}
	return right - left
}
