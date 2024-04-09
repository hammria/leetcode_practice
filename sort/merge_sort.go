package sort

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	left := nums[:len(nums)/2]
	right := nums[len(nums)/2:]
	left = mergeSort(left)
	right = mergeSort(right)
	return merge(left, right)
}

func merge(nums1, nums2 []int) []int {
	if len(nums1) == 0 {
		return nums2
	}
	if len(nums2) == 0 {
		return nums1
	}
	res := make([]int, len(nums1)+len(nums2))
	i := 0
	j := 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			res[i+j] = nums1[i]
			i++
		} else {
			res[i+j] = nums2[j]
			j++
		}
	}
	for i < len(nums1) {
		res[i+j] = nums1[i]
		i++
	}
	for j < len(nums2) {
		res[i+j] = nums2[j]
		j++
	}
	return res
}
