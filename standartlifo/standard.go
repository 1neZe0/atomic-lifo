package standardlifo

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return 0
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) Peek() int {
	if len(s.items) == 0 {
		return 0
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) Len() int {
	return len(s.items)
}
