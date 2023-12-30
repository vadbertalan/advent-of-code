// https://adventofcode.com/2023/day/11

package main

import (
	"aoc/utils"
	"aoc/utils/coordinate"
	"aoc/utils/matrix"
	"fmt"
	"math"
	"time"
)

const aocDay int = 11

type coord = coordinate.Coord
type mat = matrix.Matrix[string]

const galaxyChar = "#"
const emptyChar = "."

func manh_dist(c1, c2 coord) int {
	return int(math.Abs(float64(c1.Row)-float64(c2.Row)) + math.Abs(float64(c1.Col)-float64(c2.Col)))
}

func first(lines []string) {
	fmt.Println("--- First ---")

	result := 0

	m := matrix.Matrix[string]{
		Values:      make([][]string, len(lines)),
		RowCount:    len(lines),
		ColumnCount: len(lines[0]),
	}

	gCoords := []coord{}

	for row, line := range lines {
		m.Values[row] = make([]string, len(line))
		for col, c := range line {
			m.Values[row][col] = string(string(c))

			if string(c) == galaxyChar {
				gCoords = append(gCoords, coord{Row: row, Col: col})
			}
		}
	}

	emptyRows := []int{}

	for i := 0; i < m.RowCount; i++ {
		isRowEmtpy := true
		for j := 0; j < m.ColumnCount; j++ {
			if m.Values[i][j] != emptyChar {
				isRowEmtpy = false
			}
		}
		if isRowEmtpy {
			emptyRows = append(emptyRows, i)
		}
	}

	emptyCols := []int{}
	for i := 0; i < m.ColumnCount; i++ {
		isColEmtpy := true
		for j := 0; j < m.RowCount; j++ {
			if m.Values[j][i] != emptyChar {
				isColEmtpy = false
			}
		}
		if isColEmtpy {
			emptyCols = append(emptyCols, i)
		}
	}

	for i := 0; i < len(gCoords); i++ {
		for j := i + 1; j < len(gCoords); j++ {
			dist := manh_dist(gCoords[i], gCoords[j])

			extraRows := 0
			for _, x := range emptyRows {
				if x > utils.Min(gCoords[i].Row, gCoords[j].Row) && x < utils.Max(gCoords[i].Row, gCoords[j].Row) {
					extraRows++
				}

			}

			extraCols := 0
			for _, y := range emptyCols {
				if y > utils.Min(gCoords[i].Col, gCoords[j].Col) && y < utils.Max(gCoords[i].Col, gCoords[j].Col) {
					extraCols++
				}
			}

			result += dist + extraRows + extraCols
		}
	}

	fmt.Println(result)
}

func second(lines []string) {
	fmt.Println("\n--- Second ---")

	result := 0

	m := matrix.Matrix[string]{
		Values:      make([][]string, len(lines)),
		RowCount:    len(lines),
		ColumnCount: len(lines[0]),
	}

	gCoords := []coord{}

	for row, line := range lines {
		m.Values[row] = make([]string, len(line))
		for col, c := range line {
			m.Values[row][col] = string(string(c))

			if string(c) == galaxyChar {
				gCoords = append(gCoords, coord{Row: row, Col: col})
			}
		}
	}

	emptyRows := []int{}

	for i := 0; i < m.RowCount; i++ {
		isRowEmtpy := true
		for j := 0; j < m.ColumnCount; j++ {
			if m.Values[i][j] != emptyChar {
				isRowEmtpy = false
			}
		}
		if isRowEmtpy {
			emptyRows = append(emptyRows, i)
		}
	}

	emptyCols := []int{}
	for i := 0; i < m.ColumnCount; i++ {
		isColEmtpy := true
		for j := 0; j < m.RowCount; j++ {
			if m.Values[j][i] != emptyChar {
				isColEmtpy = false
			}
		}
		if isColEmtpy {
			emptyCols = append(emptyCols, i)
		}
	}

	const galaxyMultiplier = 1000000

	for i := 0; i < len(gCoords); i++ {
		for j := i + 1; j < len(gCoords); j++ {
			dist := manh_dist(gCoords[i], gCoords[j])

			extraRows := 0
			for _, x := range emptyRows {
				if x > utils.Min(gCoords[i].Row, gCoords[j].Row) && x < utils.Max(gCoords[i].Row, gCoords[j].Row) {
					extraRows += galaxyMultiplier - 1
				}

			}

			extraCols := 0
			for _, y := range emptyCols {
				if y > utils.Min(gCoords[i].Col, gCoords[j].Col) && y < utils.Max(gCoords[i].Col, gCoords[j].Col) {
					extraCols += galaxyMultiplier - 1
				}
			}

			result += dist + extraRows + extraCols
		}
	}

	fmt.Println(result)

}

// tried 597714715262, wrong

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt()

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	first(lines)

	second(lines)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
