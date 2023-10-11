package array

import "math"

func increasingTriplet(nums []int) bool {
	mid := math.MaxInt
	left := math.MaxInt - 1
	for i := range nums {
		if nums[i] <= left {
			left = nums[i]
		} else if nums[i] <= mid {
			mid = nums[i]
		} else {
			return true
		}
	}
	return false
}
