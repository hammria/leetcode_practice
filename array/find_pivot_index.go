package array

func pivotIndex(nums []int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	var curSum int
	for i, v := range nums {
		if i > 0 {
			curSum += nums[i-1]
		}
		sum -= v
		if curSum == sum {
			return i
		}
	}
	return -1
}
