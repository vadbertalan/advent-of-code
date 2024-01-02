// https://adventofcode.com/2023/day/14

package main

import (
	"aoc/utils"
	"aoc/utils/collections"
	"aoc/utils/coordinate"
	"aoc/utils/direction"
	"aoc/utils/matrix"
	"fmt"
	"slices"
	"time"
)

const aocDay int = 14

type coord = coordinate.Coord
type mat = matrix.Matrix[string]
type dir = direction.Direction

const roundedRock = "O"
const fixRock = "#"
const emptySpace = "."

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func first(lines []string) {
	fmt.Println("--- First ---")

	result := 0

	m := matrix.Matrix[string]{
		Values:      make([][]string, len(lines)),
		RowCount:    len(lines),
		ColumnCount: len(lines[0]),
	}

	rrCoords := []coord{}

	for row, line := range lines {
		m.Values[row] = make([]string, len(line))
		for col, c := range line {
			m.Values[row][col] = string(c)

			if m.Values[row][col] == roundedRock {
				rrCoords = append(rrCoords, coord{Row: row, Col: col})
			}
		}
	}

	newCoords := []coord{}

	for _, rrCoord := range rrCoords {
		rrCanMove := func(newCoord coord) bool {
			return m.IsValidCoord(newCoord) && m.At(newCoord) == emptySpace
		}

		oldCoord := rrCoord
		newCoord := coord{Row: rrCoord.Row - 1, Col: rrCoord.Col}
		for rrCanMove(newCoord) {
			m.Set(newCoord, roundedRock)
			m.Set(oldCoord, emptySpace)
			oldCoord = newCoord
			newCoord = coord{Row: newCoord.Row - 1, Col: newCoord.Col}
		}

		newCoords = append(newCoords, oldCoord)
	}

	for _, newCoord := range newCoords {
		result += m.RowCount - newCoord.Row
	}

	fmt.Println(result)
}

const cycleCount = 1_000_000_000

func cmpWest(c1, c2 coord) int {
	if c1.Col < c2.Col {
		return -1
	}
	if c1.Col > c2.Col {
		return 1
	}
	return cmpNorth(c1, c2)
}

func cmpSouth(c1, c2 coord) int {
	if c1.Row > c2.Row {
		return -1
	}
	if c1.Row < c2.Row {
		return 1
	}
	return cmpWest(c1, c2)
}

func cmpEast(c1, c2 coord) int {
	if c1.Col > c2.Col {
		return -1
	}
	if c1.Col < c2.Col {
		return 1
	}
	return cmpNorth(c1, c2)
}

func cmpNorth(c1, c2 coord) int {
	if c1.Row < c2.Row {
		return -1
	}
	if c1.Row > c2.Row {
		return 1
	}
	return cmpWest(c1, c2)
}

func reorderRRCoordsByDir(rrCoords []coord, d dir) []coord {
	sortFuncs := map[dir]func(c1, c2 coord) int{
		direction.Up:    cmpNorth,
		direction.Left:  cmpWest,
		direction.Down:  cmpSouth,
		direction.Right: cmpEast,
	}
	slices.SortFunc(rrCoords, sortFuncs[d])
	return rrCoords
}

func shift(m mat, initCoords []coord, d dir) []coord {
	newCoords := []coord{}

	for _, initCoord := range initCoords {
		canMoveTo := func(newCoord coord) bool {
			return m.IsValidCoord(newCoord) && m.At(newCoord) == emptySpace
		}

		oldCoord := initCoord
		newCoord := initCoord.GetNewCoord(d)
		for canMoveTo(newCoord) {
			m.Set(newCoord, roundedRock)
			m.Set(oldCoord, emptySpace)
			oldCoord = newCoord
			newCoord = newCoord.GetNewCoord(d)
		}

		newCoords = append(newCoords, oldCoord)
	}

	return newCoords
}

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func getSeenKey(coords []coord) string {
	ret := ""
	for _, coord := range coords {
		ret += fmt.Sprintf("%d-%d,", coord.Row, coord.Col)
	}
	return ret
}

func second(lines []string) {
	fmt.Println("\n--- Second ---")

	m := matrix.Matrix[string]{
		Values:      make([][]string, len(lines)),
		RowCount:    len(lines),
		ColumnCount: len(lines[0]),
	}

	rrCoords := []coord{}

	for row, line := range lines {
		m.Values[row] = make([]string, len(line))
		for col, c := range line {
			m.Values[row][col] = string(c)

			if m.Values[row][col] == roundedRock {
				rrCoords = append(rrCoords, coord{Row: row, Col: col})
			}
		}
	}

	// First cycle
	rrCoords = reorderRRCoordsByDir(rrCoords, direction.Up)
	rrCoords = shift(m, rrCoords, direction.Up)

	rrCoords = reorderRRCoordsByDir(rrCoords, direction.Left)
	rrCoords = shift(m, rrCoords, direction.Left)

	rrCoords = reorderRRCoordsByDir(rrCoords, direction.Down)
	rrCoords = shift(m, rrCoords, direction.Down)

	rrCoords = reorderRRCoordsByDir(rrCoords, direction.Right)
	rrCoords = shift(m, rrCoords, direction.Right)

	// Calculate result
	score := 0
	for _, newCoord := range rrCoords {
		score += m.RowCount - newCoord.Row
	}

	// Init data structures used for registering scores
	scores := []int{}
	seen := collections.NewSet[string]()
	curSeenKey := getSeenKey(rrCoords)
	curSeenKeys := []string{}

	// Iterate from second cycle
	for !seen.Has(curSeenKey) {
		// fmt.Println("---- Cycle", seen.Size())

		// Calculate result
		score = 0
		for _, newCoord := range rrCoords {
			score += m.RowCount - newCoord.Row
		}

		// Register cycle
		seen.Add(curSeenKey)
		scores = append(scores, score)
		curSeenKeys = append(curSeenKeys, curSeenKey)

		// Advance matrix
		rrCoords = reorderRRCoordsByDir(rrCoords, direction.Up)
		rrCoords = shift(m, rrCoords, direction.Up)

		rrCoords = reorderRRCoordsByDir(rrCoords, direction.Left)
		rrCoords = shift(m, rrCoords, direction.Left)

		rrCoords = reorderRRCoordsByDir(rrCoords, direction.Down)
		rrCoords = shift(m, rrCoords, direction.Down)

		rrCoords = reorderRRCoordsByDir(rrCoords, direction.Right)
		rrCoords = shift(m, rrCoords, direction.Right)

		curSeenKey = getSeenKey(rrCoords)
	}

	indexOfFirst := -1
	for i, csk := range curSeenKeys {
		if csk == curSeenKey {
			indexOfFirst = i
			break
		}
	}
	realCycleLen := seen.Size() - indexOfFirst
	x := scores[indexOfFirst:]

	a := cycleCount - indexOfFirst - 1
	fmt.Println(x[a%realCycleLen])
}

// tried second with 95429, too low
// Your puzzle answer was 105008.

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(1)

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	first(lines)

	second(lines)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
