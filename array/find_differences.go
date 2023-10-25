package array

type void struct {
}

var placeHolder void

func findDifference(nums1 []int, nums2 []int) [][]int {
	s1 := make(map[int]void)
	s2 := make(map[int]void)
	for _, v := range nums1 {
		s1[v] = placeHolder
	}
	for _, v := range nums2 {
		s2[v] = placeHolder
	}
	for k := range s1 {
		_, ok := s2[k]
		if ok {
			delete(s1, k)
			delete(s2, k)
		}
	}
	var nums = make([][]int, 2)
	var onlyIn1 = make([]int, len(s1))
	var onlyIn2 = make([]int, len(s2))
	var i int
	for k := range s1 {
		onlyIn1[i] = k
		i++
	}
	i = 0
	for k := range s2 {
		onlyIn2[i] = k
		i++
	}
	nums[0] = onlyIn1
	nums[1] = onlyIn2
	return nums
}
