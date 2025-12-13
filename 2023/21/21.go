// https://adventofcode.com/2023/day/21

package main

import (
	"aoc/utils-go"
	"aoc/utils-go/collections"
	"aoc/utils-go/coordinate"
	"aoc/utils-go/direction"
	"aoc/utils-go/matrix"
	"flag"
	"fmt"
	"strconv"
	"time"
)

const aocDay int = 21

type coord = coordinate.Coord
type mat = matrix.Matrix[string]
type dir = direction.Direction

const rock = "#"
const startingPos = "S"

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

type cd struct {
	c coord
	d int
}

func travbfs(m mat, seen coordinate.CoordMap, maxDist int, sc coord, odd int) int {
	test := func(value string, newCoord coord, _ dir) bool {
		return !seen.ContainsCoord(newCoord) && value != rock && value == "O"
	}

	count := 0
	dist := 0
	if odd == 1 {
		dist = 1
	}

	q := collections.Queue[cd]{}
	q.Append(cd{sc, dist})

	for !q.IsEmpty() {
		node := q.Pop()
		c := node.c
		d := node.d

		if seen.ContainsCoord(c) {
			continue
		}

		seen.Add(c)

		m.Set(c, fmt.Sprint(d))

		count++

		if d == maxDist+odd {
			continue
		}

		neighborCoords := m.GetValidNeighborCoords(c, test, false)
		for _, nc := range neighborCoords {
			q.Append(cd{nc, d + 1})
		}
	}

	return count
}

func ff(m mat, seen coordinate.CoordMap, c coord, inBoundary func(c coord) bool) int {
	if seen.ContainsCoord(c) {
		return 0
	}

	seen.Add(c)

	test := func(value string, newCoord coord, _ dir) bool {
		return !seen.ContainsCoord(newCoord) && value != rock && inBoundary(newCoord)
	}

	count := 1
	neighbors := m.GetValidNeighborCoords(c, test, false)
	for _, n := range neighbors {
		count += ff(m, seen, n, inBoundary)
	}
	return count
}

func fill(m mat, stepCount int, sc coord, odd int) int {
	inBoundary := func(c coord) bool {
		return c.ManhattanDist(sc) <= stepCount
	}

	result := 0
	m.Set(sc, "0")
	seen := coordinate.CoordMap{}
	_ = ff(m, seen, sc, inBoundary)

	for i := 0; i < m.RowCount; i++ {
		for j := 0; j < m.ColumnCount; j++ {
			if seen.ContainsRowCol(i, j) {
				m.Values[i][j] = "O"
			}
		}
	}

	seen2 := coordinate.CoordMap{}
	_ = travbfs(m, seen2, stepCount, sc, odd)

	// m.PrintlnWithSpacing()

	for i := 0; i < m.RowCount; i++ {
		for j := 0; j < m.ColumnCount; j++ {
			_, err := strconv.Atoi(m.Values[i][j])
			// if is number
			if err == nil {
				c := coord{Row: i, Col: j}
				scValue, _ := strconv.Atoi(m.At(sc))
				if (c.ManhattanDist(sc)+scValue)%2 == 0 {
					result++
				}
			}
		}
	}

	return result
}

func First(lines []string, stepCount int) (strigifiedResult string) {
	fmt.Println("--- First ---")

	m := matrix.Matrix[string]{
		Values:      make([][]string, len(lines)),
		RowCount:    len(lines),
		ColumnCount: len(lines[0]),
	}

	var sc coord

	// Parse input
	for row, line := range lines {
		m.Values[row] = make([]string, len(line))
		for col, c := range line {
			m.Values[row][col] = string(c)

			if m.Values[row][col] == startingPos {
				sc = coord{Row: row, Col: col}
			}
		}
	}

	strigifiedResult = fmt.Sprint(fill(m, 64, sc, 0))
	return strigifiedResult
}

// wrong 43
// wrong 520
// wrong 2281, too low
// good 3689

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func Second(lines []string, stepCount int) (strigifiedResult string) {
	fmt.Println("\n--- Second ---")

	result := 0

	m := matrix.Matrix[string]{
		Values:      make([][]string, len(lines)),
		RowCount:    len(lines),
		ColumnCount: len(lines[0]),
	}

	// Parse input
	for row, line := range lines {
		m.Values[row] = make([]string, len(line))
		for col, c := range line {
			m.Values[row][col] = string(c)
		}
	}

	mid0 := fill(m.Clone(), 130, coord{Row: 65, Col: 65}, 0)
	mid1 := fill(m.Clone(), 130, coord{Row: 65, Col: 65}, 1)

	n := stepCount / 131

	top := fill(m.Clone(), 130, coord{Row: 130, Col: 65}, 0)
	down := fill(m.Clone(), 130, coord{Row: 0, Col: 65}, 0)
	left := fill(m.Clone(), 130, coord{Row: 65, Col: 0}, 0)
	right := fill(m.Clone(), 130, coord{Row: 65, Col: 130}, 0)

	blsm0 := fill(m.Clone(), 65, coord{Row: 130, Col: 0}, 0)
	brsm0 := fill(m.Clone(), 65, coord{Row: 130, Col: 130}, 0)
	tlsm0 := fill(m.Clone(), 65, coord{Row: 0, Col: 0}, 0)
	trsm0 := fill(m.Clone(), 65, coord{Row: 0, Col: 130}, 0)

	blbig1 := fill(m.Clone(), 195, coord{Row: 130, Col: 0}, 1)
	brbig1 := fill(m.Clone(), 195, coord{Row: 130, Col: 130}, 1)
	tlbig1 := fill(m.Clone(), 195, coord{Row: 0, Col: 0}, 1)
	trbig1 := fill(m.Clone(), 195, coord{Row: 0, Col: 130}, 1)

	// Rombus theory all over the place...
	result = (n*n)*mid0 + (n-1)*(n-1)*mid1 + top + down + left + right + (n-1)*blbig1 + (n-1)*brbig1 + (n-1)*tlbig1 + (n-1)*trbig1 + n*blsm0 + n*brsm0 + n*tlsm0 + n*trsm0

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// 152614178271301 too low
// 610444637284029 too high
// 609871691720601 too low
// 609953542705201 wrong
// 610164189202186 wrong...
// 610158160254036 wrong...
// 610158187362102 good!!

func main() {
	stepCountP := flag.Int("sc", 64, "Specify if you want to modify the step count. Example step count is 6")
	// srowP := flag.Int("srow", 65, "starting row")
	// scolP := flag.Int("scol", 65, "starting col")

	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(1)

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	result := First(lines, *stepCountP)
	fmt.Println(result)

	result = Second(lines, 26501365)
	fmt.Println(result)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
