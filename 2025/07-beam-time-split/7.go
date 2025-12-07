// https://adventofcode.com/2025/day/7

package main

import (
	"aoc/utils"
	"aoc/utils/collections"
	"aoc/utils/coordinate"
	"aoc/utils/direction"
	"aoc/utils/matrix"
	"fmt"
	"time"
)

const aocDay int = 7

const START = "S"
const SPLITTER = "^"

type coord = coordinate.Coord
type mat = matrix.Matrix[string]

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

var cache map[string]int = make(map[string]int)

func getCacheKey(c coord) string {
	return fmt.Sprintf("%d-%d", c.Row, c.Col)
}

var splittersUsed *collections.Set[coord] = collections.NewSet[coord]()

func goBeam(m mat, start coord) int {
	key := getCacheKey(start)
	_, ok := cache[key]
	if ok {
		return 0
	}

	it := start
	for it.Row < m.RowCount && m.At(it) != SPLITTER {
		it = it.GetNewCoord(direction.Down)
	}

	if it.Row == m.RowCount {
		cache[key] = 0
		return 0
	}

	// if m.At(it) == SPLITTER {

	if splittersUsed.Has(it) {
		cache[key] = 0
		return 0
	}

	splittersUsed.Add(it)
	result := 1 + goBeam(m, it.GetNewCoord(direction.Left)) + goBeam(m, it.GetNewCoord(direction.Right))
	cache[key] = result
	return result
}

func First(lines []string) (strigifiedResult string) {
	m, startCoord := matrix.ParseStringMatrixAndGetStartingPoint(lines, func(value string) bool { return value == START })

	strigifiedResult = fmt.Sprint(goBeam(m, *startCoord))
	return strigifiedResult
}

// Your puzzle answer was 1630.

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

var cache2 map[string]int = make(map[string]int)

func goBeam2(m mat, start coord) int {
	key := getCacheKey(start)
	cachedResult, ok := cache2[key]
	if ok {
		return cachedResult
	}

	it := start
	for it.Row < m.RowCount && m.At(it) != SPLITTER {
		it = it.GetNewCoord(direction.Down)
	}

	if it.Row == m.RowCount {
		result := 1
		cache2[key] = result
		return result
	}

	// if m.At(it) == SPLITTER {
	result := goBeam2(m, it.GetNewCoord(direction.Left)) + goBeam2(m, it.GetNewCoord(direction.Right))
	cache2[key] = result
	return result
}

func Second(lines []string) (strigifiedResult string) {
	m, startCoord := matrix.ParseStringMatrixAndGetStartingPoint(lines, func(value string) bool { return value == START })

	strigifiedResult = fmt.Sprint(goBeam2(m, *startCoord))
	return strigifiedResult
}

// 2122 too low
// 47857642990159 too low
// Your puzzle answer was 47857642990160.

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
