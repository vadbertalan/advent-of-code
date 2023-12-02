package main

import (
	"aoc/utils"
	"fmt"
	// "strconv"
	// "strings"
	// "unicode"
)

const aocDay int = 1

func first(lines []string) {
	println("First ---")

	for _, line := range lines {
		fmt.Println(line)
	}
}

func second(lines []string) {
	println("Second ---")
}

func main() {
	lines := utils.ReadLines(fmt.Sprintf("%d.in", aocDay))

	first(lines)

	second(lines)
}
