package main

import (
	"aoc/utils"
	"fmt"
	"time"
)

const aocDay int = 999

func first(lines []string) {
	fmt.Println("--- First ---")

	result := 0

	for _, line := range lines {
		fmt.Println(line)

	}

	fmt.Println(result)
}

func second(lines []string) {
	fmt.Println("\n--- Second ---")

}

func DYNmain() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt()

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	first(lines)

	second(lines)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
