package coordinate

import "math"

type Coord3 struct {
	X, Y, Z int
}

func (c1 Coord3) IsEqual(c2 Coord3) bool {
	return c1.X == c2.X && c1.Y == c2.Y && c1.Z == c2.Z
}

func (c1 Coord3) EuclideanDist3(c2 Coord3) float64 {
	return math.Sqrt(float64((c1.X-c2.X)*(c1.X-c2.X) + (c1.Y-c2.Y)*(c1.Y-c2.Y) + (c1.Z-c2.Z)*(c1.Z-c2.Z)))
}
