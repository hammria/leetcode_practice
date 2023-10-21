package array

func productExceptSelf(nums []int) []int {
	l := len(nums)
	left := make([]int, l)
	left[0] = 1
	right := make([]int, l)
	right[l-1] = 1
	for i := 1; i < l; i++ {
		left[i] = nums[i-1] * left[i-1]
		right[l-i-1] = nums[l-i] * right[l-i]
	}
	for i := 0; i < l; i++ {
		nums[i] = left[i] * right[i]
	}
	return nums
}
