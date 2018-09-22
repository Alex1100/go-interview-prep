package stack

type Stack struct {
	Items []int
	Size  int
}

func InitStack() *Stack {
	var items = make([]int, 0)
	return &Stack{
		Items: items,
		Size:  0,
	}
}

func (s *Stack) Insert(item int) error {
	s.Items = append(s.Items, item)
	s.Size++
	return nil
}

func (s *Stack) Pop() (int, error) {
	removed := s.Items[len(s.Items)-1]
	if len(s.Items) > 1 {
		s.Items = s.Items[0 : len(s.Items)-1]
	} else {
		s.Items = make([]int, 0)
	}

	s.Size--
	return removed, nil
}

func (s *Stack) Front() int {
	return s.Items[len(s.Items)-1]
}

func (s *Stack) IsEmpty() bool {
	if s.Size == 0 {
		return true
	}

	return false
}
