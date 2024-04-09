package heap_pq

import (
	"math/rand"
)

func findKthLargest(nums []int, k int) int {
	return find(nums, k)
}

func find(nums []int, k int) int {
	pivotIndex := rand.Intn(len(nums))
	pivot := nums[pivotIndex]
	var greater []int
	var e int
	var less []int
	for _, val := range nums {
		if val > pivot {
			greater = append(greater, val)
		} else if val < pivot {
			less = append(less, val)
		} else if val == pivot {
			e++
		}
	}
	g := len(greater)
	if g >= k {
		return find(greater, k)
	}
	if g+e >= k {
		return pivot
	}
	return find(less, k-g-e)
}
