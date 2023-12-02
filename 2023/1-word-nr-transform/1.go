// https://adventofcode.com/2023/day/1

package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const aocDay int = 1

func first(lines []string) {
	println("First ---")

	sum := 0
	for _, line := range lines {
		digitChars := ""
		for _, c := range line {
			if unicode.IsDigit(c) {
				digitChars = digitChars + string(c)
			}
		}

		nr := 0
		if len(digitChars) >= 0 {
			nr, _ = strconv.Atoi(string(digitChars[0]) + string(digitChars[len(digitChars)-1]))
		}

		sum += nr
	}

	fmt.Println("Result:", sum)
}

type indexNrPair struct {
	index int
	nr    int
}

func second(lines []string) {
	println("Second ---")

	wordNrMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	sum := 0
	for _, line := range lines {
		// Collect indexes of word numbers and the corr. numbers
		var iNrPairs []indexNrPair = []indexNrPair{}
		for wordNr, nr := range wordNrMap {
			tempLine := line
			for i := 0; i < strings.Count(line, wordNr); i++ {
				ind := strings.Index(tempLine, wordNr)

				if ind != -1 {
					iNrPairs = append(iNrPairs, indexNrPair{index: ind, nr: nr})
				}

				tempLine = strings.Replace(tempLine, wordNr, strings.Repeat("X", len(wordNr)), 1)
			}
		}

		// Insert corr. nrs: "twoeighttwo" -> "2wo8ight2wo"
		if len(iNrPairs) > 0 {
			for _, iNrPair := range iNrPairs {
				line = line[:iNrPair.index] + fmt.Sprint(iNrPair.nr) + line[iNrPair.index+1:]
			}
		}

		// Do part1 logic
		digitChars := ""
		for _, c := range line {
			if unicode.IsDigit(c) {
				digitChars = digitChars + string(c)
			}
		}

		nr := 0
		if len(digitChars) > 0 {
			nr, _ = strconv.Atoi(string(digitChars[0]) + string(digitChars[len(digitChars)-1]))
		}

		// fmt.Println(originalLine, line, nr)

		sum += nr
	}

	fmt.Println("Result:", sum)
}

func main() {
	lines := utils.ReadLines(fmt.Sprintf("%d.in", aocDay))

	first(lines)

	second(lines)
}
