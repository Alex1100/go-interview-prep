package stack

type Stack struct {
	items []int
	size  int
}

func InitStack() *Stack {
	var items = make([]int, 0)
	return &Stack{
		items: items,
		size:  0,
	}
}

func (s *Stack) Insert(item int) error {
	s.items = append(s.items, item)
	s.size++
	return nil
}

func (s *Stack) Pop() (int, error) {
	removed := s.items[len(s.items)-1]
	if len(s.items) > 1 {
		s.items = s.items[0 : len(s.items)-1]
	} else {
		s.items = make([]int, 0)
	}

	s.size--
	return removed, nil
}

func (s *Stack) Front() int {
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	if s.Size() == 0 {
		return true
	}

	return false
}

func (s *Stack) Items() []int {
	return s.items
}

func (s *Stack) Size() int {
	return s.size
}
