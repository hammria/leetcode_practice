package main

func gcdOfStrings(str1 string, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)
	maxPossible := gcd(len1, len2)
	if maxPossible == len1 && len1 == len2 && str1 != str2 {
		maxPossible--
	}
	for maxPossible >= 1 {
		if str1[0:maxPossible] == str2[0:maxPossible] {
			if isMultipleOf(str1, maxPossible) && isMultipleOf(str2, maxPossible) {
				return str1[0:maxPossible]
			}
		}

		maxPossible--
	}
	return ""

}

func isMultipleOf(str string, div int) bool {
	if len(str)%div != 0 {
		return false
	}
	for i := 1; i < len(str)/div; i++ {
		for j := 0; j < div; j++ {
			if str[i*div+j] != str[j] {
				return false
			}
		}
	}
	return true
}

func gcd(a, b int) int {
	if a == b {
		return a
	} else if a > b {
		return gcd(a-b, b)
	} else {
		return gcd(a, b-a)
	}
}
