package two_ptr

type void struct {
}

var placeholder void
var vowels = map[string]void{
	"a": placeholder,
	"e": placeholder,
	"i": placeholder,
	"o": placeholder,
	"u": placeholder,
	"A": placeholder,
	"E": placeholder,
	"I": placeholder,
	"O": placeholder,
	"U": placeholder,
}

func reverseVowels(s string) string {
	chars := []byte(s)
	left := 0
	right := len(s) - 1
	for left < right {
		for left < right {
			if _, ok := vowels[string(chars[left])]; !ok {
				left++
			} else {
				break
			}
		}
		for left < right {
			if _, ok := vowels[string(chars[right])]; !ok {
				right--
			} else {
				break
			}
		}
		temp := chars[left]
		chars[left] = chars[right]
		chars[right] = temp
		left++
		right--
	}
	return string(chars)
}
