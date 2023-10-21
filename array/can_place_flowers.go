package array

func canPlaceFlowers(flowerbed []int, n int) bool {
	var i int
	var left, right bool
	l := len(flowerbed)
	for i < l {
		for i < l && flowerbed[i] != 0 {
			left = true
			i++
		}
		var block int
		for i < l && flowerbed[i] == 0 {
			block++
			i++
		}
		right = i < l
		if left && right {
			n -= (block - 1) / 2
		} else if !left && !right {
			n -= (block + 1) / 2
		} else {
			n -= block / 2
		}
		if n <= 0 {
			return true
		}
	}
	return false
}
