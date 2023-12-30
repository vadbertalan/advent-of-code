// https://adventofcode.com/2023/day/15

package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const aocDay int = 15

func hash(str string) int {
	ret := 0

	for _, c := range str {
		ret += int(c)
		ret *= 17
		ret %= 256
	}

	return ret
}

func first(lines []string) {
	fmt.Println("--- First ---")

	result := 0

	tokens := strings.Split(lines[0], ",")

	for _, token := range tokens {
		result += hash(token)
	}

	fmt.Println(result)
}

type commandType string

const (
	minus commandType = "-"
	equal             = "="
)

type command struct {
	typ  commandType
	lens *lens
}

type lens struct {
	label string
	focal int
}

func parseToken(tokenStr string) command {
	if strings.Contains(tokenStr, string(minus)) {
		label := strings.Split(tokenStr, string(minus))[0]
		return command{minus, &lens{label: label}}
	}
	if strings.Contains(tokenStr, string(equal)) {
		label, focalStr := utils.SplitIn2(tokenStr, string(equal))
		focal, _ := strconv.Atoi(focalStr)
		return command{equal, &lens{label: label, focal: focal}}
	}
	panic(fmt.Sprintf("could not parse token %s", tokenStr))
}

func second(lines []string) {
	fmt.Println("\n--- Second ---")

	tokens := strings.Split(lines[0], ",")

	const boxesLen = 256

	boxes := make([][]lens, boxesLen)
	for i := 0; i < boxesLen; i++ {
		boxes[i] = []lens{}
	}

	// do the steps and store lens
	for _, token := range tokens {
		cmd := parseToken(token)
		boxI := hash(cmd.lens.label)
		existingItem, index := utils.Find(boxes[boxI], func(lens lens) bool {
			return lens.label == cmd.lens.label
		})
		if cmd.typ == equal {
			if existingItem != nil {
				existingItem.focal = cmd.lens.focal
			} else {
				boxes[boxI] = append(boxes[boxI], *cmd.lens)
			}
		} else { // minus
			if existingItem != nil {
				boxes[boxI] = utils.RemoveOnIndexOrderPreserved(boxes[boxI], index)
			}
		}
	}

	// Calc lens score
	result := 0
	for boxNr, box := range boxes {
		for slotNumberOfLens, lens := range box {
			result += ((boxNr + 1) * (slotNumberOfLens + 1) * (lens.focal))
		}
	}

	fmt.Println(result)
}

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(1)

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	first(lines)

	second(lines)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
