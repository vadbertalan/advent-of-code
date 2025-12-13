// https://adventofcode.com/2024/day/19

package main

import (
	"aoc/utils-go"
	"fmt"
	"strings"
	"time"
)

const aocDay int = 19

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

var cache = map[string]bool{}

func canDesignBeDisplayed(design string, towels []string) bool {
	if design == "" {
		return true
	}

	if val, ok := cache[design]; ok {
		return val
	}

	for _, towel := range towels {
		if remDesign, ok := strings.CutPrefix(design, towel); ok {
			result := canDesignBeDisplayed(remDesign, towels)

			cache[design] = result

			if result {
				return true
			}
		}
	}
	return false
}

func First(lines []string) (strigifiedResult string) {
	// Reset cache for tests
	cache = map[string]bool{}

	result := 0

	towels := strings.Split(lines[0], ", ")

	for _, design := range lines[2:] {
		if canDesignBeDisplayed(design, towels) {
			result++
		}
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// correct 298, example 6

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

var cache2 = map[string]int{}

func canDesignBeDisplayed2(design string, towels []string) int {
	if design == "" {
		return 1
	}

	if val, ok := cache2[design]; ok {
		return val
	}

	count := 0
	for _, towel := range towels {
		if remDesign, ok := strings.CutPrefix(design, towel); ok {
			result := canDesignBeDisplayed2(remDesign, towels)

			cache2[remDesign] = result

			count += result
		}
	}
	return count
}

func Second(lines []string) (strigifiedResult string) {
	// Reset cache for tests
	cache2 = map[string]int{}

	result := 0

	towels := strings.Split(lines[0], ", ")

	for _, design := range lines[2:] {
		localResult := canDesignBeDisplayed2(design, towels)
		result += localResult
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// 97500 incorrect
// 572248688842069 correct, example 16

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
