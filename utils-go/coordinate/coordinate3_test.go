package coordinate

import (
	"math"
	"testing"
)

func TestEuclideanDist3(t *testing.T) {
	tests := []struct {
		c1, c2   Coord3
		expected float64
	}{
		{Coord3{0, 0, 0}, Coord3{0, 0, 0}, 0},
		{Coord3{1, 0, 0}, Coord3{0, 0, 0}, 1},
		{Coord3{0, 2, 0}, Coord3{0, 0, 0}, 2},
		{Coord3{0, 0, 3}, Coord3{0, 0, 0}, 3},
		{Coord3{1, 2, 2}, Coord3{4, 6, 5}, math.Sqrt(34)},
		{Coord3{-1, -2, -3}, Coord3{1, 2, 3}, math.Sqrt(56)},
	}

	for _, tt := range tests {
		result := tt.c1.EuclideanDist3(tt.c2)
		if math.Abs(result-tt.expected) > 1e-9 {
			t.Errorf("EuclideanDist3(%v, %v) = %v; want %v", tt.c1, tt.c2, result, tt.expected)
		}
	}
}

func TestIsEqual(t *testing.T) {
	tests := []struct {
		c1, c2   Coord3
		expected bool
	}{
		{Coord3{0, 0, 0}, Coord3{0, 0, 0}, true},
		{Coord3{1, 2, 3}, Coord3{1, 2, 3}, true},
		{Coord3{-1, -2, -3}, Coord3{-1, -2, -3}, true},
		{Coord3{1, 2, 3}, Coord3{3, 2, 1}, false},
		{Coord3{0, 0, 0}, Coord3{0, 1, 0}, false},
		{Coord3{0, 0, 0}, Coord3{0, 0, 1}, false},
		{Coord3{0, 0, 0}, Coord3{1, 0, 0}, false},
	}

	for _, tt := range tests {
		result := tt.c1.IsEqual(tt.c2)
		if result != tt.expected {
			t.Errorf("IsEqual(%v, %v) = %v; want %v", tt.c1, tt.c2, result, tt.expected)
		}
	}
}
