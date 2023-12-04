package utils

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("Stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("Stack is empty")
	}
	ret, _ := s.Top()
	s.items = s.items[:len(s.items)-1]
	return ret, nil
}

func (s *Stack) Print() {
	for _, item := range s.items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}
