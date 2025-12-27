// https://adventofcode.com/2025/day/10

package main

import (
	"aoc/utils-go"
	"aoc/utils-go/collections"
	"fmt"
	"math"
	"strings"
	"time"
)

const aocDay int = 10

const ON = '#'
const OFF = '.'

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func findIndex(s string, ch byte) int {
	for i := 0; i < len(s); i++ {
		if s[i] == ch {
			return i
		}
	}
	return -1
}

func findLastIndex(s string, ch byte) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ch {
			return i
		}
	}
	return -1
}

type button []int

func parseButtons(s string) []button {
	ret := []button{}
	buttonStrs := strings.Split(s, " ")
	for _, buttonStr := range buttonStrs {
		ret = append(ret, utils.ParseInts(buttonStr[1:len(buttonStr)-1], ","))
	}
	return ret
}

var allSolutions = [][]*button{}
var minSolLength = math.MaxInt

var backtrackArray collections.Set[*button]

func isValid(targetPattern string) bool {
	counts := make([]int, len(targetPattern))
	for _, btn := range backtrackArray.GetValues() {
		for _, btnIdx := range *btn {
			counts[btnIdx]++
		}
	}

	for i := 0; i < len(targetPattern); i++ {
		targetNr := 0
		if targetPattern[i] == '#' {
			targetNr = 1
		}
		if targetNr != (counts[i] % 2) {
			return false
		}
	}

	return true
}

func backtr(targetPattern string, possibilities collections.Set[*button], maxSize int) {
	ps := possibilities.GetValues()
	for _, ps := range ps {
		backtrackArray.Add(ps)

		if isValid(targetPattern) {
			allSolutions = append(allSolutions, backtrackArray.GetValues())
			if len(backtrackArray.GetValues()) < minSolLength {
				minSolLength = len(backtrackArray.GetValues())
			}
		}

		if backtrackArray.Size() < maxSize {
			possibilities.Remove(ps)
			backtr(targetPattern, possibilities, maxSize)
			possibilities.Add(ps)
		}

		backtrackArray.Remove(ps)
	}
}

func First(lines []string) (strigifiedResult string) {
	result := 0

	for _, line := range lines {
		// Parse lights: [.##.]
		lightsStart := findIndex(line, '[')
		lightsEnd := findIndex(line, ']')
		lights := line[lightsStart+1 : lightsEnd]

		// Parse buttons: (3) (1,3) (2) (2,3) (0,2) (0,1)
		buttonsStart := lightsEnd + 2
		buttonsEnd := findLastIndex(line, ')')
		buttonsStr := line[buttonsStart : buttonsEnd+1]
		buttons := parseButtons(buttonsStr)

		backtrackArray = *collections.NewSet[*button]()
		buttonPs := make([]*button, len(buttons))
		for i, btn := range buttons {
			buttonPs[i] = &btn
		}

		possibilities := *(collections.NewSetFromArray(buttonPs))

		var minButtonPresses int = -1

		for i := 1; i <= possibilities.Size(); i++ {
			allSolutions = [][]*button{}
			minSolLength = math.MaxInt

			backtr(lights, possibilities, i)

			if len(allSolutions) > 0 {
				minButtonPresses = i
				break
			}
		}

		// fmt.Println(minButtonPresses, utils.Map(allSolutions[0], func(btn *button) button { return *btn }))
		// fmt.Println("----")

		result += minButtonPresses
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// Your puzzle answer was 438

// Second part is solved in `2025/10/10.ts`

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(1)

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	fmt.Println("--- First ---")
	result := First(lines)
	fmt.Println(result)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
