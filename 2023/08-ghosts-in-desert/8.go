// https://adventofcode.com/2023/day/8

package main

import (
	"aoc/utils"
	"fmt"
	"strings"
	"time"
)

const aocDay int = 8

type to struct {
	left, right string
}

func first(lines []string) {
	fmt.Println("--- First ---")

	stepCount := 0

	// region Parse input
	instructions := lines[0]

	m := map[string]to{}
	for _, line := range lines[2:] {
		var from, left, right string
		fmt.Sscanf(line, "%3s = (%3s, %3s)", &from, &left, &right)

		m[from] = to{left, right}
	}
	// endregion Parse input

	const startItem = "AAA"
	const stopItem = "ZZZ"
	current := startItem
	canStop := func() bool {
		return current == stopItem
	}

outerFor:
	for !canStop() {
		for _, leftOrRight := range instructions {
			if leftOrRight == 'L' {
				current = m[current].left
			} else if leftOrRight == 'R' {
				current = m[current].right
			} else {
				panic("Instruction is not L or R")
			}

			stepCount++

			if canStop() {
				break outerFor
			}
		}
	}

	fmt.Println(stepCount)
}

func second(lines []string) {
	fmt.Println("--- Second ---")

	// region Parse input
	instructions := lines[0]

	m := map[string]to{}
	for _, line := range lines[2:] {
		var from, left, right string
		fmt.Sscanf(line, "%3s = (%3s, %3s)", &from, &left, &right)

		m[from] = to{left, right}
	}

	items := utils.Keys(m)
	// endregion Parse input

	const startItemSuffix = 'A'
	const endItemSuffix = 'Z'
	startItems := utils.Filter(items, func(item string) bool {
		return strings.HasSuffix(item, string(startItemSuffix))
	})
	canStop := func(current string) bool {
		return strings.HasSuffix(current, string(endItemSuffix))
	}

	stepCounts := make([]int, len(startItems))

	for i, current := range startItems {
		stepCount := 0

	outerFor:
		for !canStop(current) {
			for _, leftOrRight := range instructions {
				if leftOrRight == 'L' {
					current = m[current].left
				} else if leftOrRight == 'R' {
					current = m[current].right
				} else {
					panic("Instruction is not L or R")
				}

				stepCount++

				if canStop(current) {
					break outerFor
				}
			}
		}

		stepCounts[i] = stepCount
	}

	result := stepCounts[0]
	for _, stepCount := range stepCounts[1:] {
		result = utils.LCM(stepCount, result)
	}

	fmt.Println(result)
}

// tried second with 281, but too low

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt()

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	first(lines)

	second(lines)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
