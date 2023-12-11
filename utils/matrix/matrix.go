package matrix

import (
	"aoc/utils/coordinate"
	"aoc/utils/direction"
	"fmt"
)

type coord = coordinate.Coord

type Matrix[T any] struct {
	Values      [][]T
	RowCount    int
	ColumnCount int
}

func (m Matrix[T]) Print() {
	for i := 0; i < m.RowCount; i++ {
		for j := 0; j < m.ColumnCount; j++ {
			fmt.Print(m.Values[i][j])
		}
		fmt.Println()
	}
}

func (m Matrix[T]) IsValidCoord(c coord) bool {
	return c.Row >= 0 && c.Row < m.RowCount && c.Col >= 0 && c.Col < m.ColumnCount
}

func (m Matrix[T]) At(c coord) T {
	if !m.IsValidCoord(c) {
		panic(fmt.Sprintf("Invalid matrix coordinate: (%d, %d)", c.Row, c.Col))
	}

	return m.Values[c.Row][c.Col]
}

type offset struct {
	dir       direction.Direction
	rowOffset int
	colOffset int
}

func getOffsetsArray(diagonal bool) []offset {
	if diagonal {
		return []offset{
			{direction.Up, -1, 0},
			{direction.UpRight, -1, 1},
			{direction.Right, 0, 1},
			{direction.RightDown, 1, 1},
			{direction.Down, 1, 0},
			{direction.DownLeft, 1, -1},
			{direction.Left, 0, -1},
			{direction.LeftUp, -1, -1},
		}
	}
	return []offset{
		{direction.Up, -1, 0},
		{direction.Right, 0, 1},
		{direction.Down, 1, 0},
		{direction.Left, 0, -1},
	}
}

// Clockwise iteration of the provided coordinate's neighbors, starting from 12 o'clock.
// Returns all neighboring coordinates that passes the test or an empty array if none
// passed the test.
func (m Matrix[T]) GetValidNeighborCoords(c coord, test func(value T, neighborCoord coord, dir direction.Direction) bool, diagonal bool) (coords []coord) {
	offsets := getOffsetsArray(diagonal)
	neighborCount := len(offsets)
	for i := 0; i < neighborCount; i++ {
		xx := c.Row + offsets[i].rowOffset
		yy := c.Col + offsets[i].colOffset
		newCoord := coord{Row: xx, Col: yy}
		if m.IsValidCoord(newCoord) {
			val := m.Values[xx][yy]
			if test(val, newCoord, offsets[i].dir) {
				coords = append(coords, newCoord)
			}
		}
	}
	return coords
}

// Clockwise iteration of the provided coordinate's neighbors, starting from 12 o'clock.
// Returns a pointer to the first neighboring coordinate that passes the test or `nil`
// if there is no valid neighbor.
func (m Matrix[T]) GetFirstValidNeighbor(c coord, test func(val T, neighborCoord coord, dir direction.Direction) bool, diagonal bool) *coord {
	validNeighborCoords := m.GetValidNeighborCoords(c, test, diagonal)
	if len(validNeighborCoords) > 0 {
		return &validNeighborCoords[0]
	}
	return nil
}
