package array

func findMaxAverage(nums []int, k int) float64 {
	var i int
	var cur int
	for i < k {
		cur += nums[i]
		i++
	}
	maxSum := cur
	for i < len(nums) {
		cur = cur + nums[i] - nums[i-k]
		if cur > maxSum {
			maxSum = cur
		}
		i++
	}
	return float64(maxSum) / float64(k)
}
