// https://adventofcode.com/2023/day/5

package main

import (
	"aoc/utils"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const aocDay int64 = 5

func parseNrLine1(str string) (int, int, int) {
	split := strings.Fields(str)
	dest, _ := strconv.Atoi(split[0])
	source, _ := strconv.Atoi(split[1])
	count, _ := strconv.Atoi(split[2])
	return dest, source, count
}

func parseNrLine2(str string) (int64, int64, int64) {
	split := strings.Fields(str)
	// dest, _ := strconv.Atoi(split[0])
	// source, _ := strconv.Atoi(split[1])
	// count, _ := strconv.Atoi(split[2])
	dest, _ := strconv.ParseInt(split[0], 10, 64)
	source, _ := strconv.ParseInt(split[1], 10, 64)
	count, _ := strconv.ParseInt(split[2], 10, 64)
	return int64(dest), int64(source), int64(count)
}

type dsc struct {
	dest, source, count, diff int
}

func parseSection(sectionStr string) func(v int) int {
	split := strings.Split(sectionStr, "\n")
	dscArray := []dsc{}
	for _, line := range split[1:] {
		dest, source, count := parseNrLine1(line)
		dscArray = append(dscArray, dsc{dest, source, count, dest - source})
	}
	return func(v int) int {
		for _, dscItem := range dscArray {
			if v >= dscItem.source && v < dscItem.source+dscItem.count {
				// return dscItem.dest + (v - dscItem.source)
				return v + dscItem.diff
			}
		}
		return v
	}
}

func first(input string) {
	fmt.Println("--- First ---")

	minLoc := 99999999999999

	sections := strings.Split(input, "\n\n")

	_, seedsStr := utils.SplitIn2(sections[0], ": ")
	seedStrs := strings.Fields(seedsStr)
	seeds := []int{}
	for _, str := range seedStrs {
		nr, _ := strconv.Atoi(str)
		seeds = append(seeds, nr)
	}

	toSoil := parseSection(sections[1])
	toFert := parseSection(sections[2])
	toWater := parseSection(sections[3])
	toLight := parseSection(sections[4])
	toTemp := parseSection(sections[5])
	toHum := parseSection(sections[6])
	toLoc := parseSection(sections[7])

	for _, seed := range seeds {
		loc := toLoc(toHum(toTemp(toLight(toWater(toFert(toSoil(seed)))))))
		if loc < minLoc {
			minLoc = loc
		}
	}

	fmt.Println("Min location:", minLoc)
}

// tried with, 1901620491 too high

// Part 2 utilities

type funcRule struct {
	min, max, diff int64
}

type pointOnLine struct {
	value int64
	rule  funcRule
}

func insertRuleSorted(sortedRules []pointOnLine, rule funcRule) []pointOnLine {
	i := sort.Search(len(sortedRules), func(i int) bool { return sortedRules[i].value > rule.min })
	sortedRules = append(sortedRules, pointOnLine{})
	copy(sortedRules[i+1:], sortedRules[i:])
	// here it would be enough to pass value only
	sortedRules[i] = pointOnLine{value: rule.min, rule: rule}

	i = sort.Search(len(sortedRules), func(i int) bool { return sortedRules[i].value > rule.max })
	sortedRules = append(sortedRules, pointOnLine{})
	copy(sortedRules[i+1:], sortedRules[i:])
	// here it would be enough to pass value only
	sortedRules[i] = pointOnLine{value: rule.max, rule: rule}

	return sortedRules
}

func Add64(left, right int64) int64 {
	if right > 0 {
		if left > math.MaxInt64-right {
			panic("errorrrrrr flw")
		}
	} else {
		if left < math.MinInt64-right {
			panic("errorrrrrr flw")
		}
	}
	return left + right
}

func compose2Funcs(fRules, gRules []funcRule) []funcRule {
	// Order interval boundaries (A start of interval, a end of interval) on line like this:
	//
	// ---[-----)[-----[------)[-----)--------)---
	//
	//	  A     aB     C      cD     d        b
	pointsOnLine := []pointOnLine{}

	for _, fRule := range fRules {
		pointsOnLine = insertRuleSorted(pointsOnLine, fRule)
	}
	for _, gRule := range gRules {
		pointsOnLine = insertRuleSorted(pointsOnLine, gRule)
	}

	var allFuncRules []funcRule = make([]funcRule, len(fRules)+len(gRules))
	copy(allFuncRules[:len(fRules)], fRules)
	copy(allFuncRules[len(fRules):], gRules)

	newFuncRules := []funcRule{}
	for i := 0; i < len(pointsOnLine)-1; i++ {
		var diff int64 = 0

		for j := 0; j < len(allFuncRules); j++ {
			if allFuncRules[j].min <= pointsOnLine[i].value && allFuncRules[j].max > pointsOnLine[i].value {
				diff += allFuncRules[j].diff
				// diff = Add64(diff, allFuncRules[j].diff)
			}
		}

		if pointsOnLine[i].value != pointsOnLine[i+1].value {
			funcRule := funcRule{min: pointsOnLine[i].value, max: pointsOnLine[i+1].value, diff: diff}
			newFuncRules = append(newFuncRules, funcRule)
		}
	}
	return newFuncRules
}

func parseSection2(sectionStr string) []funcRule {
	split := strings.Split(sectionStr, "\n")

	rules := []funcRule{}

	for _, line := range split[1:] {
		dest, source, count := parseNrLine2(line)
		rules = append(rules, funcRule{min: source, max: source + count, diff: dest - source})
	}

	return rules
}

func funcRules2Func(funcRules []funcRule) func(x int64) int64 {
	return func(x int64) int64 {
		for _, rule := range funcRules {
			if x >= rule.min && x < rule.max {
				// return Add64(x, rule.diff)
				return x + rule.diff
			}
		}
		return x
	}
}

func second(input string) {
	fmt.Println("\n--- Second ---")

	var minLoc int64 = math.MaxInt64

	sections := strings.Split(input, "\n\n")

	_, seedsStr := utils.SplitIn2(sections[0], ": ")
	seedStrs := strings.Fields(seedsStr)
	seedRanges := []int64{}
	for _, str := range seedStrs {
		nr, _ := strconv.ParseInt(str, 10, 64)
		seedRanges = append(seedRanges, int64(nr))
	}

	// seeds := []int64{}
	// for i := 0; i < len(seedRanges); i += 2 {
	// 	for j := seedRanges[i]; j < seedRanges[i]+seedRanges[i+1]; j++ {
	// 		seeds = append(seeds, j)
	// 	}
	// }

	toSoilFuncRules := parseSection2(sections[1])
	toFertFuncRules := parseSection2(sections[2])

	// toSoilFuncRules := parseSection2(sections[1])
	// toFertFuncRules := parseSection2(sections[2])
	// toWaterFuncRules := parseSection2(sections[3])
	// toLightFuncRules := parseSection2(sections[4])
	// toTempFuncRules := parseSection2(sections[5])
	// toHumFuncRules := parseSection2(sections[6])
	// toLocFuncRules := parseSection2(sections[7])

	fr1 := compose2Funcs(toSoilFuncRules, toFertFuncRules)
	f1 := funcRules2Func(fr1)
	// fmt.Println(f1(79))
	// fmt.Println(f1(92))
	// fmt.Println(f1(55))
	// fmt.Println(f1(67))
	fmt.Println(f1(50))
	fmt.Println(f1(51))
	fmt.Println(f1(52))
	fmt.Println(f1(53))
	fmt.Println(f1(54))
	fmt.Println(f1(55))
	// gigachadFuncRules := compose2Funcs(toLocFuncRules, compose2Funcs(toHumFuncRules, compose2Funcs(toTempFuncRules, compose2Funcs(toLightFuncRules, compose2Funcs(toWaterFuncRules, fr1)))))
	// gigachadFunc := funcRules2Func(gigachadFuncRules)

	// for i := 0; i < len(seedRanges); i += 2 {
	// 	// for j := seedRanges[i]; j < seedRanges[i]+seedRanges[i+1]; j++ {
	// 	// 	seeds = append(seeds, j)
	// 	// }
	// 	loc1 := gigachadFunc(seedRanges[i])
	// 	loc2 := gigachadFunc(seedRanges[i] + seedRanges[i+1] - 1)
	// 	gigachadMin := utils.Min(int64(loc1), int64(loc2))
	// 	fmt.Println(seedRanges[i], loc1, loc2, gigachadMin)
	// 	if gigachadMin < minLoc {
	// 		minLoc = gigachadMin
	// 	}
	// }

	for i := 0; i < len(seedRanges); i += 2 {
		loc1 := f1(seedRanges[i])
		loc2 := f1(seedRanges[i] + seedRanges[i+1] - 1)
		gigachadMin := utils.Min(loc1, loc2)
		fmt.Println(seedRanges[i], loc1, loc2, gigachadMin)
		if gigachadMin < minLoc {
			minLoc = gigachadMin
		}
	}

	fmt.Println("Min location:", minLoc)
}

// greedy worked out, in 445 s

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(1)

	// lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))
	inputBytes, _ := os.ReadFile(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))
	input := string(inputBytes)

	// strip last \n
	first(input[:len(input)-1])

	// strip last \n
	second(input[:len(input)-1])

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
