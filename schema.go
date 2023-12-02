package main

import (
	"aoc/utils"
	"fmt"
	// "strconv"
	// "strings"
	// "unicode"
)

const aocDay int = 999

func first(lines []string) {
	println("First ---")

	for _, line := range lines {
		fmt.Println(line)
	}
}

func second(lines []string) {
	println("Second ---")
}

func DYNmain() {
	inputFileExtension := utils.GetInputFileExt()

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	first(lines)

	second(lines)
}
