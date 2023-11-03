package queue

func predictPartyVictory(senate string) string {
	var r, d, dropR, dropD int
	for _, c := range senate {
		if c == 'R' {
			r++
		} else {
			d++
		}
	}
	arr := []byte(senate)
	for r > 0 && d > 0 {
		for i, c := range arr {
			if c == 'R' {
				if dropR > 0 {
					dropR--
					r--
					arr[i] = 'x'
				} else {
					dropD++
				}
			} else if c == 'D' {
				if dropD > 0 {
					dropD--
					d--
					arr[i] = 'x'
				} else {
					dropR++
				}
			}
		}
	}
	if r == 0 {
		return "Dire"
	}
	return "Radiant"
}
