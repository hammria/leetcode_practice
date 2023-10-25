package array

type void struct {
}

func uniqueOccurrences(arr []int) bool {

	var occurrence = make(map[int]int)
	for _, v := range arr {
		count, ok := occurrence[v]
		if !ok {
			occurrence[v] = 1
		} else {
			occurrence[v] = count + 1
		}
	}
	var placeHolder void
	var set = make(map[int]void)
	for _, v := range occurrence {
		_, ok := set[v]
		if !ok {
			set[v] = placeHolder
		} else {
			return false
		}
	}
	return true
}
