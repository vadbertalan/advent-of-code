// https://adventofcode.com/2025/day/6

package main

import (
	"aoc/utils-go"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const aocDay int = 6

func parseProblemsAndSigns(lines []string) (problems [][]int, signs []string) {
	for _, sign := range strings.Split(lines[len(lines)-1], " ") {
		if sign == "" {
			continue
		}
		signs = append(signs, sign)
	}

	for _, line := range lines[:len(lines)-1] {
		col := 0
		for _, num := range strings.Split(line, " ") {
			if num == "" {
				continue
			}
			if len(problems) <= col {
				problems = append(problems, []int{})
			}
			problems[col] = append(problems[col], utils.Atoi(num))
			col++
		}
	}
	return problems, signs
}

func solveProblem(problem []int, sign string) int {
	answer := 0
	if sign == "*" {
		answer = 1
	}
	for _, val := range problem {
		if sign == "+" {
			answer += val
		} else if sign == "*" {
			answer *= val
		} else {
			panic("unknown sign")
		}
	}
	return answer

}

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func First(lines []string) (strigifiedResult string) {
	result := 0

	problems, signs := parseProblemsAndSigns(lines)

	for i, problem := range problems {
		result += solveProblem(problem, signs[i])
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// Your puzzle answer was 5877594983578

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func parseMatrixParcel(m []string, rowStart, rowEnd, colStart, colEnd int) []int {
	values := []int{}
	for j := colStart; j < colEnd; j++ {
		num := 0
		for i := rowStart; i < rowEnd; i++ {
			c := m[i][j]
			if c != ' ' {
				num = num*10 + utils.Atoi(string(c))
			}
		}
		values = append(values, num)
	}
	return values
}

// Your puzzle answer was 11159825706149

func Second(strMatrix []string) (strigifiedResult string) {
	result := 0

	problems, signs := parseProblemsAndSigns(strMatrix)

	mCol := 0

	for problemIndex, problem := range problems {
		maxVal := problem[0]
		for _, val := range problem {
			if val > maxVal {
				maxVal = val
			}
		}
		lenOfMaxValue := len(strconv.Itoa(maxVal))

		realProblem := parseMatrixParcel(strMatrix, 0, len(strMatrix)-1, mCol, mCol+lenOfMaxValue)
		mCol += lenOfMaxValue + 1

		signOfProblem := signs[problemIndex]
		answer := solveProblem(realProblem, signOfProblem)
		result += answer
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// Your puzzle answer was 11159825706149

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
