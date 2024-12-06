// https://adventofcode.com/2024/day/6

package main

import (
	"aoc/utils"
	"aoc/utils/coordinate"
	"aoc/utils/direction"
	"aoc/utils/matrix"
	"fmt"
	"time"
)

const aocDay int = 6

type coord = coordinate.Coord
type mat = matrix.Matrix[string]

const startCoordChar = "^"
const obstacleChar = "#"

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

var coordMap = make(coordinate.CoordMap)

func travel1(m mat, currentCoord coord, dirOffset coordinate.DirOffset) int {
	coordMap.Add(currentCoord)

	newCoord := currentCoord.GetNewCoord(dirOffset.Dir)

	if !m.IsValidCoord(newCoord) {
		return 0
	}

	if m.Values[newCoord.Row][newCoord.Col] == obstacleChar {
		newDirOffset := coordinate.GetClockwise90DegreeNeighborOffset(dirOffset)
		return 1 + travel1(m, currentCoord.GetNewCoord(newDirOffset.Dir), newDirOffset)
	}

	return 1 + travel1(m, newCoord, dirOffset)
}

func First(lines []string) (strigifiedResult string) {
	fmt.Println("--- First ---")

	guardCoord := coord{Row: -1, Col: -1}

	m := matrix.Matrix[string]{
		Values:      make([][]string, len(lines)),
		RowCount:    len(lines),
		ColumnCount: len(lines[0]),
	}

	for row, line := range lines {
		m.Values[row] = make([]string, len(line))
		for col, c := range line {
			m.Values[row][col] = string(c)

			if m.Values[row][col] == startCoordChar {
				guardCoord = coord{Row: row, Col: col}
			}
		}
	}

	coordMap.Clear()
	travel1(m, guardCoord, coordinate.GetOffsetForDir(direction.Up))

	result := len(coordMap)

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// correct answers: 41, 5101

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

var travelMap = make(coordinate.TravelMap)

func isInLoop(m mat, currentCoord coord, dirOffset coordinate.DirOffset, potentialObstacleCoord coord) bool {
	travelMap.Add(currentCoord, dirOffset.Dir)

	newCoord := currentCoord.GetNewCoord(dirOffset.Dir)

	if !m.IsValidCoord(newCoord) {
		return false
	}

	if travelMap.ContainsCoordAndDir(newCoord, dirOffset.Dir) {
		return true
	}

	didHitObstacle := m.Values[newCoord.Row][newCoord.Col] == obstacleChar || newCoord.IsEqual(potentialObstacleCoord)
	if didHitObstacle {
		newDirOffset := coordinate.GetClockwise90DegreeNeighborOffset(dirOffset)

		// We stay in place, do not move AND turn at the same time
		return isInLoop(m, currentCoord, newDirOffset, potentialObstacleCoord)
	}

	return isInLoop(m, newCoord, dirOffset, potentialObstacleCoord)
}

func Second(lines []string) (strigifiedResult string) {
	fmt.Println("\n--- Second ---")

	// ⚔️ Leaving this here to mark a historical battlefield ⚔️
	// debug.SetMaxStack(3000000000)

	guardCoord := coord{Row: -1, Col: -1}

	m := matrix.Matrix[string]{
		Values:      make([][]string, len(lines)),
		RowCount:    len(lines),
		ColumnCount: len(lines[0]),
	}

	for row, line := range lines {
		m.Values[row] = make([]string, len(line))
		for col, c := range line {
			m.Values[row][col] = string(c)

			if m.Values[row][col] == startCoordChar {
				guardCoord = coord{Row: row, Col: col}
			}
		}
	}

	// Setup coordMap of path with function from part 1
	coordMap.Clear()
	travel1(m, guardCoord, coordinate.GetOffsetForDir(direction.Up))

	result := 0
	for _, newObstacleCoord := range coordMap.GetAllCoordValues() {
		travelMap.Clear()
		if isInLoop(m, guardCoord, coordinate.GetOffsetForDir(direction.Up), newObstacleCoord) {
			result++
		}
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// 260 too low
// 443 too low
// 2112 too high
// 2036 not the right answer
// 2035 not the right answer
// 2037 not the right answer
// 2049 not the right answer -> here i considered the potential obstacle
// correct 1951, example 6

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
