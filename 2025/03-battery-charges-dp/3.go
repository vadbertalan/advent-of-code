// https://adventofcode.com/2025/day/3

package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"time"
)

const aocDay int = 3

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func First(lines []string) (strigifiedResult string) {
	result := 0

	for _, line := range lines {
		max := 0
		for i := 0; i < len(line)-1; i++ {
			for j := i + 1; j < len(line); j++ {
				nr, _ := strconv.Atoi(fmt.Sprintf("%c%c", line[i], line[j]))
				if nr > max {
					max = nr
				}
			}
		}
		result += max
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// Your puzzle answer was 17435

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func Second(lines []string) (strigifiedResult string) {
	result := 0

	const DIGIT_COUNT = 12

	for _, line := range lines {
		// Skipping 0 index to make code semantically easier to read
		dp := [100][DIGIT_COUNT + 1]int{}

		dp[0][1] = utils.Atoi(string(line[0]))

		for i := 1; i < len(line); i++ {
			currentDigit := utils.Atoi(string(line[i]))
			maxPossibleDigit := min(i+1, DIGIT_COUNT)

			for k := maxPossibleDigit; k > 1; k-- {
				nrOfKDigits := dp[i-1][k-1]*10 + currentDigit

				dp[i][k] = max(nrOfKDigits, dp[i-1][k])
			}

			dp[i][1] = max(dp[i-1][1], currentDigit)
		}

		result += dp[len(line)-1][DIGIT_COUNT]
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// Your puzzle answer was 172886048065379

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
