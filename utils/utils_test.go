package utils

import (
	"reflect"
	"testing"
)

func TestCloneArray(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{}, []int{}},
		{[]int{42}, []int{42}},
	}

	for _, test := range tests {
		result := CloneArray(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("CloneArray(%v) = %v; want %v", test.input, result, test.expected)
		}
		if &result == &test.input {
			t.Errorf("CloneArray(%v) returned the same reference", test.input)
		}
	}
}

func TestSome(t *testing.T) {
	tests := []struct {
		input     []int
		predicate func(int) bool
		expected  bool
	}{
		{[]int{1, 2, 3, 4, 5}, func(n int) bool { return n > 3 }, true},
		{[]int{1, 2, 3, 4, 5}, func(n int) bool { return n > 5 }, false},
		{[]int{1, 2, 3, 4, 5}, func(n int) bool { return n < 0 }, false},
		{[]int{1, 2, 3, 4, 5}, func(n int) bool { return n == 3 }, true},
		{[]int{}, func(n int) bool { return n > 0 }, false},
	}

	for _, test := range tests {
		result := Some(test.input, test.predicate)
		if result != test.expected {
			t.Errorf("Some(%v, predicate) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		input     []int
		transform func(int) int
		expected  []int
	}{
		{[]int{1, 2, 3}, func(n int) int { return n * 2 }, []int{2, 4, 6}},
		{[]int{1, 2, 3}, func(n int) int { return n + 1 }, []int{2, 3, 4}},
		{[]int{}, func(n int) int { return n * 2 }, []int{}},
		{[]int{1, 2, 3}, func(n int) int { return n - 1 }, []int{0, 1, 2}},
	}

	for _, test := range tests {
		result := Map(test.input, test.transform)
		if len(result) == 0 && len(test.expected) == 0 {
			continue
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Map(%v, transform) = %v; want %v", test.input, result, test.expected)
		}
	}
}
