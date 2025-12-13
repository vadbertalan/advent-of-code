// https://adventofcode.com/2024/day/18

package main

import (
	"aoc/utils-go"
	"aoc/utils-go/coordinate"
	"fmt"
	"strings"
	"time"

	"github.com/rameshputalapattu/heapq"
)

const aocDay int = 18

type coord = coordinate.Coord

//	____            _     _
//
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

type node struct {
	coord coord
	dist  int
	path  []coord
}

func less(a, b node) bool {
	return a.dist < b.dist
}

func isValidCoord(c coord, rowCount, colCount int) bool {
	return c.Row >= 0 && c.Row <= rowCount && c.Col >= 0 && c.Col <= colCount
}

func dijkstraFromStartToEndCoord(startCoord coord, endCoord coord, obstacles coordinate.CoordMap, rowCount, colCount int) int {
	pq := heapq.NewPQWithItems([]node{}, less)
	pq.Push(node{coord: startCoord, dist: 0, path: []coord{startCoord}})

	seen := coordinate.CoordMap{}

	for pq.Len() > 0 {
		item := pq.Pop()

		if item.coord == endCoord {
			return item.dist
		}

		if seen.ContainsCoord(item.coord) {
			continue
		}

		seen.Add(item.coord)

		offsets := coordinate.GetOffsetsArray(false)
		for _, offset := range offsets {
			newCoord := item.coord.GetNewCoord(offset.Dir)

			if !isValidCoord(newCoord, rowCount, colCount) || obstacles.ContainsCoord(newCoord) {
				continue
			}

			pathToNewCoord := append(item.path, newCoord)
			pq.Push(node{coord: newCoord, dist: item.dist + 1, path: utils.CloneArray(pathToNewCoord)})
		}
	}

	return -1
}

func First(lines []string, rowCount, colCount, byteCount int) (strigifiedResult string) {
	bytesCoordMap := coordinate.CoordMap{}

	for _, line := range lines[:byteCount] {
		xyStrs := strings.Split(line, ",")
		xys := utils.ConvertToInts(xyStrs)

		// ??????? I think the input is wrong => X and Y are swapped
		// bytesCoordMap.Add(coord{Row: xys[0], Col: xys[1]})
		bytesCoordMap.Add(coord{Row: xys[1], Col: xys[0]})
	}

	result := dijkstraFromStartToEndCoord(coord{Row: 0, Col: 0}, coord{Row: rowCount, Col: colCount}, bytesCoordMap, rowCount, colCount)

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// correct: 334, example: 22

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func Second(lines []string, rowCount, colCount, startByteCount int) (strigifiedResult string) {
	bytesCoordMap := coordinate.CoordMap{}

	for _, line := range lines[:startByteCount] {
		xyStrs := strings.Split(line, ",")
		xys := utils.ConvertToInts(xyStrs)
		bytesCoordMap.Add(coord{Row: xys[1], Col: xys[0]})
	}

	for _, line := range lines[startByteCount:] {
		xyStrs := strings.Split(line, ",")
		xys := utils.ConvertToInts(xyStrs)
		bytesCoordMap.Add(coord{Row: xys[1], Col: xys[0]})

		pathLength := dijkstraFromStartToEndCoord(coord{Row: 0, Col: 0}, coord{Row: rowCount, Col: colCount}, bytesCoordMap, rowCount, colCount)

		if pathLength == -1 {
			strigifiedResult = fmt.Sprintf("%d,%d", xys[0], xys[1])
			return strigifiedResult
		}
	}

	return "No result found"
}

// 4,3 incorrect
// 20,12 correct, example 6,1

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(1)

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	fmt.Println("--- First ---")
	result := First(lines, 6, 6, 12)
	// result := First(lines, 70, 70, 1024)
	fmt.Println(result)

	fmt.Println("\n--- Second ---")
	// result = Second(lines, 6, 6, 12)
	result = Second(lines, 70, 70, 1024)
	fmt.Println(result)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
