// https://adventofcode.com/2024/day/10

package main

import (
	"aoc/utils"
	"aoc/utils/coordinate"
	"aoc/utils/matrix"
	"fmt"
	"time"
)

const aocDay int = 10

type coord = coordinate.Coord
type mat = matrix.Matrix[int]

func test(currentCoord, nextCoord coordinate.Coord, currentValue, nextValue int) bool {
	return nextValue-currentValue == 1
}

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func First(lines []string) (strigifiedResult string) {
	result := 0

	m := matrix.ParseDigitMatrix(lines)

	trailHeads := m.GetAllCoordsWhich(func(value int) bool {
		return value == 0
	})

	nineHeights := m.GetAllCoordsWhich(func(value int) bool {
		return value == 9
	})

	for _, trailHead := range trailHeads {
		for _, nineHeight := range nineHeights {
			if m.IsPathBetween(trailHead, nineHeight, test) {
				result++
			}
		}
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// 182 incorrect - pls read the final sentence of the question
// correct 607, example 36, e2 4

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func Second(lines []string) (strigifiedResult string) {
	result := 0

	m := matrix.ParseDigitMatrix(lines)

	trailHeads := m.GetAllCoordsWhich(func(value int) bool {
		return value == 0
	})

	nineHeights := m.GetAllCoordsWhich(func(value int) bool {
		return value == 9
	})

	for _, trailHead := range trailHeads {
		for _, nineHeight := range nineHeights {
			score := m.CountPathsBetween(trailHead, nineHeight, test)
			result += score
		}
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// correct 1384, example 81, e2 13

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(2)

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
