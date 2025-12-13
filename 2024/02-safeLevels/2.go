// https://adventofcode.com/2024/day/2

package main

import (
	"aoc/utils-go"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"time"
)

const aocDay int = 2

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func isSafe1(row []int) bool {
	previous := -1
	isIncreasing := true

	for j, cell := range row {
		if j == 0 {
			previous = cell
		} else {
			if j == 1 {
				if cell < previous {
					isIncreasing = false
				}
			} else {
				if isIncreasing {
					if cell <= previous {
						break
					}
				} else {
					if cell >= previous {
						break
					}
				}
			}

			if math.Abs(float64(cell-previous)) != 1 && math.Abs(float64(cell-previous)) != 2 && math.Abs(float64(cell-previous)) != 3 {
				break
			}
		}

		if j == len(row)-1 {
			return true
		}

		previous = cell
	}

	return false
}

func First(lines []string) string {
	fmt.Println("--- First ---")

	result := 0

	for _, line := range lines {
		intChars := strings.Split(line, " ")
		row := make([]int, len(intChars))
		for col, c := range intChars {
			x, _ := strconv.Atoi(c)
			row[col] = x
		}

		if isSafe1(row) {
			result++
		}
	}

	return fmt.Sprint(result)
}

// correct 472, example 2

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func isTooFar(cell1, cell2 int) bool {
	diff := math.Abs(float64(cell1 - cell2))
	return diff != 1 && diff != 2 && diff != 3
}

func isSafe3(row []int) (bool, bool, int) {
	cellZero := row[0]

	if len(row) == 1 {
		return true, true, cellZero
	}

	restIsIncreasing, restIsSafe, secondCell := isSafe3(row[1:])

	if len(row) == 2 {
		return cellZero < secondCell, !isTooFar(cellZero, secondCell), cellZero
	}

	isInLine := restIsIncreasing && cellZero < secondCell || !restIsIncreasing && cellZero > secondCell
	isSafe := restIsSafe && isInLine && !isTooFar(cellZero, secondCell)

	return restIsIncreasing, isSafe, cellZero
}

func Second(lines []string) string {
	fmt.Println("--- Second ---")

	result := 0

	for _, line := range lines {
		intChars := strings.Split(line, " ")
		row := make([]int, len(intChars))
		for col, c := range intChars {
			x, _ := strconv.Atoi(c)
			row[col] = x
		}

		for i := 0; i < len(row); i++ {
			newRow := slices.Concat(row[:i], row[i+1:])
			if _, isSafe, _ := isSafe3(newRow); isSafe {
				result++
				break
			}
		}
	}

	return fmt.Sprint(result)
}

// correct 520, example 4

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
