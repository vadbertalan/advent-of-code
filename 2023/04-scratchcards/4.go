// https://adventofcode.com/2023/day/4

package main

import (
	"aoc/utils"
	"aoc/utils/collections"
	"fmt"
	"strconv"
	"strings"
)

const aocDay int = 4

func contains(nrs []int, nr int) bool {
	for _, existingNr := range nrs {
		if nr == existingNr {
			return true
		}
	}
	return false
}

func first(lines []string) {
	println("First ---")

	points := 0

	for _, line := range lines {
		cardPoints := 0

		_, numbers := utils.SplitIn2(line, ": ")
		winningNrsStr, localNrsStr := utils.SplitIn2(numbers, " | ")

		winningNrsStrList := strings.Fields(winningNrsStr)
		winningNrs := []int{}
		for _, str := range winningNrsStrList {
			nr, _ := strconv.Atoi(str)
			winningNrs = append(winningNrs, nr)
		}

		localNrsStrList := strings.Fields(localNrsStr)
		for _, str := range localNrsStrList {
			localNr, _ := strconv.Atoi(str)

			if contains(winningNrs, localNr) {
				if cardPoints == 0 {
					cardPoints = 1
				} else {
					cardPoints = cardPoints << 1
				}
			}
		}

		points += cardPoints
	}

	fmt.Println("Points:", points)
}

func second(lines []string) {
	println("\nSecond ---")

	initialCardCount := len(lines)

	yieldCounts := map[int]int{}

	for _, line := range lines {
		cardYields := 0

		cardIdStr, numbers := utils.SplitIn2(line, ": ")

		cardId, _ := strconv.Atoi(strings.Fields(cardIdStr)[1])

		winningNrsStr, localNrsStr := utils.SplitIn2(numbers, " | ")

		winningNrsStrList := strings.Fields(winningNrsStr)
		winningNrs := []int{}
		for _, str := range winningNrsStrList {
			nr, _ := strconv.Atoi(str)
			winningNrs = append(winningNrs, nr)
		}

		localNrsStrList := strings.Fields(localNrsStr)
		for _, str := range localNrsStrList {
			localNr, _ := strconv.Atoi(str)

			if contains(winningNrs, localNr) {
				cardYields += 1
			}
		}

		yieldCounts[cardId] = cardYields
	}

	cardCount := 0

	cardStack := collections.Stack[int]{}
	for cardId := 1; cardId <= initialCardCount; cardId++ {
		cardStack.Push(cardId)
	}

	for !cardStack.IsEmpty() {
		processedCard := cardStack.Pop()

		cardCount += 1

		for card := processedCard + 1; card < processedCard+1+yieldCounts[processedCard]; card++ {
			if card <= initialCardCount {
				cardStack.Push(card)
			}
		}
	}

	fmt.Println("Nr of cards:", cardCount)
}

// Tried 9441994 but failed, it is too low

func main() {
	inputFileExtension := utils.GetInputFileExt(1)

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	first(lines)

	second(lines)
}
