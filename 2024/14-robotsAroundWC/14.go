// https://adventofcode.com/2024/day/14

package main

import (
	"aoc/utils"
	"aoc/utils/coordinate"
	"fmt"
	"time"
)

const aocDay int = 14

type coord = coordinate.Coord

// # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #
// # . . . . . . . . . . . . . . . . . . . . . . . . . . . . . #
// # . . . . . . . . . . . . . . . . . . . . . . . . . . . . . #
// # . . . . . . . . . . . . . . . . . . . . . . . . . . . . . #
// # . . . . . . . . . . . . . . . . . . . . . . . . . . . . . #
// # . . . . . . . . . . . . . . # . . . . . . . . . . . . . . #
// # . . . . . . . . . . . . . # # # . . . . . . . . . . . . . #
// # . . . . . . . . . . . . # # # # # . . . . . . . . . . . . #
// # . . . . . . . . . . . # # # # # # # . . . . . . . . . . . #
// # . . . . . . . . . . # # # # # # # # # . . . . . . . . . . #
// # . . . . . . . . . . . . # # # # # . . . . . . . . . . . . #
// # . . . . . . . . . . . # # # # # # # . . . . . . . . . . . #
// # . . . . . . . . . . # # # # # # # # # . . . . . . . . . . #
// # . . . . . . . . . # # # # # # # # # # # . . . . . . . . . #
// # . . . . . . . . # # # # # # # # # # # # # . . . . . . . . #
// # . . . . . . . . . . # # # # # # # # # . . . . . . . . . . #
// # . . . . . . . . . # # # # # # # # # # # . . . . . . . . . #
// # . . . . . . . . # # # # # # # # # # # # # . . . . . . . . #
// # . . . . . . . # # # # # # # # # # # # # # # . . . . . . . #
// # . . . . . . # # # # # # # # # # # # # # # # # . . . . . . #
// # . . . . . . . . # # # # # # # # # # # # # . . . . . . . . #
// # . . . . . . . # # # # # # # # # # # # # # # . . . . . . . #
// # . . . . . . # # # # # # # # # # # # # # # # # . . . . . . #
// # . . . . . # # # # # # # # # # # # # # # # # # # . . . . . #
// # . . . . # # # # # # # # # # # # # # # # # # # # # . . . . #
// # . . . . . . . . . . . . . # # # . . . . . . . . . . . . . #
// # . . . . . . . . . . . . . # # # . . . . . . . . . . . . . #
// # . . . . . . . . . . . . . # # # . . . . . . . . . . . . . #
// # . . . . . . . . . . . . . . . . . . . . . . . . . . . . . #
// # . . . . . . . . . . . . . . . . . . . . . . . . . . . . . #
// # . . . . . . . . . . . . . . . . . . . . . . . . . . . . . #
// # . . . . . . . . . . . . . . . . . . . . . . . . . . . . . #
// # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

type RobotData struct {
	coord          coordinate.Coord
	velRow, velCol int
}

func First(lines []string, rowCount, colCount int) (strigifiedResult string) {
	result := 0
	tl, tr, bl, br := 0, 0, 0, 0

	// line example: p=9,5 v=-3,-3
	for _, line := range lines {
		s1, s2 := utils.SplitIn2(line, " ")
		_, positionStr := utils.SplitIn2(s1, "=")
		_, velocityString := utils.SplitIn2(s2, "=")
		c := coordinate.ParseCoordStr(positionStr, ",")
		velocityCoord := coordinate.ParseCoordStr(velocityString, ",")

		// inverse the row and col
		c.Row, c.Col = c.Col, c.Row

		// inverse the vel x y
		robot := RobotData{coord: c, velRow: velocityCoord.Col, velCol: velocityCoord.Row}

		for i := 0; i < 100; i++ {
			newCoord := coord{Row: robot.coord.Row + robot.velRow, Col: robot.coord.Col + robot.velCol}
			if newCoord.Row < 0 {
				newCoord.Row = rowCount + newCoord.Row
			}
			if newCoord.Col < 0 {
				newCoord.Col = colCount + newCoord.Col
			}
			if newCoord.Row >= rowCount {
				newCoord.Row = newCoord.Row % rowCount
			}
			if newCoord.Col >= colCount {
				newCoord.Col = newCoord.Col % colCount
			}

			robot.coord = newCoord
		}

		if robot.coord.Row < rowCount/2 && robot.coord.Col < colCount/2 {
			tl++
		} else if robot.coord.Row < rowCount/2 && robot.coord.Col > colCount/2 {
			tr++
		} else if robot.coord.Row > rowCount/2 && robot.coord.Col < colCount/2 {
			bl++
		} else if robot.coord.Row > rowCount/2 && robot.coord.Col > colCount/2 {
			br++
		}
	}

	fmt.Println(tl, tr, bl, br)
	result = tl * tr * bl * br
	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// 47658600 too low, forgot to modify input size
// correct 228410028, example 12

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func Second(lines []string, rowCount, colCount int) (strigifiedResult string) {
	robots := make([]*RobotData, 0)

	// line example: p=9,5 v=-3,-3
	for _, line := range lines {
		s1, s2 := utils.SplitIn2(line, " ")
		_, positionStr := utils.SplitIn2(s1, "=")
		_, velocityString := utils.SplitIn2(s2, "=")
		c := coordinate.ParseCoordStr(positionStr, ",")
		velocityCoord := coordinate.ParseCoordStr(velocityString, ",")

		// inverse the row and col
		c.Row, c.Col = c.Col, c.Row

		// inverse the vel x y
		robot := RobotData{coord: c, velRow: velocityCoord.Col, velCol: velocityCoord.Row}
		robots = append(robots, &robot)
	}

	secondsElapsed := 0
	for {
		coordMap := coordinate.CoordMap{}

		for _, robot := range robots {
			newCoord := coord{Row: robot.coord.Row + robot.velRow, Col: robot.coord.Col + robot.velCol}
			if newCoord.Row < 0 {
				newCoord.Row = rowCount + newCoord.Row
			}
			if newCoord.Col < 0 {
				newCoord.Col = colCount + newCoord.Col
			}
			if newCoord.Row >= rowCount {
				newCoord.Row = newCoord.Row % rowCount
			}
			if newCoord.Col >= colCount {
				newCoord.Col = newCoord.Col % colCount
			}

			coordMap.Add(newCoord)

			robot.coord = newCoord
		}

		secondsElapsed++

		// horizontal: 18 121 224 ... 3005 3108 => f(x) = 18 + 103x
		// vertical: 77 178 279 ... 3006 3107 => f(x) = 77 + 101x

		if (secondsElapsed-18)%103 == 0 {
			for i := 0; i < rowCount; i++ {
				for j := 0; j < colCount; j++ {
					if coordMap.ContainsRowCol(i, j) {
						fmt.Printf("# ")
					} else {
						fmt.Printf(". ")
					}
				}
				fmt.Println()
			}

			fmt.Println("Seconds elapsed:", secondsElapsed)

			var input string
			fmt.Scanln(&input)
			if input == "s" {
				break
			}
			fmt.Println()
			fmt.Println()
		}
	}

	strigifiedResult = fmt.Sprint(secondsElapsed)
	return strigifiedResult
}

// 18 incorrect
// 77 incorrect
// correct 8258

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(1)

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	fmt.Println("--- First ---")
	// result := First(lines, 7, 11)
	result := First(lines, 103, 101)
	fmt.Println(result)

	fmt.Println("\n--- Second ---")
	result = Second(lines, 103, 101)
	fmt.Println(result)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
