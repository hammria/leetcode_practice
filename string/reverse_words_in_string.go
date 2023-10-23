package main

import "strings"

func reverseWords(s string) string {
	var words = strings.Fields(s)
	var sb strings.Builder
	for i := range words {
		sb.WriteString(words[len(words)-1-i])
		if i != len(words)-1 {
			sb.WriteString(" ")
		}
	}
	return sb.String()
}
