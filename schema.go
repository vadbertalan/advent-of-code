// https://adventofcode.com/yyyy/day/dd

package main

import (
	"aoc/utils"
	"aoc/utils/coordinate"
	"aoc/utils/matrix"
	"fmt"
	"time"
)

const aocDay int = 999

type coord = coordinate.Coord
type mat = matrix.Matrix[string]

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func First(lines []string) (strigifiedResult string) {
	result := 0

	for _, line := range lines {
		fmt.Println(line)

	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func Second(lines []string) (strigifiedResult string) {
	result := 0

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

func DYNmain() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(1)

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	result := First(lines)
	fmt.Println("--- First ---")
	fmt.Println(result)

	result = Second(lines)
	fmt.Println("\n--- Second ---")
	fmt.Println(result)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
