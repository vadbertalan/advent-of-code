// https://adventofcode.com/2023/day/16

package main

import (
	"aoc/utils-go"
	"aoc/utils-go/collections"
	"aoc/utils-go/coordinate"
	"aoc/utils-go/direction"
	"aoc/utils-go/matrix"
	"fmt"
	"math"
	"time"
)

const aocDay int = 16

type coord = coordinate.Coord
type do = coordinate.DirOffset
type direc = direction.Direction
type mat = matrix.Matrix[string]

type itinerary struct {
	coord coord
	dir   direc
}

func goBeam(m mat, c coord, dir direc, seen *collections.Set[itinerary], energized *collections.Set[coord]) {
	it := itinerary{c, dir}
	if seen.Has(it) {
		return
	}

	for m.IsValidCoord(c) && m.At(c) == "." {
		seen.Add(it)
		energized.Add(c)
		c = c.GetNewCoord(dir)
	}

	if !m.IsValidCoord(c) {
		return
	}

	energized.Add(c)

	if m.At(c) == "|" {
		if dir == direction.Left || dir == direction.Right {
			goBeam(m, c.GetNewCoord(direction.Up), direction.Up, seen, energized)
			goBeam(m, c.GetNewCoord(direction.Down), direction.Down, seen, energized)
		} else {
			goBeam(m, c.GetNewCoord(dir), dir, seen, energized)
		}
	} else if m.At(c) == "-" {
		if dir == direction.Up || dir == direction.Down {
			goBeam(m, c.GetNewCoord(direction.Left), direction.Left, seen, energized)
			goBeam(m, c.GetNewCoord(direction.Right), direction.Right, seen, energized)
		} else {
			goBeam(m, c.GetNewCoord(dir), dir, seen, energized)
		}
	} else if m.At(c) == "/" {
		switch dir {
		case direction.Up:
			goBeam(m, c.GetNewCoord(direction.Right), direction.Right, seen, energized)
		case direction.Left:
			goBeam(m, c.GetNewCoord(direction.Down), direction.Down, seen, energized)
		case direction.Down:
			goBeam(m, c.GetNewCoord(direction.Left), direction.Left, seen, energized)
		case direction.Right:
			goBeam(m, c.GetNewCoord(direction.Up), direction.Up, seen, energized)
		}
	} else if m.At(c) == "\\" {
		switch dir {
		case direction.Up:
			goBeam(m, c.GetNewCoord(direction.Left), direction.Left, seen, energized)
		case direction.Left:
			goBeam(m, c.GetNewCoord(direction.Up), direction.Up, seen, energized)
		case direction.Down:
			goBeam(m, c.GetNewCoord(direction.Right), direction.Right, seen, energized)
		case direction.Right:
			goBeam(m, c.GetNewCoord(direction.Down), direction.Down, seen, energized)
		}
	}
}

func countEnergized(m mat, startCoord coord, startDir direc) int {
	seen := collections.NewSet[itinerary]()
	energized := collections.NewSet[coord]()

	goBeam(m, startCoord, startDir, seen, energized)

	return energized.Size()
}

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func First(lines []string) (strigifiedResult string) {
	fmt.Println("--- First ---")

	m := matrix.ParseStringMatrix(lines)

	result := countEnergized(m, coord{Row: 0, Col: 0}, direction.Right)

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func Second(lines []string) (strigifiedResult string) {
	fmt.Println("\n--- Second ---")

	m := matrix.ParseStringMatrix(lines)

	max := math.MinInt

	// Top row -> dir down
	for col := 0; col < m.ColumnCount; col++ {
		count := countEnergized(m, coord{Row: 0, Col: col}, direction.Down)
		// fmt.Println("starting from ", coord{Row: 0, Col: col}, " -> ", count)
		if count > max {
			max = count
		}
	}

	// Bottom row -> dir up
	for col := 0; col < m.ColumnCount; col++ {
		count := countEnergized(m, coord{Row: m.RowCount - 1, Col: col}, direction.Up)
		// fmt.Println("starting from ", coord{Row: m.RowCount - 1, Col: col}, " -> ", count)
		if count > max {
			max = count
		}
	}

	// First column -> dir right
	for row := 0; row < m.RowCount; row++ {
		count := countEnergized(m, coord{Row: row, Col: 0}, direction.Right)
		// fmt.Println("starting from ", coord{Row: row, Col: 0}, " -> ", count)
		if count > max {
			max = count
		}
	}

	// Rightmost column -> dir left
	for row := 0; row < m.RowCount; row++ {
		count := countEnergized(m, coord{Row: row, Col: m.ColumnCount - 1}, direction.Left)
		// fmt.Println("starting from ", coord{Row: row, Col: m.ColumnCount - 1}, " -> ", count)
		if count > max {
			max = count
		}
	}

	result := max
	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

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
