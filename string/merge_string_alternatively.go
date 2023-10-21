package main

func mergeAlternately(word1 string, word2 string) string {
	l1 := len(word1)
	l2 := len(word2)
	var res = make([]byte, l1+l2)
	var i int
	for i < l1 && i < l2 {
		res[2*i] = word1[i]
		res[2*i+1] = word2[i]
		i++
	}
	if l1 < l2 {
		for i < l2 {
			res[i+l1] = word2[i]
			i++
		}
	} else {
		for i < l1 {
			res[i+l2] = word1[i]
			i++
		}
	}
	return string(res)
}
