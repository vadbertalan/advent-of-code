// https://adventofcode.com/2024/day/3

package main

import (
	"aoc/utils"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

const aocDay int = 3

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func First(lines []string) (strigifiedResult string) {
	fmt.Println("--- First ---")

	s := ""

	for _, line := range lines {
		s += line
	}

	result := 0

	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := re.FindAllString(s, -1)

	for _, match := range matches {
		_, nrs := utils.SplitIn2(match, "(")
		first, nr2str := utils.SplitIn2(nrs, ",")
		second := nr2str[:len(nr2str)-1]

		nr1, _ := strconv.Atoi(first)
		nr2, _ := strconv.Atoi(second)

		result += nr1 * nr2
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// correct 161289189, example 161

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func Second(lines []string) (strigifiedResult string) {
	fmt.Println("\n--- Second ---")

	s := ""

	for _, line := range lines {
		s += line
	}

	result := 0

	re := regexp.MustCompile(`(mul\(\d+,\d+\)|do\(\)|don't\(\))`)
	matches := re.FindAllString(s, -1)

	enabled := true

	for _, match := range matches {
		if match == "do()" {
			enabled = true
		} else if match == "don't()" {
			enabled = false
		} else {
			if !enabled {
				continue
			}

			_, nrs := utils.SplitIn2(match, "(")
			first, nr2str := utils.SplitIn2(nrs, ",")
			second := nr2str[:len(nr2str)-1]

			nr1, _ := strconv.Atoi(first)
			nr2, _ := strconv.Atoi(second)

			result += nr1 * nr2
		}
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// correct 83595109, example 48

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(2)

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	result := First(lines)
	fmt.Println(result)

	result = Second(lines)
	fmt.Println(result)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
