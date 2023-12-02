// https://adventofcode.com/2023/day/2

package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

const aocDay int = 2

func first(lines []string) {
	println("First ---")

	maxMap := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	validSums := 0

	for _, line := range lines {
		split := strings.Split(line, ": ")

		idStr := strings.Split(split[0], " ")[1]
		id, _ := strconv.Atoi(idStr)

		sets := strings.Split(split[1], "; ")

		invalidGame := false

	sets:
		for _, set := range sets {
			tokens := strings.Split(set, ", ")
			for _, token := range tokens {
				s := strings.Split(token, " ")
				count, _ := strconv.Atoi(s[0])
				color := s[1]

				if count > maxMap[color] {
					invalidGame = true
					break sets
				}
			}
		}

		if !invalidGame {
			validSums += id
		}
	}

	fmt.Println("valid sums", validSums)
}

func second(lines []string) {
	println("Second ---")

	powerSums := 0

	for _, line := range lines {
		maxMap := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		split := strings.Split(line, ": ")

		// idStr := strings.Split(split[0], " ")[1]
		// id, _ := strconv.Atoi(idStr)
		// fmt.Printf("--- game %d \n", id)

		sets := strings.Split(split[1], "; ")

		for _, set := range sets {
			tokens := strings.Split(set, ", ")
			for _, token := range tokens {
				countStr, color := utils.SplitIn2(token, " ")
				count, _ := strconv.Atoi(countStr)

				if count > maxMap[color] {
					maxMap[color] = count
				}
			}
		}

		// fmt.Printf("would be possible with red %d, green %d, blue %d\n", maxMap["red"], maxMap["green"], maxMap["blue"])
		powerSums += (maxMap["red"] * maxMap["green"] * maxMap["blue"])
	}

	fmt.Println("power sums", powerSums)
}

func main() {
	lines := utils.ReadLines(fmt.Sprintf("%d.in", aocDay))

	first(lines)

	second(lines)
}
