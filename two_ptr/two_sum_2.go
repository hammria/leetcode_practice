package two_ptr

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	res := []int{left, right}
	for left < right {
		if numbers[left]+numbers[right] > target {
			right--
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			res[0] = left + 1
			res[1] = right + 1
			break
		}
	}
	return res
}
