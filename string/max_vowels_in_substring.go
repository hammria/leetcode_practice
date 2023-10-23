package main

func maxVowels(s string, k int) int {
	var i int
	var cur int
	for i < k {
		if isVowel(s[i]) {
			cur++
		}
		i++
	}
	maxCount := cur
	for i < len(s) {
		if isVowel(s[i]) {
			cur++
		}
		if isVowel(s[i-k]) {
			cur--
		}
		if cur > maxCount {
			maxCount = cur
		}
		i++
	}
	return maxCount
}

func isVowel(b byte) bool {
	if b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' {
		return true
	}
	return false
}
