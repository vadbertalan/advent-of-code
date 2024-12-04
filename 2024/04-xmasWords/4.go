// https://adventofcode.com/2024/day/4

package main

import (
	"aoc/utils"
	"aoc/utils/coordinate"
	"aoc/utils/matrix"
	"fmt"
	"time"
)

const aocDay int = 4

type coord = coordinate.Coord
type mat = matrix.Matrix[string]

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func getNextLetter(letter string) string {
	switch letter {
	case "X":
		return "M"
	case "M":
		return "A"
	case "A":
		return "S"
	default:
		return "❌" // This is easter egg for Boti ✅❌
	}
}

func isXMASText(m mat, c coord, dirOffset coordinate.DirOffset, nextLetter string) bool {
	newCoord := coord{Row: c.Row + dirOffset.RowOffset, Col: c.Col + dirOffset.ColOffset}
	if !m.IsValidCoord(newCoord) || m.Values[newCoord.Row][newCoord.Col] != nextLetter {
		return false
	}
	if nextLetter == "S" {
		return true
	}
	return isXMASText(m, newCoord, dirOffset, getNextLetter(nextLetter))
}

func First(lines []string) (strigifiedResult string) {
	fmt.Println("--- First ---")

	result := 0

	// Declare and parse matrix
	m := matrix.ParseStringMatrix(lines)

	for i := 0; i < m.RowCount; i++ {
		for j := 0; j < m.ColumnCount; j++ {
			if m.Values[i][j] == "X" {
				for _, dirOffset := range coordinate.GetOffsetsArray(true) {
					if isXMASText(m, coord{Row: i, Col: j}, dirOffset, "M") {
						// fmt.Println("Found XMAS text at", i, j, "with direction", dirOffset)
						result++
					}
				}
			}
		}
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// correct 2551, example 18

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func canFinishFromA(m mat, aCoord coord, dirOffset coordinate.DirOffset) bool {
	newCoord := aCoord.GetNewCoord(dirOffset.Dir)
	if !m.IsValidCoord(newCoord) || m.Values[newCoord.Row][newCoord.Col] != "S" {
		return false
	}

	perpendicularOffsets := coordinate.GetPerpendicularOffsets(dirOffset)

	newCoord1 := aCoord.GetNewCoord(perpendicularOffsets[0].Dir)
	newCoord2 := aCoord.GetNewCoord(perpendicularOffsets[1].Dir)

	if !m.IsValidCoord(newCoord1) || !m.IsValidCoord(newCoord2) {
		return false
	}

	letter1 := m.Values[newCoord1.Row][newCoord1.Col]
	letter2 := m.Values[newCoord2.Row][newCoord2.Col]
	if letter1 == "M" && letter2 == "S" || letter1 == "S" && letter2 == "M" {
		return true
	}
	return false
}

func isMASText(m mat, c coord, dirOffset coordinate.DirOffset, nextLetter string) bool {
	newCoord := c.GetNewCoord(dirOffset.Dir)
	if !m.IsValidCoord(newCoord) || m.Values[newCoord.Row][newCoord.Col] != nextLetter {
		return false
	}
	if nextLetter == "A" {
		return canFinishFromA(m, newCoord, dirOffset)
	}
	return false
}

func Second(lines []string) (strigifiedResult string) {
	fmt.Println("\n--- Second ---")

	result := 0
	foundACoords := []coord{}

	// Declare and parse matrix
	m := matrix.ParseStringMatrix(lines)

	for i := 0; i < m.RowCount; i++ {
		for j := 0; j < m.ColumnCount; j++ {
			if m.Values[i][j] == "M" {
				for _, dirOffset := range coordinate.GetOnlyDiagonalOffsets() {
					c := coord{Row: i, Col: j}
					if isMASText(m, c, dirOffset, "A") {
						aCoord := c.GetNewCoord(dirOffset.Dir)

						if utils.Contains(foundACoords, aCoord) {
							continue
						}

						foundACoords = append(foundACoords, aCoord)
						// fmt.Println("Found MAS text at", i, j, "with direction", dirOffset)
						result++
					}
				}
			}
		}
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// correct 1985, example 9

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(4)

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	result := First(lines)
	fmt.Println(result)

	result = Second(lines)
	fmt.Println(result)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
