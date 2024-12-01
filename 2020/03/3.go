// https://adventofcode.com/2020/day/3

package main

import (
	"aoc/utils"
	"aoc/utils/coordinate"
	"aoc/utils/matrix"
	"fmt"
	"time"
)

const aocDay int = 3

type coord = coordinate.Coord

const treeChar = '#'

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func First(lines []string) (strigifiedResult string) {
	fmt.Println("--- First ---")

	result := 0

	// Declare matrix
	m := matrix.Matrix[string]{
		Values:      make([][]string, len(lines)),
		RowCount:    len(lines),
		ColumnCount: len(lines[0]),
	}

	// Read matrix
	for row, line := range lines {
		m.Values[row] = make([]string, len(line))
		for col, c := range line {
			m.Values[row][col] = string(c)
		}
	}

	currentCoord := coord{Row: 0, Col: 0}

	// while currentCoord Y did not pass the last row
	for currentCoord.Row < m.RowCount {
		currentCoord.Col += 3
		currentCoord.Row += 1

		if currentCoord.Row >= m.RowCount {
			break
		}

		if currentCoord.Col >= m.ColumnCount {
			currentCoord.Col -= m.ColumnCount
		}

		// if hit a tree, increment
		if m.Values[currentCoord.Row][currentCoord.Col] == string(treeChar) {
			result++
		}
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// correct 274, example correct 7

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func Second(lines []string) (strigifiedResult string) {
	fmt.Println("\n--- Second ---")

	// Declare matrix
	m := matrix.Matrix[string]{
		Values:      make([][]string, len(lines)),
		RowCount:    len(lines),
		ColumnCount: len(lines[0]),
	}

	// Read matrix
	for row, line := range lines {
		m.Values[row] = make([]string, len(line))
		for col, c := range line {
			m.Values[row][col] = string(c)
		}
	}

	// Solve

	result := 1

	slopeVariants := []struct {
		Right int
		Down  int
	}{
		{Right: 1, Down: 1},
		{Right: 3, Down: 1},
		{Right: 5, Down: 1},
		{Right: 7, Down: 1},
		{Right: 1, Down: 2},
	}

	for _, slope := range slopeVariants {
		currentCoord := coord{Row: 0, Col: 0}

		localResult := 0

		// while currentCoord Y did not pass the last row
		for currentCoord.Row < m.RowCount {
			currentCoord.Col += slope.Right
			currentCoord.Row += slope.Down

			if currentCoord.Row >= m.RowCount {
				break
			}

			if currentCoord.Col >= m.ColumnCount {
				currentCoord.Col -= m.ColumnCount
			}

			// if hit a tree, increment
			if m.Values[currentCoord.Row][currentCoord.Col] == string(treeChar) {
				localResult++
			}
		}

		result *= localResult
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// correct 6050183040, example correct 336

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(1)

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	result := First(lines)
	fmt.Println(result)

	result = Second(lines)
	fmt.Println(result)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
