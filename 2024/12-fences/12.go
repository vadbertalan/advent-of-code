// https://adventofcode.com/2024/day/12

package main

import (
	"aoc/utils"
	"aoc/utils/coordinate"
	"aoc/utils/direction"
	"aoc/utils/matrix"
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

const aocDay int = 12

type coord = coordinate.Coord
type mat = matrix.Matrix[string]

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func getCellPerimeter(m mat, coord coord) int {
	perimeter := 0

	offests := coordinate.GetOffsetsArray(false)
	for _, offset := range offests {
		newCoord := coord.GetNewCoord(offset.Dir)
		if !m.IsValidCoord(newCoord) || (m.At(newCoord) != m.At(coord)) {
			perimeter++
		}
	}

	return perimeter
}

func travel(m mat, coord coord, seenMap coordinate.CoordMap) (area, perimeter int) {
	seenMap.Add(coord)

	offests := coordinate.GetOffsetsArray(false)
	for _, offset := range offests {
		newCoord := coord.GetNewCoord(offset.Dir)
		if !seenMap.ContainsCoord(newCoord) && m.IsValidCoord(newCoord) && m.At(newCoord) == m.At(coord) {
			newArea, newPerimeter := travel(m, newCoord, seenMap)
			area += newArea
			perimeter += newPerimeter
		}
	}

	currentPerimeter := getCellPerimeter(m, coord)

	return area + 1, perimeter + currentPerimeter
}

func RandString(n int) string {
	b := make([]byte, n)
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func First(lines []string) (strigifiedResult string) {
	result := 0

	m := matrix.ParseStringMatrix(lines)

	seenMap := make(coordinate.CoordMap)

	regionData := make(map[string][2]int)
	for i := 0; i < len(m.Values); i++ {
		for j := 0; j < len(m.Values[i]); j++ {
			coord := coordinate.Coord{Row: i, Col: j}
			if !seenMap.ContainsCoord(coord) {
				a, p := travel(m, coord, seenMap)
				regionId := RandString(10)
				regionData[regionId] = [2]int{a, p}
			}
		}
	}

	for _, data := range regionData {
		result += data[0] * data[1]
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// 79014 incorrect
// 1573474 correct

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

var fencedMap = coordinate.TravelMap{}

func travelFenceInDir(m mat, fenceDir direction.Direction, coord coord, travelDir direction.Direction) {
	fencedMap.Add(coord, fenceDir)

	newCoord := coord.GetNewCoord(travelDir)

	if m.IsValidCoord(newCoord) && m.At(newCoord) == m.At(coord) {
		diagonalCoord := newCoord.GetNewCoord(fenceDir)
		if m.IsValidCoord(diagonalCoord) && m.At(diagonalCoord) == m.At(coord) {
			return
		}

		travelFenceInDir(m, fenceDir, newCoord, travelDir)
	}
}

func getCellPerimeter2(m mat, coord coord) int {
	perimeter := 0

	offests := coordinate.GetOffsetsArray(false)
	for _, offset := range offests {
		newCoord := coord.GetNewCoord(offset.Dir)
		if !m.IsValidCoord(newCoord) || (m.At(newCoord) != m.At(coord)) {
			if fencedMap.ContainsCoordAndDir(coord, offset.Dir) {
				continue
			}
			perimeter++
			dir1 := coordinate.GetCounterClockwise90DegreeNeighborOffset(offset)
			dir2 := coordinate.GetClockwise90DegreeNeighborOffset(offset)
			travelFenceInDir(m, offset.Dir, coord, dir1.Dir)
			travelFenceInDir(m, offset.Dir, coord, dir2.Dir)
		}
	}

	return perimeter
}

func travel2(m mat, coord coord, seenMap coordinate.CoordMap) (area, perimeter int) {
	seenMap.Add(coord)

	offsets := coordinate.GetOffsetsArray(false)
	for _, offset := range offsets {
		newCoord := coord.GetNewCoord(offset.Dir)
		if !seenMap.ContainsCoord(newCoord) && m.IsValidCoord(newCoord) && m.At(newCoord) == m.At(coord) {
			newArea, newPerimeter := travel2(m, newCoord, seenMap)
			area += newArea
			perimeter += newPerimeter
		}
	}

	currentPerimeter := getCellPerimeter2(m, coord)

	return area + 1, perimeter + currentPerimeter
}

func Second(lines []string) (strigifiedResult string) {
	result := 0

	// Clear map for tests
	fencedMap.Clear()

	m := matrix.ParseStringMatrix(lines)

	seenMap := make(coordinate.CoordMap)

	regionData := make(map[string][2]int)
	for i := 0; i < len(m.Values); i++ {
		for j := 0; j < len(m.Values[i]); j++ {
			coord := coordinate.Coord{Row: i, Col: j}
			if !seenMap.ContainsCoord(coord) {
				a, p := travel2(m, coord, seenMap)
				regionId := RandString(10)
				regionData[regionId] = [2]int{a, p}
			}
		}
	}

	for _, data := range regionData {
		result += data[0] * data[1]
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// correct: 966476

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(5)

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
