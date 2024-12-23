// https://adventofcode.com/2023/day/9

package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const aocDay int = 9

func First(lines []string) int {
	fmt.Println("--- First ---")

	result := 0

	for _, line := range lines {
		nrStrList := strings.Fields(line)
		nrs := []int{}
		for _, str := range nrStrList {
			nr, _ := strconv.Atoi(str)
			nrs = append(nrs, nr)
		}

		nrsToAdd := []int{nrs[len(nrs)-1]}
		diffs := nrs
		for {
			newdiffs := make([]int, len(diffs)-1)
			for i := 0; i < len(diffs)-1; i++ {
				newdiffs[i] = diffs[i+1] - diffs[i]
			}
			nrsToAdd = append(nrsToAdd, newdiffs[len(newdiffs)-1])
			if utils.Every(newdiffs, func(item int) bool { return item == 0 }) {
				break
			}
			diffs = newdiffs

		}

		sum := 0
		for _, nr := range nrsToAdd {
			sum += nr
		}
		result += sum
	}

	return result
}

func Second(lines []string) int {
	fmt.Println("\n--- Second ---")

	result := 0

	for _, line := range lines {
		nrStrList := strings.Fields(line)
		nrs := []int{}
		for _, str := range nrStrList {
			nr, _ := strconv.Atoi(str)
			nrs = append(nrs, nr)
		}

		nrsToExt := []int{nrs[0]}
		diffs := nrs
		for {
			newdiffs := make([]int, len(diffs)-1)
			for i := 0; i < len(diffs)-1; i++ {
				newdiffs[i] = diffs[i+1] - diffs[i]
			}
			nrsToExt = append(nrsToExt, newdiffs[0])
			if utils.Every(newdiffs, func(item int) bool { return item == 0 }) {
				break
			}
			diffs = newdiffs

		}

		diff := 0
		for i := len(nrsToExt) - 1; i >= 0; i-- {
			diff = nrsToExt[i] - diff
		}
		result += diff
	}

	return result
}

// 228, wrong

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(1)

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	fmt.Println(First(lines))

	fmt.Println(Second(lines))

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
