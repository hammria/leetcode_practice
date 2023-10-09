package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	nums2 := []int{}
	nums3 := []int{1, 4}
	fmt.Println(searchRange(nums, 6))
	fmt.Println(searchRange(nums, 8))
	fmt.Println(searchRange(nums2, 8))
	fmt.Println(searchRange(nums3, 4))

}

func searchRange(nums []int, target int) []int {
	var res = []int{-1, -1}
	lower := 0
	upper := len(nums) - 1
	index := -1
	mid := 0
	for lower <= upper {
		mid = (lower + upper) / 2
		if nums[mid] == target {
			upper = mid - 1
			index = mid
		} else if nums[mid] < target {
			lower = mid + 1
		} else {
			upper = mid - 1
		}
	}
	res[0] = index

	lower = 0
	upper = len(nums) - 1
	index = -1
	for lower <= upper {
		mid = (lower + upper) / 2
		if nums[mid] == target {
			lower = mid + 1
			index = mid
		} else if nums[mid] < target {
			lower = mid + 1
		} else {
			upper = mid - 1
		}

	}
	res[1] = index

	return res
}
