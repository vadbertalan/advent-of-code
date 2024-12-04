package formulae

import (
	"aoc/utils/coordinate"
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
