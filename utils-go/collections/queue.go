package collections

import "fmt"

type Queue[T any] struct {
	items []T
}

func NewQueue[T any]() *Queue[T] {
	q := &Queue[T]{}
	return q
}

func (s *Queue[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Queue[T]) First() T {
	if s.IsEmpty() {
		print("Queue is empty")
	}
	return s.items[0]
}

func (s *Queue[T]) FirstSafe() (*T, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("Queue is empty")
	}
	return &s.items[0], nil
}

func (s *Queue[T]) Append(newItem T) {
	s.items = append(s.items, newItem)
}

func (s *Queue[T]) Pop() T {
	if s.IsEmpty() {
		panic("Queue is empty")
	}
	ret := s.First()
	s.items = s.items[1:]
	return ret
}

func (s *Queue[T]) PopSafe() (*T, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("Queue is empty")
	}
	ret, _ := s.FirstSafe()
	s.items = s.items[1:]
	return ret, nil
}

func (s *Queue[T]) Print() {
	for _, item := range s.items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}
