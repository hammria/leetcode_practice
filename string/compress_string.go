package main

import "strconv"

func compress(chars []byte) int {
	var slow, fast int
	for fast < len(chars) {
		cur := chars[fast]
		chars[slow] = cur
		count := 0
		for fast < len(chars) && cur == chars[fast] {
			fast++
			count++
		}
		slow++
		if count > 1 {
			b := strconv.Itoa(count)
			for i := 0; i < len(b); i++ {
				chars[slow] = b[i]
				slow++
			}
		}
	}
	return slow
}
