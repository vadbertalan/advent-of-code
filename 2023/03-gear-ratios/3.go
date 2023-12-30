// https://adventofcode.com/2023/day/3

package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"unicode"
)

const aocDay int = 3

type matrix struct {
	values [][]rune
	n      int
	m      int
}

//region First part utilities

type partNrCand struct {
	nr       int
	isPartNr bool
}

func detIfPartNr(x int, ys []int, mat matrix) bool {
	for _, y := range ys {
		if isNeighboringSpecialChar(x, y, mat) {
			return true
		}
	}
	return false
}

func isNeighboringSpecialChar(x, y int, mat matrix) bool {
	xAdd := []int{0, 1, 1, 1, 0, -1, -1, -1}
	yAdd := []int{1, 1, 0, -1, -1, -1, 0, 1}
	for i := 0; i < 8; i++ {
		xx := x + xAdd[i]
		yy := y + yAdd[i]
		if isValidCoord(xx, yy, mat) && !unicode.IsDigit(mat.values[xx][yy]) && mat.values[xx][yy] != '.' {
			return true
		}
	}
	return false
}

// end region First part utilities

func first(lines []string) {
	println("First ---")

	var n, m int
	n = len(lines)
	m = len(lines[0])

	var matValues = make([][]rune, n)
	var nrs []partNrCand

	// Record matrix rune values
	for lineInd, line := range lines {
		matValues[lineInd] = make([]rune, m)
		for charInd, char := range line {
			matValues[lineInd][charInd] = char
		}
	}

	mat := matrix{matValues, n, m}

	for lineInd, line := range lines {
		nr := 0
		wasDigit := false
		indices := []int{}

		for charInd, char := range line {
			if unicode.IsDigit(char) {
				wasDigit = true
				digit, _ := strconv.Atoi(string(char))
				nr = nr*10 + digit
				indices = append(indices, charInd)
			} else {
				if wasDigit {
					isPartNr := detIfPartNr(lineInd, indices, mat)
					nrs = append(nrs, partNrCand{nr, isPartNr})

					nr = 0
					indices = []int{}
				}

				wasDigit = false
			}
		}

		if wasDigit {
			isPartNr := detIfPartNr(lineInd, indices, mat)
			nrs = append(nrs, partNrCand{nr, isPartNr})
		}
	}

	sum := 0
	for _, nr := range nrs {
		if nr.isPartNr {
			sum += nr.nr
		}
	}
	fmt.Println("Result:", sum)
}

// Second part utilities

const gearSymbol = '*'

type coord struct {
	x, y int
}

type gearCand struct {
	values []int
	coord  coord
}

func isValidCoord(x, y int, mat matrix) bool {
	return x >= 0 && x < mat.n && y >= 0 && y < mat.m
}

func getNeighboringStarCoord(x, y int, mat matrix) (coords []*coord) {
	xAdd := []int{0, 1, 1, 1, 0, -1, -1, -1}
	yAdd := []int{1, 1, 0, -1, -1, -1, 0, 1}
	for i := 0; i < 8; i++ {
		xx := x + xAdd[i]
		yy := y + yAdd[i]
		if isValidCoord(xx, yy, mat) && mat.values[xx][yy] == gearSymbol {
			coords = append(coords, &coord{xx, yy})
		}
	}
	return coords
}

func contains(coords []coord, coord coord) bool {
	for _, existingCoord := range coords {
		if coord.x == existingCoord.x && coord.y == existingCoord.y {
			return true
		}
	}
	return false
}

func getGearCandCoords(x int, ys []int, mat matrix) (gearCandCoords []coord) {
	for _, y := range ys {
		coords := getNeighboringStarCoord(x, y, mat)
		for _, coord := range coords {
			if coord != nil {
				// There might be multiple connections to the same gear
				if !contains(gearCandCoords, *coord) {
					gearCandCoords = append(gearCandCoords, *coord)
				}
			} else {
				panic("Oh no, this is bad! This pointer should not be nil")
			}
		}
	}
	return gearCandCoords
}

func getGearId(gearCoord coord) string {
	return fmt.Sprintf("%d-%d", gearCoord.x, gearCoord.y)
}

// Gather neighbor gear coordinates of this number and for that
// gear object register this number
func processFinishedNumber(nr int, lineInd int, indices []int, gearCands map[string]*gearCand, mat matrix) {
	gearCandCoords := getGearCandCoords(lineInd, indices, mat)
	for _, gearCandCoord := range gearCandCoords {
		gearKey := getGearId(gearCandCoord)
		gearCand, ok := gearCands[gearKey]
		if ok {
			// `nr` is for sure a part number as it's neighbor with a special char (with the gear *)
			gearCand.values = append(gearCand.values, nr)
		} else {
			panic(fmt.Sprintf("Oh uh! All gear cands should be registered in the map already. %v", gearCandCoord))
		}
	}
}

// end region Second part utilities

func second(lines []string) {
	println("Second ---")

	var n, m int
	n = len(lines)
	m = len(lines[0])

	var matValues = make([][]rune, n)
	gearCands := make(map[string]*gearCand)

	// Record matrix rune values
	for lineInd, line := range lines {
		matValues[lineInd] = make([]rune, m)
		for charInd, char := range line {
			matValues[lineInd][charInd] = char

			if char == gearSymbol {
				gearCoord := coord{lineInd, charInd}
				gearCands[getGearId(gearCoord)] = &gearCand{[]int{}, gearCoord}
			}
		}
	}

	mat := matrix{matValues, n, m}

	for lineInd, line := range lines {
		nr := 0
		wasDigit := false
		indices := []int{}

		for charInd, char := range line {
			if unicode.IsDigit(char) {
				wasDigit = true
				digit, _ := strconv.Atoi(string(char))
				nr = nr*10 + digit
				indices = append(indices, charInd)
			} else {
				if wasDigit {
					processFinishedNumber(nr, lineInd, indices, gearCands, mat)

					nr = 0
					indices = []int{}
				}

				wasDigit = false
			}
		}

		if wasDigit {
			processFinishedNumber(nr, lineInd, indices, gearCands, mat)
		}
	}

	sum := 0
	for _, gearCand := range gearCands {
		if len(gearCand.values) == 2 {
			sum += gearCand.values[0] * gearCand.values[1]
		}
	}

	fmt.Println("Result:", sum)
}

// I tried for second 83364799, but too low

func main() {
	inputFileExtension := utils.GetInputFileExt(1)

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	first(lines)

	second(lines)
}
