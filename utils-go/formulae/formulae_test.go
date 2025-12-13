package formulae

import (
	"aoc/utils-go/coordinate"
	"testing"
)

func TestPointInPolygon(t *testing.T) {
	tests := []struct {
		name     string
		poly     []coordinate.Coord
		point    coordinate.Coord
		expected bool
	}{
		{
			name: "point inside square",
			poly: []coordinate.Coord{
				{Row: 0, Col: 0},
				{Row: 0, Col: 10},
				{Row: 10, Col: 10},
				{Row: 10, Col: 0},
			},
			point:    coordinate.Coord{Row: 5, Col: 5},
			expected: true,
		},
		{
			name: "point outside square",
			poly: []coordinate.Coord{
				{Row: 0, Col: 0},
				{Row: 0, Col: 10},
				{Row: 10, Col: 10},
				{Row: 10, Col: 0},
			},
			point:    coordinate.Coord{Row: 15, Col: 15},
			expected: false,
		},
		{
			name: "point on edge 1",
			poly: []coordinate.Coord{
				{Row: 0, Col: 0},
				{Row: 0, Col: 10},
				{Row: 10, Col: 10},
				{Row: 10, Col: 0},
			},
			point:    coordinate.Coord{Row: 5, Col: 0},
			expected: true,
		},
		{
			name: "point on edge 2",
			poly: []coordinate.Coord{
				{Row: 0, Col: 0},
				{Row: 0, Col: 10},
				{Row: 10, Col: 10},
				{Row: 10, Col: 0},
			},
			point:    coordinate.Coord{Row: 10, Col: 5},
			expected: true,
		},
		{
			name: "point on edge 3",
			poly: []coordinate.Coord{
				{Row: 0, Col: 0},
				{Row: 0, Col: 10},
				{Row: 10, Col: 10},
				{Row: 10, Col: 0},
			},
			point:    coordinate.Coord{Row: 5, Col: 10},
			expected: true,
		},
		{
			name: "point on edge 4",
			poly: []coordinate.Coord{
				{Row: 0, Col: 0},
				{Row: 0, Col: 10},
				{Row: 10, Col: 10},
				{Row: 10, Col: 0},
			},
			point:    coordinate.Coord{Row: 0, Col: 5},
			expected: true,
		},
		{
			name: "point at vertex 1",
			poly: []coordinate.Coord{
				{Row: 0, Col: 0},
				{Row: 0, Col: 10},
				{Row: 10, Col: 10},
				{Row: 10, Col: 0},
			},
			point:    coordinate.Coord{Row: 0, Col: 0},
			expected: true,
		},
		{
			name: "point at vertex 2",
			poly: []coordinate.Coord{
				{Row: 0, Col: 0},
				{Row: 0, Col: 10},
				{Row: 10, Col: 10},
				{Row: 10, Col: 0},
			},
			point:    coordinate.Coord{Row: 10, Col: 0},
			expected: true,
		},
		{
			name: "point at vertex 3",
			poly: []coordinate.Coord{
				{Row: 0, Col: 0},
				{Row: 0, Col: 10},
				{Row: 10, Col: 10},
				{Row: 10, Col: 0},
			},
			point:    coordinate.Coord{Row: 10, Col: 10},
			expected: true,
		},
		{
			name: "point at vertex 4",
			poly: []coordinate.Coord{
				{Row: 0, Col: 0},
				{Row: 0, Col: 10},
				{Row: 10, Col: 10},
				{Row: 10, Col: 0},
			},
			point:    coordinate.Coord{Row: 0, Col: 10},
			expected: true,
		},
		{
			name: "point inside triangle",
			poly: []coordinate.Coord{
				{Row: 0, Col: 0},
				{Row: 10, Col: 5},
				{Row: 0, Col: 10},
			},
			point:    coordinate.Coord{Row: 3, Col: 5},
			expected: true,
		},
		{
			name: "point on triangle vertex 1",
			poly: []coordinate.Coord{
				{Row: 0, Col: 0},
				{Row: 10, Col: 5},
				{Row: 0, Col: 10},
			},
			point:    coordinate.Coord{Row: 0, Col: 0},
			expected: true,
		},
		{
			name: "point on triangle vertex 2",
			poly: []coordinate.Coord{
				{Row: 0, Col: 0},
				{Row: 10, Col: 5},
				{Row: 0, Col: 10},
			},
			point:    coordinate.Coord{Row: 10, Col: 5},
			expected: true,
		},
		{
			name:     "polygon with less than 3 vertices",
			poly:     []coordinate.Coord{{Row: 0, Col: 0}, {Row: 1, Col: 1}},
			point:    coordinate.Coord{Row: 0, Col: 0},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PointInPolygon(tt.poly, tt.point)
			if result != tt.expected {
				t.Errorf("PointInPolygon() = %v; want %v", result, tt.expected)
			}
		})
	}
}
