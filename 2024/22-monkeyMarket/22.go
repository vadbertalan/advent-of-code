// https://adventofcode.com/2024/day/22

package main

import (
	"aoc/utils-go"
	"aoc/utils-go/collections"
	"fmt"
	"math"
	"strconv"
	"time"
)

const aocDay int = 22

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func mix(secretNumber, value int) int {
	xorResult := secretNumber ^ value
	return xorResult
}

func prune(secretNumber int) int {
	return secretNumber % 16777216
}

func proceed(secretNumber int) int {
	step1Result := secretNumber * 64
	secretNumber = prune(mix(secretNumber, step1Result))

	step2Result := secretNumber / 32
	secretNumber = prune(mix(secretNumber, step2Result))

	step3Result := secretNumber * 2048
	secretNumber = prune(mix(secretNumber, step3Result))

	return secretNumber
}

func First(lines []string) (strigifiedResult string) {
	result := 0
	const NR_OF_SECRET_NUMBERS = 2000

	for _, line := range lines {
		secretNumber, _ := strconv.Atoi(line)
		for i := 0; i < NR_OF_SECRET_NUMBERS; i++ {
			secretNumber = proceed(secretNumber)
		}
		result += secretNumber
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// correct: 16299144133, example 37327623

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func Second(lines []string) (strigifiedResult string) {
	result := 0
	const NR_OF_SECRET_NUMBERS = 2000

	monkeyCount := len(lines)

	setOfSequences := collections.NewSet[string]()

	sequences := map[int]map[string]int{}

	for i, line := range lines {
		subsequence := [4]int{}
		sequences[i+1] = map[string]int{}
		secretNumber, _ := strconv.Atoi(line)
		prevPrice := secretNumber % 10
		for j := 0; j < NR_OF_SECRET_NUMBERS; j++ {
			secretNumber = proceed(secretNumber)
			price := secretNumber % 10

			if j >= 3 {
				if j == 3 {
					subsequence[3] = price - prevPrice
				} else {
					subsequence[0] = subsequence[1]
					subsequence[1] = subsequence[2]
					subsequence[2] = subsequence[3]
					subsequence[3] = price - prevPrice
				}
				subseqStr := fmt.Sprintf("%d,%d,%d,%d", subsequence[0], subsequence[1], subsequence[2], subsequence[3])

				setOfSequences.Add(subseqStr)

				if _, ok := sequences[i+1][subseqStr]; !ok {
					sequences[i+1][subseqStr] = price
				}
			} else {
				subsequence[j] = price - prevPrice
			}

			prevPrice = price
		}
	}

	max := math.MinInt64
	for _, seq := range setOfSequences.GetValues() {
		bananaCount := 0
		for i := 1; i <= monkeyCount; i++ {
			if _, ok := sequences[i][seq]; ok {
				bananaCount += sequences[i][seq]
			}
		}
		if bananaCount > max {
			max = bananaCount
		}
	}
	result = max

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// correct: 1896 (best sequence "0,-1,-1,2"), example 23

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(2)

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
