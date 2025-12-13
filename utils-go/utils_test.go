package utils

import (
	"reflect"
	"slices"
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
func TestAbs(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{input: 5, expected: 5},
		{input: -5, expected: 5},
		{input: 0, expected: 0},
		{input: -123456, expected: 123456},
		{input: 123456, expected: 123456},
	}

	for _, test := range tests {
		result := Abs(test.input)
		if result != test.expected {
			t.Errorf("Abs(%d) = %d; want %d", test.input, result, test.expected)
		}
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{input: 0, expected: 1},
		{input: 1, expected: 1},
		{input: 2, expected: 2},
		{input: 3, expected: 6},
		{input: 4, expected: 24},
		{input: 5, expected: 120},
		{input: 10, expected: 3628800},
	}

	for _, test := range tests {
		result := Factorial(test.input)
		if result != test.expected {
			t.Errorf("Factorial(%d) = %d; want %d", test.input, result, test.expected)
		}
	}
}

func TestFilterDuplicates(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 2, 3, 4, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{1, 1, 1, 1, 1}, expected: []int{1}},
		{input: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{}, expected: []int{}},
		{input: []int{5, 5, 5, 5, 5, 5}, expected: []int{5}},
	}

	testsWithStrings := []struct {
		input    []string
		expected []string
	}{
		{input: []string{"a", "b", "b", "c", "d", "d", "e"}, expected: []string{"a", "b", "c", "d", "e"}},
		{input: []string{"a", "a", "a", "a", "a"}, expected: []string{"a"}},
	}

	for _, test := range slices.Concat(tests) {
		result := FilterDuplicates(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("FilterDuplicates(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
	for _, test := range slices.Concat(testsWithStrings) {
		result := FilterDuplicates(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("FilterDuplicates(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestAtoiEdgeCases(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{input: "2147483647", expected: 2147483647},   // max int32
		{input: "-2147483648", expected: -2147483648}, // min int32
		{input: "00123", expected: 123},               // leading zeros
		{input: "-0", expected: 0},                    // negative zero
		{input: "1", expected: 1},                     // single digit
	}

	for _, test := range tests {
		result := Atoi(test.input)
		if result != test.expected {
			t.Errorf("Atoi(%q) = %d; want %d", test.input, result, test.expected)
		}
	}
}

func TestAtoiPanic(t *testing.T) {
	tests := []string{
		"abc",
		"12a34",
		"1.5",
		"",
		"  ",
		"1 2 3",
	}

	for _, input := range tests {
		t.Run(input, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Atoi(%q) should panic but didn't", input)
				}
			}()
			Atoi(input)
		})
	}
}

func TestCloneMatrix(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected [][]int
	}{
		{[][]int{{1, 2}, {3, 4}}, [][]int{{1, 2}, {3, 4}}},
		{[][]int{{1}}, [][]int{{1}}},
		{[][]int{}, [][]int{}},
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
		{[][]int{{0}, {0}, {0}}, [][]int{{0}, {0}, {0}}},
	}

	for _, test := range tests {
		result := CloneMatrix(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("CloneMatrix(%v) = %v; want %v", test.input, result, test.expected)
		}
		if len(result) > 0 && len(test.input) > 0 {
			if &result == &test.input {
				t.Errorf("CloneMatrix(%v) returned the same reference", test.input)
			}
			if &result[0] == &test.input[0] {
				t.Errorf("CloneMatrix(%v) returned the same row reference", test.input)
			}
		}
	}
}
