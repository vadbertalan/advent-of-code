// https://adventofcode.com/2023/day/10

package main

import (
	"aoc/utils-go"
	"aoc/utils-go/coordinate"
	"aoc/utils-go/direction"
	"aoc/utils-go/formulae"
	"aoc/utils-go/matrix"
	"fmt"
	"strings"
	"time"
)

const aocDay int = 10

const startMarkerChar = 'S'

type coord = coordinate.Coord
type mat = matrix.Matrix[pipeType]

type pipeType string

const (
	upDown    pipeType = "|"
	leftRight          = "-"
	downRight          = "L"
	downLeft           = "J"
	upRight            = "F"
	upLeft             = "7"
	ground             = "."
	start              = "S"
)

var verticalConnectorPipes = []pipeType{
	upDown,
}

func canConnect(from, to pipeType, dir direction.Direction) bool {
	switch dir {
	case direction.Up:
		switch from {
		case upDown:
			return utils.Contains([]pipeType{start, upDown, upRight, upLeft}, to)
		case leftRight:
			return false
		case downRight:
			return utils.Contains([]pipeType{start, upDown, upRight, upLeft}, to)
		case downLeft:
			return utils.Contains([]pipeType{start, upDown, upRight, upLeft}, to)
		case upRight:
			return false
		case upLeft:
			return false
		case ground:
			return false
		case start:
			return utils.Contains([]pipeType{upDown, upRight, upLeft}, to)
		}
	case direction.Right:
		switch from {
		case upDown:
			return false
		case leftRight:
			return utils.Contains([]pipeType{start, leftRight, upLeft, downLeft}, to)
		case downRight:
			return utils.Contains([]pipeType{start, leftRight, upLeft, downLeft}, to)
		case downLeft:
			return false
		case upRight:
			return utils.Contains([]pipeType{start, leftRight, upLeft, downLeft}, to)
		case upLeft:
			return false
		case ground:
			return false
		case start:
			return utils.Contains([]pipeType{leftRight, upLeft, downLeft}, to)
		}
	case direction.Down:
		switch from {
		case upDown:
			return utils.Contains([]pipeType{start, upDown, downLeft, downRight}, to)
		case leftRight:
			return false
		case downRight:
			return false
		case downLeft:
			return false
		case upRight:
			return utils.Contains([]pipeType{start, upDown, downLeft, downRight}, to)
		case upLeft:
			return utils.Contains([]pipeType{start, upDown, downLeft, downRight}, to)
		case ground:
			return false
		case start:
			return utils.Contains([]pipeType{upDown, downLeft, downRight}, to)
		}
	case direction.Left:
		switch from {
		case upDown:
			return false
		case leftRight:
			return utils.Contains([]pipeType{start, leftRight, upRight, downRight}, to)
		case downRight:
			return false
		case downLeft:
			return utils.Contains([]pipeType{start, leftRight, upRight, downRight}, to)
		case upRight:
			return false
		case upLeft:
			return utils.Contains([]pipeType{start, leftRight, upRight, downRight}, to)
		case ground:
			return false
		case start:
			return utils.Contains([]pipeType{leftRight, downRight, upRight}, to)
		}
	}

	return false
}

func getSecondPipeValidatorFn(m mat, startPipeCoord coord) func(pipe pipeType, neighborCoord coordinate.Coord, dir direction.Direction) bool {
	return func(pipe pipeType, neighborCoord coordinate.Coord, dir direction.Direction) bool {
		return canConnect(m.At(startPipeCoord), pipe, dir)
	}
}

func getNextPipeCoord(prevCoord, curCoord coord, m mat) *coord {
	testNext := func(nextPipe pipeType, nextCoord coord, dir direction.Direction) bool {
		return !nextCoord.IsEqual(prevCoord) && canConnect(m.At(curCoord), nextPipe, dir)
	}
	return m.GetFirstValidNeighbor(curCoord, testNext, false)
}

func first(lines []string) {
	fmt.Println("--- First ---")

	m := matrix.Matrix[pipeType]{
		Values:      make([][]pipeType, len(lines)),
		RowCount:    len(lines),
		ColumnCount: len(lines[0]),
	}

	var startPipeCoord coord

	for row, line := range lines {
		m.Values[row] = make([]pipeType, len(line))
		for col, c := range line {
			m.Values[row][col] = pipeType(string(c))

			if c == startMarkerChar {
				startPipeCoord = coord{Row: row, Col: col}
			}
		}
	}

	prevPipeCoord := &startPipeCoord
	currentPipeCoord := m.GetFirstValidNeighbor(startPipeCoord, getSecondPipeValidatorFn(m, startPipeCoord), false)

	coords := []coord{*prevPipeCoord, *currentPipeCoord}

	for !currentPipeCoord.IsEqual(startPipeCoord) {
		nextPipeCoord := getNextPipeCoord(*prevPipeCoord, *currentPipeCoord, m)
		prevPipeCoord = currentPipeCoord
		currentPipeCoord = nextPipeCoord
		coords = append(coords, *currentPipeCoord)
	}

	fmt.Println(len(coords) / 2)
}

func second(lines []string) {
	fmt.Println("\n--- Second ---")

	result := 0

	m := matrix.Matrix[pipeType]{
		Values:      make([][]pipeType, len(lines)),
		RowCount:    len(lines),
		ColumnCount: len(lines[0]),
	}

	var startPipeCoord coord

	for row, line := range lines {
		m.Values[row] = make([]pipeType, len(line))
		for col, c := range line {
			m.Values[row][col] = pipeType(string(c))
			if c == startMarkerChar {
				startPipeCoord = coord{Row: row, Col: col}
			}
		}
	}

	prevPipeCoord := &startPipeCoord
	currentPipeCoord := m.GetFirstValidNeighbor(startPipeCoord, getSecondPipeValidatorFn(m, startPipeCoord), false)

	polygonVertexes := []coord{startPipeCoord}
	boundaryCoords := []coord{*prevPipeCoord}

	for !currentPipeCoord.IsEqual(startPipeCoord) {
		if strings.Contains("7JLF", string(m.At(*currentPipeCoord))) {
			polygonVertexes = append(polygonVertexes, *currentPipeCoord)
		}
		boundaryCoords = append(boundaryCoords, *currentPipeCoord)
		nextPipeCoord := getNextPipeCoord(*prevPipeCoord, *currentPipeCoord, m)
		prevPipeCoord = currentPipeCoord
		currentPipeCoord = nextPipeCoord
	}

	A := formulae.CalcAreaShoelace(polygonVertexes)

	// fmt.Println(A)
	// fmt.Println(len(boundaryCoords))

	// Pick theorem: https://en.wikipedia.org/wiki/Pick%27s_theorem
	innerPoints := int(A - float64(len(boundaryCoords))/2 + 1)

	result = innerPoints
	fmt.Println(result)
}

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(4)

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	first(lines)

	second(lines)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
