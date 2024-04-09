package graph

func findCircleNum(isConnected [][]int) int {
	provinces := make([]int, len(isConnected))
	for i := range provinces {
		provinces[i] = i
	}

	for city := range isConnected {
		for i := city + 1; i < len(isConnected); i++ {
			if isConnected[city][i] == 1 {
				provinces[find(provinces, city)] = provinces[find(provinces, i)]
			}
		}

	}
	var ans int
	for i := range provinces {
		if provinces[i] == i {
			ans++
		}
	}
	return ans
}

func find(provinces []int, x int) int {
	if provinces[x] == x {
		return x
	}
	provinces[x] = find(provinces, provinces[x])
	return provinces[x]
}
