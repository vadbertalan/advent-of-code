// https://adventofcode.com/2025/day/2

package main

import (
	"aoc/utils-go"
	"aoc/utils-go/collections"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const aocDay int = 2

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func isInvalidId(id string) bool {
	if len(id)%2 != 0 {
		return false
	}

	for i := 0; i < len(id)/2; i++ {
		if id[i] != id[i+len(id)/2] {
			return false
		}
	}

	return true
}

func First(lines []string) (strigifiedResult string) {
	result := 0

	ranges := strings.Split(lines[0], ",")

	for _, rang := range ranges {
		leftStr, rightStr := utils.SplitIn2(rang, "-")
		left, _ := strconv.Atoi(leftStr)
		right, _ := strconv.Atoi(rightStr)

		for i := left; i <= right; i++ {
			if isInvalidId(strconv.Itoa(i)) {
				result += i
			}
		}
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// Your puzzle answer was 34826702005

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func isInvalidByK(id string, k int) bool {
	if len(id)%k != 0 {
		return false
	}

	tokensSet := collections.NewSet[string]()

	for i := 0; i < len(id); i += k {
		tokensSet.Add(id[i : i+k])
	}

	return tokensSet.Size() == 1
}

func isInvalidId2(id string) bool {
	k := len(id) / 2

	for i := k; i >= 1; i-- {
		if isInvalidByK(id, i) {
			return true
		}
	}

	return false
}

func Second(lines []string) (strigifiedResult string) {
	result := 0

	ranges := strings.Split(lines[0], ",")

	for _, rang := range ranges {
		leftStr, rightStr := utils.SplitIn2(rang, "-")
		left, _ := strconv.Atoi(leftStr)
		right, _ := strconv.Atoi(rightStr)

		for i := left; i <= right; i++ {
			if isInvalidId2(strconv.Itoa(i)) {
				result += i
			}
		}
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// Your puzzle answer was 43287141963

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
