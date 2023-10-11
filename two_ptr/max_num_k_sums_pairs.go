package two_ptr

import "sort"

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	left := 0
	right := len(nums) - 1
	ans := 0
	for left < right {
		if nums[left]+nums[right] == k {
			left++
			right--
			ans++
		} else if nums[left]+nums[right] > k {
			right--
		} else {
			left++
		}
	}
	return ans
}
