// https://adventofcode.com/2023/day/6

package main

import (
	"aoc/utils-go"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const aocDay int = 6

func first(lines []string) {
	fmt.Println("--- First ---")

	timeStrs := strings.Fields(lines[0])[1:]
	distanceStrs := strings.Fields(lines[1])[1:]

	times := []int{}
	for _, str := range timeStrs {
		nr, _ := strconv.Atoi(str)
		times = append(times, nr)
	}

	distances := []int{}
	for _, str := range distanceStrs {
		nr, _ := strconv.Atoi(str)
		distances = append(distances, nr)
	}

	winCounts := []int{}

	for j := 0; j < len(times); j++ {
		winCount := 0

		distToBeat := distances[j]
		mmpms := 0
		for hold := 0; hold <= times[j]; hold++ {
			timeLeft := times[j]

			// Hold the button
			mmpms = hold
			timeLeft -= hold

			// Let go
			dist := timeLeft * mmpms

			if dist > distToBeat {
				winCount++
			}
		}
		winCounts = append(winCounts, winCount)
	}

	mul := 1
	for _, wc := range winCounts {
		mul *= wc
	}
	fmt.Println("Multiplication result: ", mul)
}

func second_DUMB(lines []string) {
	fmt.Println("\n--- Second DUMB---")

	timeStrs := strings.Fields(lines[0])[1:]
	distanceStrs := strings.Fields(lines[1])[1:]

	timeAvailable, _ := strconv.Atoi(strings.Join(timeStrs, ""))
	distToBeat, _ := strconv.Atoi(strings.Join(distanceStrs, ""))

	winCount := 0

	for hold := 0; hold <= timeAvailable; hold++ {
		timeLeft := timeAvailable

		// Hold the button
		mmpms := hold
		timeLeft -= hold

		// Let go
		dist := timeLeft * mmpms

		if dist > distToBeat {
			winCount++
		}
	}

	fmt.Println("Win count: ", winCount)
}

func second(lines []string) {
	fmt.Println("\n--- Second ---")

	timeStrs := strings.Fields(lines[0])[1:]
	distanceStrs := strings.Fields(lines[1])[1:]

	timeAvailable, _ := strconv.Atoi(strings.Join(timeStrs, ""))
	distToBeat, _ := strconv.Atoi(strings.Join(distanceStrs, ""))

	// (T - x) * x >= D, x = hold amount, T = time available, D = distance to beat
	// ==> -x^2 + T * x - D >= 0, where rounded x1 and x2 roots are the min and max hold values
	x1, x2 := utils.SolveMasodfoku(-1, float64(timeAvailable), -float64(distToBeat))

	winCount := int(x2) - int(x1)

	fmt.Println("Win count: ", winCount)
}

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(1)

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	first(lines)

	second_DUMB(lines)

	second(lines)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
