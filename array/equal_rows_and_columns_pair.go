package array

func equalPairs(grid [][]int) int {
	var ans int
	for _, v := range grid {
		for j := 0; j < len(grid); j++ {
			k := 0
			for k < len(grid) {
				if v[k] != grid[k][j] {
					break
				}
				k++
			}
			if k == len(grid) {
				ans++
			}
		}
	}
	return ans
}
