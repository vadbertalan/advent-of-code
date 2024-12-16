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
