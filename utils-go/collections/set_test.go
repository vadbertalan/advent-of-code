package collections

import (
	"sort"
	"testing"
)

func TestSet(t *testing.T) {
	t.Run("NewSet", func(t *testing.T) {
		s := NewSet[int]()
		if s == nil {
			t.Fatal("NewSet returned nil")
		}
		if s.Size() != 0 {
			t.Errorf("NewSet should be empty, got size %d", s.Size())
		}
	})

	t.Run("NewSetFromArray", func(t *testing.T) {
		arr := []string{"a", "b", "c", "a"}
		s := NewSetFromArray(arr)
		if s.Size() != 3 {
			t.Errorf("Expected size 3, got %d", s.Size())
		}
		if !s.Has("a") || !s.Has("b") || !s.Has("c") {
			t.Errorf("Set does not contain expected elements")
		}
		if s.Has("d") {
			t.Errorf("Set contains unexpected element 'd'")
		}
	})

	t.Run("Add and Has", func(t *testing.T) {
		s := NewSet[int]()
		s.Add(1)
		if !s.Has(1) {
			t.Errorf("Set should have 1 after adding")
		}
		if s.Has(2) {
			t.Errorf("Set should not have 2")
		}
		s.Add(1) // Add existing element
		if s.Size() != 1 {
			t.Errorf("Adding existing element should not change size, got %d", s.Size())
		}
	})

	t.Run("Remove", func(t *testing.T) {
		s := NewSet[int]()
		s.Add(1)
		s.Add(2)
		s.Remove(1)
		if s.Has(1) {
			t.Errorf("Set should not have 1 after removing")
		}
		if !s.Has(2) {
			t.Errorf("Set should still have 2")
		}
		if s.Size() != 1 {
			t.Errorf("Expected size 1 after removal, got %d", s.Size())
		}
		s.Remove(3) // Remove non-existent element
		if s.Size() != 1 {
			t.Errorf("Removing non-existent element should not change size, got %d", s.Size())
		}
	})

	t.Run("Clear", func(t *testing.T) {
		s := NewSet[int]()
		s.Add(1)
		s.Add(2)
		s.Clear()
		if s.Size() != 0 {
			t.Errorf("Set should be empty after Clear, got size %d", s.Size())
		}
		if s.Has(1) || s.Has(2) {
			t.Errorf("Set should not contain any elements after Clear")
		}
	})

	t.Run("Size", func(t *testing.T) {
		s := NewSet[string]()
		if s.Size() != 0 {
			t.Errorf("Expected size 0, got %d", s.Size())
		}
		s.Add("a")
		if s.Size() != 1 {
			t.Errorf("Expected size 1, got %d", s.Size())
		}
		s.Add("b")
		s.Add("c")
		if s.Size() != 3 {
			t.Errorf("Expected size 3, got %d", s.Size())
		}
		s.Remove("a")
		if s.Size() != 2 {
			t.Errorf("Expected size 2, got %d", s.Size())
		}
	})

	t.Run("GetValues", func(t *testing.T) {
		s := NewSet[int]()
		s.Add(1)
		s.Add(2)
		s.Add(3)
		values := s.GetValues()
		if len(values) != 3 {
			t.Errorf("Expected 3 values, got %d", len(values))
		}
		// Sort for consistent comparison
		sort.Ints(values)
		expected := []int{1, 2, 3}
		for i, v := range values {
			if v != expected[i] {
				t.Errorf("Expected value %d at index %d, got %d", expected[i], i, v)
			}
		}

		emptySet := NewSet[int]()
		emptyValues := emptySet.GetValues()
		if len(emptyValues) != 0 {
			t.Errorf("Expected 0 values for empty set, got %d", len(emptyValues))
		}
	})

	t.Run("AddMulti", func(t *testing.T) {
		s := NewSet[int]()
		s.AddMulti(1, 2, 3, 1)
		if s.Size() != 3 {
			t.Errorf("Expected size 3, got %d", s.Size())
		}
		if !s.Has(1) || !s.Has(2) || !s.Has(3) {
			t.Errorf("Set does not contain expected elements after AddMulti")
		}
	})

	t.Run("Filter", func(t *testing.T) {
		s := NewSetFromArray([]int{1, 2, 3, 4, 5})
		evenFilter := func(v int) bool {
			return v%2 == 0
		}
		evenSet := s.Filter(evenFilter)
		if evenSet.Size() != 2 {
			t.Errorf("Expected size 2 for even numbers, got %d", evenSet.Size())
		}
		if !evenSet.Has(2) || !evenSet.Has(4) {
			t.Errorf("Filtered set does not contain expected even numbers")
		}
		if evenSet.Has(1) || evenSet.Has(3) || evenSet.Has(5) {
			t.Errorf("Filtered set contains unexpected odd numbers")
		}

		emptyFilter := func(v int) bool { return false }
		emptyResult := s.Filter(emptyFilter)
		if emptyResult.Size() != 0 {
			t.Errorf("Expected empty set for always false filter, got size %d", emptyResult.Size())
		}
	})

	t.Run("Union", func(t *testing.T) {
		s1 := NewSetFromArray([]int{1, 2, 3})
		s2 := NewSetFromArray([]int{3, 4, 5})
		unionSet := s1.Union(s2)
		if unionSet.Size() != 5 {
			t.Errorf("Expected union size 5, got %d", unionSet.Size())
		}
		expected := []int{1, 2, 3, 4, 5}
		for _, v := range expected {
			if !unionSet.Has(v) {
				t.Errorf("Union set missing element %d", v)
			}
		}

		s3 := NewSet[int]()
		unionWithEmpty := s1.Union(s3)
		if unionWithEmpty.Size() != 3 {
			t.Errorf("Union with empty set should be original set size, got %d", unionWithEmpty.Size())
		}
	})

	t.Run("Intersect", func(t *testing.T) {
		s1 := NewSetFromArray([]int{1, 2, 3, 4})
		s2 := NewSetFromArray([]int{3, 4, 5, 6})
		intersectSet := s1.Intersect(s2)
		if intersectSet.Size() != 2 {
			t.Errorf("Expected intersect size 2, got %d", intersectSet.Size())
		}
		if !intersectSet.Has(3) || !intersectSet.Has(4) {
			t.Errorf("Intersect set missing expected elements 3, 4")
		}
		if intersectSet.Has(1) || intersectSet.Has(5) {
			t.Errorf("Intersect set contains unexpected elements")
		}

		s3 := NewSetFromArray([]int{7, 8})
		emptyIntersect := s1.Intersect(s3)
		if emptyIntersect.Size() != 0 {
			t.Errorf("Expected empty set for no intersection, got size %d", emptyIntersect.Size())
		}
	})

	t.Run("Difference", func(t *testing.T) {
		s1 := NewSetFromArray([]int{1, 2, 3, 4})
		s2 := NewSetFromArray([]int{3, 4, 5, 6})
		diffSet := s1.Difference(s2)
		if diffSet.Size() != 2 {
			t.Errorf("Expected difference size 2, got %d", diffSet.Size())
		}
		if !diffSet.Has(1) || !diffSet.Has(2) {
			t.Errorf("Difference set missing expected elements 1, 2")
		}
		if diffSet.Has(3) || diffSet.Has(4) || diffSet.Has(5) {
			t.Errorf("Difference set contains unexpected elements")
		}

		s3 := NewSetFromArray([]int{1, 2, 3, 4})
		emptyDiff := s1.Difference(s3)
		if emptyDiff.Size() != 0 {
			t.Errorf("Expected empty set for identical sets difference, got size %d", emptyDiff.Size())
		}
	})

	t.Run("Copy", func(t *testing.T) {
		s := NewSetFromArray([]int{1, 2, 3})
		copiedSet := s.Copy()
		if copiedSet.Size() != s.Size() {
			t.Errorf("Copied set size mismatch, expected %d, got %d", s.Size(), copiedSet.Size())
		}
		if !copiedSet.Has(1) || !copiedSet.Has(2) || !copiedSet.Has(3) {
			t.Errorf("Copied set missing elements")
		}

		// Ensure it's a deep copy (modifying copy doesn't affect original)
		copiedSet.Add(4)
		if s.Has(4) {
			t.Errorf("Modifying copied set affected original set")
		}
		if copiedSet.Size() != 4 {
			t.Errorf("Copied set should have 4 elements, got %d", copiedSet.Size())
		}
		if s.Size() != 3 {
			t.Errorf("Original set should still have 3 elements, got %d", s.Size())
		}
	})
}
