package main

import (
	"aoc/utils"
	"aoc/utils/coordinate"
	"aoc/utils/matrix"
	"fmt"
	"time"
)

const aocDay int = 13

type coord = coordinate.Coord
type mat = matrix.Matrix[string]

const rock = "#"
const ash = "."

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func areRowsEqual(rowIndex1, rowIndex2 int, m mat) bool {
	for col := 0; col < m.ColumnCount; col++ {
		if m.Values[rowIndex1][col] != m.Values[rowIndex2][col] {
			return false
		}
	}
	return true
}

func areColsEqual(colIndex1, colIndex2 int, m mat) bool {
	for row := 0; row < m.RowCount; row++ {
		if m.Values[row][colIndex1] != m.Values[row][colIndex2] {
			return false
		}
	}
	return true
}

func first(lines []string) {
	fmt.Println("--- First ---")

	result := 0

	matrixLines := []string{}
	for _, line := range lines {
		if line != "" {
			matrixLines = append(matrixLines, line)
		} else {
			m := matrix.ParseStringMatrix(matrixLines)

			// Check vertical mirrors
			for mirrorAfterIndex := 0; mirrorAfterIndex < m.ColumnCount-1; mirrorAfterIndex++ {
				isMirror := true
				dist := 1
				for i := mirrorAfterIndex; i >= 0 && i+dist < m.ColumnCount; i-- {
					if !areColsEqual(i, i+dist, m) {
						isMirror = false
						break
					}
					dist += 2
				}

				if isMirror {
					result += mirrorAfterIndex + 1
					break
				}
			}

			// Check horizontal mirrors
			for mirrorAfterIndex := 0; mirrorAfterIndex < m.RowCount-1; mirrorAfterIndex++ {
				isMirror := true
				dist := 1
				for i := mirrorAfterIndex; i >= 0 && i+dist < m.RowCount; i-- {
					if !areRowsEqual(i, i+dist, m) {
						isMirror = false
						break
					}
					dist += 2
				}

				if isMirror {
					result += 100 * (mirrorAfterIndex + 1)
					break
				}
			}

			matrixLines = []string{}
		}

	}

	fmt.Println(result)
}

// tried 2005, but wrong
// Your puzzle answer was 37025.

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func areRowsEqual2(rowIndex1, rowIndex2 int, m mat) (bool, bool) {
	diffFound := false
	for col := 0; col < m.ColumnCount; col++ {
		if m.Values[rowIndex1][col] != m.Values[rowIndex2][col] {
			if diffFound {
				return false, false
			}
			diffFound = true
		}
	}
	return true, diffFound
}

func areColsEqual2(colIndex1, colIndex2 int, m mat) (bool, bool) {
	diffFound := false
	for row := 0; row < m.RowCount; row++ {
		if m.Values[row][colIndex1] != m.Values[row][colIndex2] {
			if diffFound {
				return false, false
			}
			diffFound = true
		}
	}
	return true, diffFound
}

func second(lines []string) {
	fmt.Println("\n--- Second ---")

	result := 0

	matrixLines := []string{}
	for _, line := range lines {
		if line != "" {
			matrixLines = append(matrixLines, line)
		} else {
			m := matrix.ParseStringMatrix(matrixLines)

			standardVerticalScore := 0
			standardHorizontalScore := 0
			newVerticalScore := 0
			newHorizontalScore := 0

			// Check vertical mirrors
			isNewVerticalReflectionLine := false

			for mirrorAfterIndex := 0; mirrorAfterIndex < m.ColumnCount-1; mirrorAfterIndex++ {
				diffCount := 0
				isMirror := true
				dist := 1
				for i := mirrorAfterIndex; i >= 0 && i+dist < m.ColumnCount; i-- {
					areEqual, diffFound := areColsEqual2(i, i+dist, m)
					if !areEqual {
						isMirror = false
						break
					}
					if diffFound {
						if diffCount > 0 {
							isMirror = false
							break
						} else {
							diffCount++
						}
					}
					dist += 2
				}

				if isMirror {
					if diffCount > 0 {
						newVerticalScore += mirrorAfterIndex + 1
						isNewVerticalReflectionLine = true
					} else {
						standardVerticalScore += mirrorAfterIndex + 1
					}
				}
			}

			// Check horizontal mirrors
			isNewHorizontalReflectionLine := false
			for mirrorAfterIndex := 0; mirrorAfterIndex < m.RowCount-1; mirrorAfterIndex++ {
				diffCount := 0
				isMirror := true
				dist := 1
				for i := mirrorAfterIndex; i >= 0 && i+dist < m.RowCount; i-- {
					areEqual, diffFound := areRowsEqual2(i, i+dist, m)
					if !areEqual {
						isMirror = false
						break
					}
					if diffFound {
						if diffCount > 0 {
							isMirror = false
							break
						} else {
							diffCount++
						}
					}
					dist += 2
				}

				if isMirror {
					if diffCount > 0 {
						newHorizontalScore += 100 * (mirrorAfterIndex + 1)
						isNewHorizontalReflectionLine = true
					} else {
						standardHorizontalScore += mirrorAfterIndex + 1
					}
				}
			}

			if isNewVerticalReflectionLine && !isNewHorizontalReflectionLine {
				result += newVerticalScore
			}
			if isNewHorizontalReflectionLine && !isNewVerticalReflectionLine {
				result += newHorizontalScore
			}
			if !isNewVerticalReflectionLine && !isNewHorizontalReflectionLine {
				result += newVerticalScore + newHorizontalScore
			}

			matrixLines = []string{}
		}

	}

	fmt.Println(result)
}

// too high 39444
// too high 45715
// too low 	28400
// wrong 	28669
// wrong 	23615
// Your puzzle answer was 32854.

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(1)

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	// Mark end of last matrix
	lines = append(lines, "")

	first(lines)

	second(lines)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
