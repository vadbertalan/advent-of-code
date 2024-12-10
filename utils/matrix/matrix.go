package matrix

import (
	"aoc/utils/coordinate"
	"aoc/utils/direction"
	"fmt"
)

type coord = coordinate.Coord

type Matrix[T comparable] struct {
	Values      [][]T
	RowCount    int
	ColumnCount int
}

func (m Matrix[T]) Clone() Matrix[T] {
	mClone := Matrix[T]{RowCount: m.RowCount, ColumnCount: m.ColumnCount, Values: make([][]T, len(m.Values))}
	for i := 0; i < m.RowCount; i++ {
		mClone.Values[i] = make([]T, len(m.Values[i]))
		for j := 0; j < m.ColumnCount; j++ {
			mClone.Values[i][j] = m.Values[i][j]
		}
	}
	return mClone
}

func (m Matrix[T]) PrintWithSmallSpacing() {
	for i := 0; i < m.RowCount; i++ {
		for j := 0; j < m.ColumnCount; j++ {
			fmt.Printf("%2v", m.Values[i][j])
		}
		fmt.Println()
	}
}

func (m Matrix[T]) PrintWithSpacing() {
	for i := 0; i < m.RowCount; i++ {
		for j := 0; j < m.ColumnCount; j++ {
			fmt.Printf("%5v", m.Values[i][j])
		}
		fmt.Println()
	}
}

func (m Matrix[T]) Print() {
	for i := 0; i < m.RowCount; i++ {
		for j := 0; j < m.ColumnCount; j++ {
			fmt.Print(m.Values[i][j])
		}
		fmt.Println()
	}
}

func (m Matrix[T]) Println() {
	m.Print()
	fmt.Println()
}

func (m Matrix[T]) PrintlnWithSpacing() {
	m.PrintWithSpacing()
	fmt.Println()
}

func (m Matrix[T]) PrintlnWithSmallSpacing() {
	m.PrintWithSmallSpacing()
	fmt.Println()
}

func (m Matrix[T]) PrintlnWithOverride(test func(i, j int, value T) bool, override string) {
	for i := 0; i < m.RowCount; i++ {
		for j := 0; j < m.ColumnCount; j++ {
			if test(i, j, m.Values[i][j]) {
				fmt.Printf("%2v", override)
			} else {
				fmt.Printf("%2v", m.Values[i][j])
			}
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

func (m Matrix[T]) Set(c coord, value T) {
	if !m.IsValidCoord(c) {
		panic(fmt.Sprintf("Invalid matrix coordinate: (%d, %d)", c.Row, c.Col))
	}

	m.Values[c.Row][c.Col] = value
}

func (m Matrix[T]) Count(value T) (count int) {
	for i := 0; i < m.RowCount; i++ {
		for j := 0; j < m.ColumnCount; j++ {
			if m.Values[i][j] == value {
				count++
			}
		}
	}
	return count
}

// Clockwise iteration of the provided coordinate's neighbors, starting from 12 o'clock.
// Returns all neighboring coordinates that passes the test or an empty array if none
// passed the test.
func (m Matrix[T]) GetValidNeighborCoords(c coord, test func(value T, neighborCoord coord, dir direction.Direction) bool, diagonal bool) (coords []coord) {
	offsets := coordinate.GetOffsetsArray(diagonal)
	neighborCount := len(offsets)
	for i := 0; i < neighborCount; i++ {
		xx := c.Row + offsets[i].RowOffset
		yy := c.Col + offsets[i].ColOffset
		newCoord := coord{Row: xx, Col: yy}
		if m.IsValidCoord(newCoord) {
			val := m.Values[xx][yy]
			if test(val, newCoord, offsets[i].Dir) {
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

func ParseStringMatrix(lines []string) Matrix[string] {
	if len(lines) == 0 {
		return Matrix[string]{}
	}

	m := Matrix[string]{
		Values:      make([][]string, len(lines)),
		RowCount:    len(lines),
		ColumnCount: len(lines[0]),
	}

	for row, line := range lines {
		m.Values[row] = make([]string, len(line))
		for col, c := range line {
			m.Values[row][col] = string(c)
		}
	}

	return m
}

func ParseDigitMatrix(lines []string) Matrix[int] {
	if len(lines) == 0 {
		return Matrix[int]{}
	}

	m := Matrix[int]{
		Values:      make([][]int, len(lines)),
		RowCount:    len(lines),
		ColumnCount: len(lines[0]),
	}

	for row, line := range lines {
		m.Values[row] = make([]int, len(line))
		for col, c := range line {
			m.Values[row][col] = int(c - '0')
		}
	}

	return m
}

func (m Matrix[T]) GetAllCoordsWhich(test func(value T) bool) (coords []coord) {
	for i := 0; i < m.RowCount; i++ {
		for j := 0; j < m.ColumnCount; j++ {
			if test(m.Values[i][j]) {
				coords = append(coords, coord{Row: i, Col: j})
			}
		}
	}
	return coords
}

func (m Matrix[T]) IsPathBetween(currentCoord coordinate.Coord, endCoord coordinate.Coord, test func(currentCoord, nextCoord coordinate.Coord, currentValue, nextValue T) bool) bool {
	seenCoordMap := make(coordinate.CoordMap)
	return m.isPathBetweenInternal(currentCoord, endCoord, seenCoordMap, test)
}

func (m Matrix[T]) isPathBetweenInternal(currentCoord coordinate.Coord, endCoord coordinate.Coord, seenCoordMap coordinate.CoordMap, test func(currentCoord, nextCoord coordinate.Coord, currentValue, nextValue T) bool) bool {
	if currentCoord == endCoord {
		return true
	}

	seenCoordMap.Add(currentCoord)

	dirOffsets := coordinate.GetOffsetsArray(false)
	for _, dirOffset := range dirOffsets {
		nextCoord := currentCoord.GetNewCoord(dirOffset.Dir)

		if !m.IsValidCoord(nextCoord) {
			continue
		}

		if !seenCoordMap.ContainsCoord(nextCoord) &&
			test(currentCoord, nextCoord, m.At(currentCoord), m.At(nextCoord)) {
			if m.isPathBetweenInternal(nextCoord, endCoord, seenCoordMap, test) {
				return true
			}
		}
	}

	return false
}

func (m Matrix[T]) CountPathsBetween(currentCoord coordinate.Coord, endCoord coordinate.Coord, test func(currentCoord, nextCoord coordinate.Coord, currentValue, nextValue T) bool) int {
	seenCoordMap := make(coordinate.CoordMap)
	return m.countPathsBetweenInternal(currentCoord, endCoord, seenCoordMap, test)
}

func (m Matrix[T]) countPathsBetweenInternal(currentCoord coord, endCoord coord, seenCoordMap coordinate.CoordMap, test func(currentCoord, nextCoord coordinate.Coord, currentValue, nextValue T) bool) int {
	if currentCoord == endCoord {
		return 1
	}

	seenCoordMap.Add(currentCoord)

	count := 0
	dirOffsets := coordinate.GetOffsetsArray(false)
	for _, dirOffset := range dirOffsets {
		nextCoord := currentCoord.GetNewCoord(dirOffset.Dir)

		if !m.IsValidCoord(nextCoord) {
			continue
		}

		if !seenCoordMap.ContainsCoord(nextCoord) &&
			test(currentCoord, nextCoord, m.At(currentCoord), m.At(nextCoord)) {
			count += m.countPathsBetweenInternal(nextCoord, endCoord, seenCoordMap.Copy(), test)
		}
	}

	return count
}
