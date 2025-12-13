// https://adventofcode.com/2025/day/4

package main

import (
	"aoc/utils-go"
	"aoc/utils-go/coordinate"
	"aoc/utils-go/matrix"
	"fmt"
	"time"
)

const aocDay int = 4

const ROLL string = "@"
const EMPTY string = "."

type coord = coordinate.Coord
type mat = matrix.Matrix[string]

func hasLessThan4AdjacentRolls(c coord, m mat) bool {
	count := 0

	offsets := coordinate.GetOffsetsArray(true)

	for _, offset := range offsets {
		newCoord := c.GetNewCoord(offset.Dir)
		if m.IsValidCoord(newCoord) && m.At(newCoord) == ROLL {
			count++
		}
	}

	return count < 4
}

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func First(lines []string) (strigifiedResult string) {
	result := 0

	m := matrix.ParseStringMatrix(lines)

	for i := 0; i < len(m.Values); i++ {
		for j := 0; j < len(m.Values[i]); j++ {
			coord := coordinate.Coord{Row: i, Col: j}
			if m.At(coord) == ROLL && hasLessThan4AdjacentRolls(coord, m) {
				result++
			}
		}
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// Your puzzle answer was 1569.

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func Second(lines []string) (strigifiedResult string) {
	result := 0

	m := matrix.ParseStringMatrix(lines)

	for {
		removedCount := 0
		for i := 0; i < len(m.Values); i++ {
			for j := 0; j < len(m.Values[i]); j++ {
				coord := coordinate.Coord{Row: i, Col: j}
				if m.At(coord) == ROLL && hasLessThan4AdjacentRolls(coord, m) {
					removedCount++
					m.Set(coord, EMPTY)
				}
			}
		}
		if removedCount == 0 {
			break
		} else {
			result += removedCount
		}
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// Your puzzle answer was 9280.

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(1)

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	fmt.Println("--- First ---")
	result := First(lines)
	fmt.Println(result)

	fmt.Println("\n--- Second ---")
	result = Second(lines)
	fmt.Println(result)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
