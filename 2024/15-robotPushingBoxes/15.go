// https://adventofcode.com/2024/day/15

package main

import (
	"aoc/utils"
	"aoc/utils/coordinate"
	"aoc/utils/direction"
	"aoc/utils/matrix"
	"fmt"
	"time"
)

const aocDay int = 15

type coord = coordinate.Coord
type mat = matrix.Matrix[string]
type dir = direction.Direction

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func moveBox(m mat, boxCoord coord, dir dir) bool {
	newCoord := boxCoord.GetNewCoord(dir)

	if m.At(newCoord) == "#" {
		return false
	}

	if m.At(newCoord) == "." {
		m.Set(newCoord, "O")
		return true
	}

	// it's a wall
	if moveBox(m, newCoord, dir) {
		m.Set(newCoord, "O")
		return true
	}

	return false
}

func moveRobot(m mat, robotCoord coord, dir dir) (newRobotCoord coord) {
	newCoord := robotCoord.GetNewCoord(dir)

	if m.At(newCoord) == "#" {
		return robotCoord
	}

	if m.At(newCoord) == "." {
		m.Set(newCoord, "@")
		m.Set(robotCoord, ".")

		return newCoord
	}

	// it's a wall
	if moveBox(m, newCoord, dir) {
		m.Set(newCoord, "@")
		m.Set(robotCoord, ".")

		return newCoord
	}

	return robotCoord
}

func First(lines []string) (strigifiedResult string) {
	result := 0

	_, emptyLineIndex := utils.Find(lines, func(line string) bool { return line == "" })

	m, startCoord := matrix.ParseStringMatrixAndGetStartingPoint(lines[:emptyLineIndex], func(value string) bool { return value == "@" })

	robotCoord := *startCoord
	for _, line := range lines[emptyLineIndex+1:] {
		for _, c := range line {
			if c == '<' {
				robotCoord = moveRobot(m, robotCoord, direction.Left)
			} else if c == '>' {
				robotCoord = moveRobot(m, robotCoord, direction.Right)
			} else if c == '^' {
				robotCoord = moveRobot(m, robotCoord, direction.Up)
			} else if c == 'v' {
				robotCoord = moveRobot(m, robotCoord, direction.Down)
			}
		}
	}

	// Count result
	m.ForEach(func(i, j int, value string) {
		if value == "O" {
			result += j + i*100
		}
	})

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// correct: 1446158, e1: 2028, e2: 10092

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func moveBox2(m mat, boxLeftCoord, boxRightCoord coord, dir dir) bool {
	boxLeftNewCoord := boxLeftCoord.GetNewCoord(dir)
	boxRightNewCoord := boxRightCoord.GetNewCoord(dir)

	switch dir {
	case direction.Left:
		{
			if m.At(boxLeftNewCoord) == "#" {
				return false
			}
			if m.At(boxLeftNewCoord) == "." {
				m.Set(boxLeftNewCoord, "[")
				m.Set(boxRightNewCoord, "]")
				m.Set(boxRightCoord, ".")
				return true
			}
			// it's a ] box
			if moveBoxR(m, boxLeftNewCoord, dir) {
				m.Set(boxLeftNewCoord, "[")
				m.Set(boxRightNewCoord, "]")
				m.Set(boxRightCoord, ".")
				return true
			}
		}
	case direction.Right:
		{
			if m.At(boxRightNewCoord) == "#" {
				return false
			}
			if m.At(boxRightNewCoord) == "." {
				m.Set(boxLeftNewCoord, "[")
				m.Set(boxRightNewCoord, "]")
				m.Set(boxLeftCoord, ".")
				return true
			}
			// it's a [ box
			if moveBoxL(m, boxRightNewCoord, dir) {
				m.Set(boxLeftNewCoord, "[")
				m.Set(boxRightNewCoord, "]")
				m.Set(boxLeftCoord, ".")
				return true
			}
		}
	case direction.Up, direction.Down:
		{
			if m.At(boxLeftNewCoord) == "#" || m.At(boxRightNewCoord) == "#" {
				return false
			}
			if m.At(boxLeftNewCoord) == "." && m.At(boxRightNewCoord) == "." {
				m.Set(boxLeftNewCoord, "[")
				m.Set(boxRightNewCoord, "]")
				m.Set(boxLeftCoord, ".")
				m.Set(boxRightCoord, ".")
				return true
			}
			if m.At(boxLeftNewCoord) == "[" && moveBoxL(m, boxLeftNewCoord, dir) {
				m.Set(boxLeftNewCoord, "[")
				m.Set(boxRightNewCoord, "]")
				m.Set(boxLeftCoord, ".")
				m.Set(boxRightCoord, ".")
				return true
			}
			if m.At(boxLeftNewCoord) == "]" && m.At(boxRightNewCoord) == "." && moveBoxR(m, boxLeftNewCoord, dir) {
				m.Set(boxLeftNewCoord, "[")
				m.Set(boxRightNewCoord, "]")
				m.Set(boxLeftCoord, ".")
				m.Set(boxRightCoord, ".")
				return true
			}
			if m.At(boxLeftNewCoord) == "." && m.At(boxRightNewCoord) == "[" && moveBoxL(m, boxRightNewCoord, dir) {
				m.Set(boxLeftNewCoord, "[")
				m.Set(boxRightNewCoord, "]")
				m.Set(boxLeftCoord, ".")
				m.Set(boxRightCoord, ".")
				return true
			}
			if m.At(boxLeftNewCoord) == "]" && moveBoxR(m, boxLeftNewCoord, dir) && m.At(boxRightNewCoord) == "[" && moveBoxL(m, boxRightNewCoord, dir) {
				m.Set(boxLeftNewCoord, "[")
				m.Set(boxRightNewCoord, "]")
				m.Set(boxLeftCoord, ".")
				m.Set(boxRightCoord, ".")
				return true
			}
		}
	}
	return false
}

func moveBoxR(m mat, boxRightCoord coord, dir dir) bool {
	boxLeftCoord := coord{Row: boxRightCoord.Row, Col: boxRightCoord.Col - 1}

	return moveBox2(m, boxLeftCoord, boxRightCoord, dir)
}

func moveBoxL(m mat, boxLeftCoord coord, dir dir) bool {
	boxRightCoord := coord{Row: boxLeftCoord.Row, Col: boxLeftCoord.Col + 1}

	return moveBox2(m, boxLeftCoord, boxRightCoord, dir)
}

func moveRobot2(m mat, robotCoord coord, dir dir) (newMat mat, newRobotCoord coord) {
	newCoord := robotCoord.GetNewCoord(dir)

	if m.At(newCoord) == "#" {
		return m, robotCoord
	}

	if m.At(newCoord) == "." {
		m.Set(newCoord, "@")
		m.Set(robotCoord, ".")

		return m, newCoord
	}

	// it's a wall

	// need a copy because deep recursive instance will do the move, while it is possible that those moves should not be possible
	// so I revert those changes like this
	mCopy := m.Clone()
	if mCopy.At(newCoord) == "[" && moveBoxL(mCopy, newCoord, dir) || mCopy.At(newCoord) == "]" && moveBoxR(mCopy, newCoord, dir) {
		mCopy.Set(newCoord, "@")
		mCopy.Set(robotCoord, ".")

		return mCopy, newCoord
	}

	return m, robotCoord
}

func Second(lines []string) (strigifiedResult string) {
	result := 0

	_, emptyLineIndex := utils.Find(lines, func(line string) bool { return line == "" })

	mStr := lines[:emptyLineIndex]

	m := mat{
		Values:      make([][]string, len(mStr)),
		RowCount:    len(mStr),
		ColumnCount: len(mStr[0]) * 2,
	}

	var startCoord coord

	// Get extended matrix
	for row, line := range mStr {
		m.Values[row] = make([]string, len(line)*2)
		for col, c := range line {
			if c == '#' {
				m.Values[row][2*col] = "#"
				m.Values[row][2*col+1] = "#"
			} else if c == '.' {
				m.Values[row][2*col] = "."
				m.Values[row][2*col+1] = "."
			} else if c == '@' {
				m.Values[row][2*col] = "@"
				m.Values[row][2*col+1] = "."

				startCoord = coord{Row: row, Col: 2 * col}
			} else {
				// it's a wall: "O"
				m.Values[row][2*col] = "["
				m.Values[row][2*col+1] = "]"
			}
		}
	}

	// Perform moves
	robotCoord := startCoord
	for _, line := range lines[emptyLineIndex+1:] {
		for _, c := range line {
			if c == '<' {
				m, robotCoord = moveRobot2(m, robotCoord, direction.Left)
			} else if c == '>' {
				m, robotCoord = moveRobot2(m, robotCoord, direction.Right)
			} else if c == '^' {
				m, robotCoord = moveRobot2(m, robotCoord, direction.Up)
			} else if c == 'v' {
				m, robotCoord = moveRobot2(m, robotCoord, direction.Down)
			}
		}
	}

	// Count result
	m.ForEach(func(i, j int, value string) {
		if value == "[" {
			result += j + i*100
		}
	})

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// 1461173 too high
// 1448229 too high -> tried to see min dist of left and right edges
// 1433400 too low -> tried to see only row distances, just to understand better where the result should be
// correct: 1446175, e3 9021

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(3)

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
