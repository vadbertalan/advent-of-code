package main

import (
	"aoc/utils-go"
	"fmt"
	"strconv"
	"strings"
)

const aocDay int = 2

type rule struct {
	Letter rune
	Min    int
	Max    int
}

type pwEntry struct {
	rul rule
	pw  string
}

func parseLine(line string) pwEntry {
	split := strings.Split(line, ": ")

	ruleSplit := strings.Split(split[0], " ")
	pw := split[1]

	minmax := strings.Split(ruleSplit[0], "-")
	letter := ruleSplit[1][0]

	min, _ := strconv.Atoi(minmax[0])
	max, _ := strconv.Atoi(minmax[1])

	r := new(rule)
	r.Letter = rune(letter)
	r.Min = min
	r.Max = max
	return pwEntry{*r, pw}
}

func parseLines(lines []string) (entries []pwEntry) {
	for _, line := range lines {
		entries = append(entries, parseLine(line))
	}
	return entries
}

func isPwEntryValid1(pwe pwEntry) bool {
	count := strings.Count(pwe.pw, string(pwe.rul.Letter))
	return count >= pwe.rul.Min && count <= pwe.rul.Max
}

func first(lines []string) {
	println("First ---")
	validCount := 0
	for _, e := range parseLines((lines)) {
		if isPwEntryValid1(e) {
			validCount++
		}
	}
	fmt.Println("valid count =", validCount)
}

func isPwEntryValid2(pwe pwEntry) bool {
	var isFirst int = 0
	var isSecond int = 0
	if pwe.pw[pwe.rul.Min-1] == byte(pwe.rul.Letter) {
		isFirst = 1
	}
	if pwe.pw[pwe.rul.Max-1] == byte(pwe.rul.Letter) {
		isSecond = 1
	}
	// XOR
	if isFirst^isSecond == 1 {
		return true
	}
	return false
}

func second(lines []string) {
	println("Second ---")

	validCount := 0
	for _, e := range parseLines((lines)) {
		if isPwEntryValid2(e) {
			validCount++
		}
	}
	fmt.Println("valid count =", validCount)
}

func main() {
	lines := utils.ReadLines(fmt.Sprintf("%d.in", aocDay))

	first(lines)

	second(lines)
}
