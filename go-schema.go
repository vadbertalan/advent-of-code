package main

import (
	"aoc/utils"
	"fmt"
	// "strconv"
	// "strings"
	// "unicode"
)

const aocDay int = 1

func first(input []int) {
	println("First ---")
}

func second(input []int) {
	println("Second ---")
}

func main() {
	lines := utils.ReadLines(fmt.Sprintf("%d.in", aocDay))
	ints := utils.ConvertToInts(lines)

	// fmt.Println(ints)

	first(ints)

	second(ints)
}
