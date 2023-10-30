package stack

import (
	"strconv"
	"strings"
)

type stack struct {
	data []string
}

func decodeString(s string) string {
	var container stack
	for _, c := range s {
		if c == ']' {
			var pattern string
			for container.peek() != "[" {
				pattern = container.pop() + pattern
			}
			container.pop()
			var times string
			for !container.isEmpty() && container.peek() >= "0" && container.peek() <= "9" {
				times = container.pop() + times
			}
			count, _ := strconv.Atoi(times)
			container.push(strings.Repeat(pattern, count))
		} else {
			container.push(string(c))
		}
	}
	return container.toString()
}

func (s *stack) peek() string {
	return s.data[len(s.data)-1]
}

func (s *stack) push(str string) {
	s.data = append(s.data, str)
}

func (s *stack) pop() string {
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top

}

func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *stack) toString() string {
	return strings.Join(s.data, "")
}
