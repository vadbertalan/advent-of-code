package formulae

import (
	"aoc/utils-go/coordinate"
	"math"
)

// Calculate area of a polygon with Shoelace formulae https://en.wikipedia.org/wiki/Shoelace_formula
// See `Other formulas` section. Absolute value is needed because of the order of the vertices.
// Extra explanation: https://www.youtube.com/watch?v=bGWK76_e-LM&t=233s
func CalcAreaShoelace(vertexCoords []coordinate.Coord) float64 {
	A := 0.0
	for i := 0; i < len(vertexCoords); i++ {
		multiplier := 1.0
		if i == 0 {
			multiplier = float64(vertexCoords[1].Col - vertexCoords[len(vertexCoords)-1].Col)
		} else if i == len(vertexCoords)-1 {
			multiplier = float64(vertexCoords[0].Col - vertexCoords[len(vertexCoords)-2].Col)
		} else {
			multiplier = float64(vertexCoords[i+1].Col - vertexCoords[i-1].Col)
		}
		A += float64(vertexCoords[i].Row) * multiplier
	}
	A = math.Abs(float64(A)) / 2
	return A
}

// onSegment checks whether point p lies exactly on line segment (a,b).
func onSegment(a, b, p coordinate.Coord) bool {
	// Check if p is collinear with a->b using cross product
	cross := (b.Row-a.Row)*(p.Col-a.Col) - (b.Col-a.Col)*(p.Row-a.Row)
	if cross != 0 {
		return false
	}

	// Check if p lies within bounding box of segment
	if p.Row < min(a.Row, b.Row) || p.Row > max(a.Row, b.Row) {
		return false
	}
	if p.Col < min(a.Col, b.Col) || p.Col > max(a.Col, b.Col) {
		return false
	}

	return true
}

// Ray Casting algorithm to determine if point is in polygon
// https://people.utm.my/shahabuddin/?p=6277#:~:text=The%20ray%20casting%20algorithm%20involves,point%20is%20outside%20the%20polygon.
func PointInPolygon(poly []coordinate.Coord, p coordinate.Coord) bool {
	n := len(poly)
	if n < 3 {
		return false
	}

	inside := false

	for i := 0; i < n; i++ {
		a := poly[i]
		b := poly[(i+1)%n]

		// Check if point is exactly on the segment
		if onSegment(a, b, p) {
			return true
		}

		// Check if this edge crosses a horizontal ray to the right of p
		intersects := (a.Row > p.Row) != (b.Row > p.Row)
		if intersects {
			xIntersection := float64(a.Col) +
				float64(p.Row-a.Row)*float64(b.Col-a.Col)/float64(b.Row-a.Row)

			if float64(p.Col) < xIntersection {
				inside = !inside
			}
		}
	}

	return inside
}
