package collections

import "fmt"

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Top() T {
	if s.IsEmpty() {
		print("Stack is empty")
	}
	return s.items[len(s.items)-1]
}

func (s *Stack[T]) TopSafe() (*T, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("Stack is empty")
	}
	return &s.items[len(s.items)-1], nil
}

func (s *Stack[T]) Push(data T) {
	s.items = append(s.items, data)
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		panic("Stack is empty")
	}
	ret := s.Top()
	s.items = s.items[:len(s.items)-1]
	return ret
}

func (s *Stack[T]) PopSafe() (*T, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("Stack is empty")
	}
	ret, _ := s.TopSafe()
	s.items = s.items[:len(s.items)-1]
	return ret, nil
}

func (s *Stack[T]) Print() {
	for _, item := range s.items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}
