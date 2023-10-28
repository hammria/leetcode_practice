package stack

type stack struct {
	data []int
	len  int
}

func asteroidCollision(asteroids []int) []int {
	var container = stack{
		data: make([]int, len(asteroids)),
		len:  0,
	}

	for _, v := range asteroids {
		if container.isEmpty() || container.peek() < 0 || v > 0 {
			container.push(v)
			continue
		}
		var destroyed bool
		for !container.isEmpty() && container.peek() > 0 {
			if container.peek()+v > 0 {
				destroyed = true
				break
			} else if container.peek()+v == 0 {
				container.pop()
				destroyed = true
				break
			} else {
				container.pop()
			}
		}
		if !destroyed {
			container.push(v)
		}
	}
	return container.toArray()
}

func (s *stack) peek() int {
	return s.data[s.len-1]
}

func (s *stack) push(i int) {
	s.data[s.len] = i
	s.len++
}

func (s *stack) pop() int {
	top := s.data[s.len-1]
	s.data[s.len-1] = 0
	s.len--
	return top

}

func (s *stack) isEmpty() bool {
	return s.len == 0
}

func (s *stack) toArray() []int {
	return s.data[0:s.len]
}
