package array

import "sort"

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	arr1 := calOccur(word1)
	arr2 := calOccur(word2)
	for i := range arr1 {
		if (arr1[i] == 0 && arr2[i] != 0) || (arr1[i] != 0 && arr2[i] == 0) {
			return false
		}
	}
	return arrEquals(arr1, arr2)
}

func arrEquals(arr1, arr2 []int) bool {
	sort.Ints(arr1)
	sort.Ints(arr2)
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func calOccur(word string) []int {
	var ans = make([]int, 26)
	for i := range word {
		ans[word[i]-'a']++
	}
	return ans
}
