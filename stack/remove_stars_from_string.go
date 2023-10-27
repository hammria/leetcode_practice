package stack

func removeStars(s string) string {
	var stack = make([]byte, len(s))
	tail := 0
	for i := range s {
		if s[i] == '*' && len(stack) >= 0 {
			tail--
		} else {
			stack[tail] = s[i]
			tail++
		}
	}
	return string(stack[0:tail])
}
