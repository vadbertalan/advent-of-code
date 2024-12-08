// https://adventofcode.com/2024/day/8

package main

import (
	"aoc/utils"
	"aoc/utils/coordinate"
	"aoc/utils/matrix"
	"fmt"
	"math"
	"time"
)

const aocDay int = 8

type coord = coordinate.Coord

func getXforYOnLine(y, x1, y1, x2, y2 int) int {
	return x1 + (y-y1)*(x2-x1)/(y2-y1)
}

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func First(lines []string) (strigifiedResult string) {
	result := 0

	m := matrix.ParseStringMatrix(lines)

	clusters := make(map[string][]coord)

	for i, row := range m.Values {
		for j, cell := range row {
			if cell != "." {
				clusters[cell] = append(clusters[cell], coord{Row: i, Col: j})
			}
		}
	}

	coordMap := make(coordinate.CoordMap)

	for _, coords := range clusters {
		for i := 0; i < len(coords)-1; i++ {
			for j := i + 1; j < len(coords); j++ {
				colDiff := int(math.Abs(float64(coords[i].Col - coords[j].Col)))

				var leftCoord, rightCoord coord
				if coords[i].Col < coords[j].Col {
					leftCoord = coords[i]
					rightCoord = coords[j]
				} else {
					leftCoord = coords[j]
					rightCoord = coords[i]
				}

				newCoord1Y := leftCoord.Col - colDiff
				newCoord1X := getXforYOnLine(newCoord1Y, leftCoord.Row, leftCoord.Col, rightCoord.Row, rightCoord.Col)
				newCoord1 := coord{Row: newCoord1X, Col: newCoord1Y}

				newCoord2Y := rightCoord.Col + colDiff
				newCoord2X := getXforYOnLine(newCoord2Y, leftCoord.Row, leftCoord.Col, rightCoord.Row, rightCoord.Col)
				newCoord2 := coord{Row: newCoord2X, Col: newCoord2Y}

				if m.IsValidCoord(newCoord1) {
					coordMap.Add(newCoord1)
				}
				if m.IsValidCoord(newCoord2) {
					coordMap.Add(newCoord2)
				}
			}
		}
	}

	result = len(coordMap)
	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// correct 344, example 14

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func getAntinodeCoordLeft(leftCoord, rightCoord coord) coord {
	colDiff := int(math.Abs(float64(leftCoord.Col - rightCoord.Col)))

	newCoordY := leftCoord.Col - colDiff
	newCoordX := getXforYOnLine(newCoordY, leftCoord.Row, leftCoord.Col, rightCoord.Row, rightCoord.Col)
	antinodeCoord := coord{Row: newCoordX, Col: newCoordY}

	return antinodeCoord
}

func getAntinodeCoordRight(leftCoord, rightCoord coord) coord {
	colDiff := int(math.Abs(float64(leftCoord.Col - rightCoord.Col)))

	newCoordY := rightCoord.Col + colDiff
	newCoordX := getXforYOnLine(newCoordY, leftCoord.Row, leftCoord.Col, rightCoord.Row, rightCoord.Col)
	antinodeCoord := coord{Row: newCoordX, Col: newCoordY}

	return antinodeCoord
}

func Second(lines []string) (strigifiedResult string) {
	result := 0

	m := matrix.ParseStringMatrix(lines)

	clusters := make(map[string][]coord)

	for i, row := range m.Values {
		for j, cell := range row {
			if cell != "." {
				clusters[cell] = append(clusters[cell], coord{Row: i, Col: j})
			}
		}
	}

	coordMap := make(coordinate.CoordMap)

	for _, coords := range clusters {
		for i := 0; i < len(coords)-1; i++ {
			for j := i + 1; j < len(coords); j++ {
				var originalLeftCoord, originalRightCoord coord
				if coords[i].Col < coords[j].Col {
					originalLeftCoord = coords[i]
					originalRightCoord = coords[j]
				} else {
					originalLeftCoord = coords[j]
					originalRightCoord = coords[i]
				}

				coordMap.Add(originalLeftCoord)
				coordMap.Add(originalRightCoord)

				// while we have valid anti-node coords in left dir, get all anti-node coords
				leftCoord := originalLeftCoord
				rightCoord := originalRightCoord
				leftAntinodeCoord := getAntinodeCoordLeft(leftCoord, rightCoord)
				for m.IsValidCoord(leftAntinodeCoord) {
					coordMap.Add(leftAntinodeCoord)

					rightCoord = leftCoord
					leftCoord = leftAntinodeCoord
					leftAntinodeCoord = getAntinodeCoordLeft(leftCoord, rightCoord)
				}

				// while we have valid anti-node coords in right dir, get all anti-node coords
				leftCoord = originalLeftCoord
				rightCoord = originalRightCoord
				rightAntinodeCoord := getAntinodeCoordRight(leftCoord, rightCoord)
				for m.IsValidCoord(rightAntinodeCoord) {
					coordMap.Add(rightAntinodeCoord)

					leftCoord = rightCoord
					rightCoord = rightAntinodeCoord
					rightAntinodeCoord = getAntinodeCoordRight(leftCoord, rightCoord)
				}
			}
		}
	}

	result = len(coordMap)
	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// correct 1182, example 34

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
