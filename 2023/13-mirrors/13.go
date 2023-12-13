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

func parseMatrix(lines []string) mat {
	m := matrix.Matrix[string]{
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
			m := parseMatrix(matrixLines)

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
	for p, line := range lines {
		if line != "" {
			matrixLines = append(matrixLines, line)
		} else {
			fmt.Println("------")
			fmt.Println(p)
			m := parseMatrix(matrixLines)

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
					fmt.Println("	Vertical mirror after index:", mirrorAfterIndex)
					if diffCount > 0 {
						newVerticalScore += mirrorAfterIndex + 1
						fmt.Println("	Mirror with diffcount", diffCount)
						isNewVerticalReflectionLine = true
					} else {
						standardVerticalScore += mirrorAfterIndex + 1
						fmt.Println("	Mirror standard")
					}
					// break
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
					fmt.Println("Horizontal mirror after index:", mirrorAfterIndex)

					if diffCount > 0 {
						newHorizontalScore += 100 * (mirrorAfterIndex + 1)
						fmt.Println("	Mirror with 1 diffcount", diffCount)
						isNewHorizontalReflectionLine = true
					} else {
						standardHorizontalScore += mirrorAfterIndex + 1
						fmt.Println("	Mirror standard")
					}
					// break
				}
			}

			if isNewVerticalReflectionLine && isNewHorizontalReflectionLine {
				fmt.Println("Oh uh, trouble! -> both new vertical and horizontal refl lines")
			}
			if isNewVerticalReflectionLine && !isNewHorizontalReflectionLine {
				fmt.Println("	Adding vertical only for new ver refl line")
				result += newVerticalScore
			}
			if isNewHorizontalReflectionLine && !isNewVerticalReflectionLine {
				fmt.Println("	Adding horizontal only for new hor refl line")
				result += newHorizontalScore
			}
			if !isNewVerticalReflectionLine && !isNewHorizontalReflectionLine {
				fmt.Println("	Adding standard way (no new refl lines)")
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

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt()

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	// Mark end of last matrix
	lines = append(lines, "")

	first(lines)

	second(lines)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
