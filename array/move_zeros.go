package array

func moveZeroes(nums []int) {
	var zeros int
	for i, v := range nums {
		if v == 0 {
			zeros++
		} else if zeros > 0 {
			nums[i-zeros] = v
			nums[i] = 0
		}
	}
}
