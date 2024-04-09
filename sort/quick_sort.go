package sort

import "math/rand"

func quickSort(nums []int) {
	helper(nums, 0, len(nums)-1)
}

func helper(nums []int, start, end int) {
	if start >= end {
		return
	}
	pivotIndex := start + rand.Intn(end-start)
	swap(nums, start, pivotIndex)
	pivot := nums[start]
	left := start
	right := end
	for left < right {
		for left < right && nums[left] < pivot {
			left++
		}
		for left < right && nums[right] >= pivot {
			right--
		}
		if left < right {
			swap(nums, left, right)
		} else {
			break
		}
	}
	nums[left] = pivot
	helper(nums, start, left-1)
	helper(nums, left+1, end)
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
