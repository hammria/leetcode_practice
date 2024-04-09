package heap_pq

import "container/heap"

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	n := len(nums1)
	sort(nums1, nums2, 0, n-1)
	h := msHeap{}
	sum := 0
	for i := 0; i < k; i++ {
		heap.Push(&h, nums1[i])
		sum += nums1[i]
	}
	ans := sum * nums2[k-1]
	for i := k; i < n; i++ {
		heap.Push(&h, nums1[i])
		sum += nums1[i]
		sum -= heap.Pop(&h).(int)
		if ans < sum*nums2[i] {
			ans = sum * nums2[i]
		}
	}
	return int64(ans)
}

func sort(nums1, nums2 []int, start, end int) {
	if start >= end {
		return
	}
	pivot := nums2[start]
	i := start
	j := end
	for i < j {
		for i < j && nums2[j] <= pivot {
			j--
		}
		for i < j && nums2[i] >= pivot {
			i++
		}
		if i < j {
			nums2[i], nums2[j] = nums2[j], nums2[i]
			nums1[i], nums1[j] = nums1[j], nums1[i]
		} else {
			break
		}
	}
	nums2[start], nums2[i] = nums2[i], nums2[start]
	nums1[start], nums1[i] = nums1[i], nums1[start]
	sort(nums1, nums2, start, i-1)
	sort(nums1, nums2, i+1, end)
}

type msHeap []int

func (h *msHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *msHeap) Pop() any {
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return val
}
func (h msHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h msHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h msHeap) Len() int {
	return len(h)
}
