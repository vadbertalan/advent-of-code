package collections

type Set[T comparable] struct {
	list map[T]struct{} //empty structs occupy 0 memory
}

func (s *Set[T]) Has(v T) bool {
	_, ok := s.list[v]
	return ok
}

func (s *Set[T]) Add(v T) {
	s.list[v] = struct{}{}
}

func (s *Set[T]) Remove(v T) {
	delete(s.list, v)
}

func (s *Set[T]) Clear() {
	s.list = make(map[T]struct{})
}

func (s *Set[T]) Size() int {
	return len(s.list)
}

func (s *Set[T]) GetValues() []T {
	ret := []T{}
	for key := range s.list {
		ret = append(ret, key)
	}
	return ret
}

func NewSet[T comparable]() *Set[T] {
	s := &Set[T]{}
	s.list = make(map[T]struct{})
	return s
}

//optional functionalities

// AddMulti Add multiple values in the set
func (s *Set[T]) AddMulti(list ...T) {
	for _, v := range list {
		s.Add(v)
	}
}

type FilterFunc[T comparable] func(v T) bool

// Filter returns a subset, that contains only the values that satisfies the given predicate P
func (s *Set[T]) Filter(P FilterFunc[T]) *Set[T] {
	res := NewSet[T]()
	for v := range s.list {
		if !P(v) {
			continue
		}
		res.Add(v)
	}
	return res
}

func (s *Set[T]) Union(s2 *Set[T]) *Set[T] {
	res := NewSet[T]()
	for v := range s.list {
		res.Add(v)
	}

	for v := range s2.list {
		res.Add(v)
	}
	return res
}

func (s *Set[T]) Intersect(s2 *Set[T]) *Set[T] {
	res := NewSet[T]()
	for v := range s.list {
		if !s2.Has(v) {
			continue
		}
		res.Add(v)
	}
	return res
}

// Difference returns the subset from s, that doesn't exists in s2 (param)
func (s *Set[T]) Difference(s2 *Set[T]) *Set[T] {
	res := NewSet[T]()
	for v := range s.list {
		if s2.Has(v) {
			continue
		}
		res.Add(v)
	}
	return res
}
