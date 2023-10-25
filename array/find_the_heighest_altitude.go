package array

func largestAltitude(gain []int) int {
	var cur, ans int
	for _, v := range gain {
		cur += v
		if cur > ans {
			ans = cur
		}
	}
	return ans
}
