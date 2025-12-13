// https://adventofcode.com/2024/day/7

package main

import (
	"aoc/utils-go"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const aocDay int = 7

// func printOps(leftNr int, rightNrs []int, ops []string) {
// 	fmt.Print(leftNr, " = ")
// 	for i, nr := range rightNrs {
// 		fmt.Print(nr, " ")
// 		if i < len(ops) {
// 			fmt.Print(ops[i], " ")
// 		}
// 	}
// 	fmt.Println()
// }

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func isValid1(leftNr int, rightNrs []int, ops []string, level, maxLevel int) bool {
	if level != maxLevel-1 {
		return false
	}

	result := rightNrs[0]
	for i := 0; i < maxLevel; i++ {
		if ops[i] == "+" {
			result += rightNrs[i+1]
		} else if ops[i] == "*" {
			result *= rightNrs[i+1]
		}
	}
	return result == leftNr
}

var canBeSolvedMap = map[int]bool{}

func backtrack1(leftNr int, rightNrs []int, ops []string, levelIndex, maxLevel int) {
	for _, sign := range []string{"+", "*"} {
		ops[levelIndex] = sign

		if levelIndex < maxLevel-1 {
			backtrack1(leftNr, rightNrs, ops, levelIndex+1, maxLevel)
		} else if isValid1(leftNr, rightNrs, ops, levelIndex, maxLevel) {
			// printOps(leftNr, rightNrs, ops)
			canBeSolvedMap[leftNr] = true
		}
	}
}

func First(lines []string) (strigifiedResult string) {
	result := 0

	for _, line := range lines {
		leftNrStr, rightNrsStrs := utils.SplitIn2(line, ": ")
		leftNr, _ := strconv.Atoi(leftNrStr)
		rightNrs := utils.ConvertToInts(strings.Split(rightNrsStrs, " "))

		maxLevel := len(rightNrs) - 1
		ops := make([]string, maxLevel)

		backtrack1(leftNr, rightNrs, ops, 0, maxLevel)

		if canBeSolvedMap[leftNr] {
			result += leftNr
		}
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// 19 not right
// 349 not right, I should really read the question
// 1708857123053 correct, example 3749

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func isValid2(leftNr int, rightNrs []int, ops []string, level, maxLevel int) bool {
	if level != maxLevel-1 {
		return false
	}

	result := rightNrs[0]
	for i := 0; i < maxLevel; i++ {
		if ops[i] == "+" {
			result += rightNrs[i+1]
		} else if ops[i] == "*" {
			result *= rightNrs[i+1]
		} else if ops[i] == "||" {
			newStr := strconv.Itoa(result) + strconv.Itoa(rightNrs[i+1])
			newNr, _ := strconv.Atoi(newStr)
			result = newNr
		}
	}
	return result == leftNr
}

var canBeSolvedMap2 = map[int]bool{}

func backtrack2(leftNr int, rightNrs []int, ops []string, levelIndex, maxLevel int) {
	for _, sign := range []string{"+", "*", "||"} {
		ops[levelIndex] = sign

		if levelIndex < maxLevel-1 {
			backtrack2(leftNr, rightNrs, ops, levelIndex+1, maxLevel)
		} else if isValid2(leftNr, rightNrs, ops, levelIndex, maxLevel) {
			// printOps(leftNr, rightNrs, ops)
			canBeSolvedMap2[leftNr] = true
		}
	}
}

func Second(lines []string) (strigifiedResult string) {
	result := 0

	for _, line := range lines {
		leftNrStr, rightNrsStrs := utils.SplitIn2(line, ": ")
		leftNr, _ := strconv.Atoi(leftNrStr)
		rightNrs := utils.ConvertToInts(strings.Split(rightNrsStrs, " "))

		maxLevel := len(rightNrs) - 1
		ops := make([]string, maxLevel)

		backtrack2(leftNr, rightNrs, ops, 0, maxLevel)

		if canBeSolvedMap2[leftNr] {
			result += leftNr
		}
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// 189207836795655 correct, example 11387

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(1)

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	fmt.Println("--- First ---")
	result := First(lines)
	fmt.Println(result)

	fmt.Println("\n--- Second ---")
	result = Second(lines)
	fmt.Println(result)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
