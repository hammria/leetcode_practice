package array

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var res = make([]bool, len(candies))
	var m int
	for _, v := range candies {
		if v > m {
			m = v
		}
	}
	for i := range res {
		if candies[i]+extraCandies >= m {
			res[i] = true
		}
	}
	return res
}
