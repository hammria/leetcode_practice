package main

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	sl := len(s)
	tl := len(t)
	if sl > tl {
		return false
	}
	var head int
	for i := 0; i < tl; i++ {
		if s[head] == t[i] {
			head++
			if head == sl {
				return true
			}
		}
		if sl-head > tl-i {
			return false
		}
	}
	return false
}
