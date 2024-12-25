// https://adventofcode.com/2024/day/25

package main

import (
	"aoc/utils"
	"aoc/utils/matrix"
	"fmt"
	"time"
)

const aocDay int = 25

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func First(lines []string) (strigifiedResult string) {
	result := 0

	locks := [][5]int{}
	keys := [][5]int{}

	for i := 0; i < len(lines); i++ {
		j := i
		for lines[j] != "" && j < len(lines)-1 {
			j++
		}
		// for last matrix
		if j == len(lines)-1 {
			j++
		}
		m := matrix.ParseStringMatrix(lines[i:j])

		count := 0
		pins := [5]int{}
		for k := 0; k < m.ColumnCount; k++ {
			count = 0
			for l := 1; l < m.RowCount-1; l++ {
				if m.Values[l][k] == "#" {
					count++
				}
			}
			pins[k] = count
		}

		if m.Values[0][0] == "#" {
			locks = append(locks, pins)
		} else {
			keys = append(keys, pins)
		}

		i = j
	}

	for _, lock := range locks {
		for _, key := range keys {
			keyFits := true
			for i := 0; i < 5; i++ {
				if 5-lock[i] < key[i] {
					keyFits = false
					break
				}
			}
			if keyFits {
				result++
			}
		}
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// correct: 3201, example 3

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(1)

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	fmt.Println("--- First ---")
	result := First(lines)
	fmt.Println(result)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
