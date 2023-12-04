package main

import (
	"aoc/utils"
	"fmt"
)

const aocDay int = 999

func first(lines []string) {
	fmt.Println("First ---")

	result := 0

	for _, line := range lines {
		fmt.Println(line)

	}

	fmt.Println(result)
}

func second(lines []string) {
	fmt.Println("\nSecond ---")

}

func DYNmain() {
	inputFileExtension := utils.GetInputFileExt()

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	first(lines)

	second(lines)
}
