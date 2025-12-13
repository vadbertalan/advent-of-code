// https://adventofcode.com/2024/day/5

package main

import (
	"aoc/utils-go"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const aocDay int = 5

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func First(lines []string) (strigifiedResult string) {
	fmt.Println("--- First ---")

	result := 0

	rules := [][]int{}
	pages := []string{}

	// Parse
	for i, line := range lines {
		if line == "" {
			pages = lines[i+1:]
			break
		}

		nr1str, nr2str := utils.SplitIn2(line, "|")
		nr1, _ := strconv.Atoi(nr1str)
		nr2, _ := strconv.Atoi(nr2str)

		rules = append(rules, []int{nr1, nr2})
	}

	// Solve
	for _, page := range pages {
		nrStrs := strings.Split(page, ",")
		nrs := utils.ConvertToInts(nrStrs)

		pageCorrect := true

		// Down and dirty
	out:
		for i := 0; i < len(nrs)-1; i++ {
			for j := i + 1; j < len(nrs); j++ {
				for _, rule := range rules {
					if rule[1] == nrs[i] && rule[0] == nrs[j] {
						pageCorrect = false
						break out
					}
				}
			}
		}

		if pageCorrect {
			result += nrs[len(nrs)/2]
		}
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// correct 5248, example 143

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func Second(lines []string) (strigifiedResult string) {
	fmt.Println("\n--- Second ---")

	result := 0

	rules := [][]int{}
	pages := []string{}

	// Parse
	for i, line := range lines {
		if line == "" {
			pages = lines[i+1:]
			break
		}

		nr1str, nr2str := utils.SplitIn2(line, "|")
		nr1, _ := strconv.Atoi(nr1str)
		nr2, _ := strconv.Atoi(nr2str)

		rules = append(rules, []int{nr1, nr2})
	}

	// Solve
	for _, page := range pages {
		nrStrs := strings.Split(page, ",")
		nrs := utils.ConvertToInts(nrStrs)

		pageCorrect := true

		// Down and dirty again 🫧
		for i := 0; i < len(nrs)-1; i++ {
			for j := i + 1; j < len(nrs); j++ {
				for _, rule := range rules {
					if rule[1] == nrs[i] && rule[0] == nrs[j] {
						nrs[j], nrs[i] = nrs[i], nrs[j]
						pageCorrect = false
					}
				}
			}
		}

		if !pageCorrect {
			result += nrs[len(nrs)/2]
		}
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// correct 4507, example 123

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(1)

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	result := First(lines)
	fmt.Println(result)

	result = Second(lines)
	fmt.Println(result)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
