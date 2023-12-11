package coordinate

type Coord struct {
	Row, Col int
}

func (c Coord) IsEqual(c2 Coord) bool {
	return c.Row == c2.Row && c.Col == c2.Col
}
