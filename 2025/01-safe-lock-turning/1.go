// https://adventofcode.com/2025/day/1

package main

import (
	"aoc/utils-go"
	"fmt"
	"strconv"
	"time"
)

const aocDay int = 1

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func First(lines []string) (strigifiedResult string) {
	result := 0

	tick := 50

	for _, line := range lines {
		amount, _ := strconv.ParseInt(line[1:], 10, 64)

		if line[0] == 'L' {
			tick -= int(amount)
			for tick < 0 {
				tick += 100
			}
		}

		if line[0] == 'R' {
			tick += int(amount)
			for tick >= 100 {
				tick -= 100
			}
		}

		if tick == 0 {
			result++
		}
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// 204 wrong
// Your puzzle answer was 980

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func Second(lines []string) (strigifiedResult string) {
	result := 0

	prevTick := -1
	tick := 50

	for _, line := range lines {
		amount, _ := strconv.ParseInt(line[1:], 10, 64)

		prevTick = tick

		if line[0] == 'L' {
			tick -= int(amount)
			for tick < 0 {
				if prevTick != 0 {
					result++
				}
				prevTick = tick
				tick += 100
			}

			if tick == 0 {
				result++
			}
		}

		if line[0] == 'R' {
			tick += int(amount)
			for tick >= 100 {
				result++
				tick -= 100
			}
		}
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// 5956 wrong
// 5467 wrong
// 4923 too low

// Your puzzle answer was 5961

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
