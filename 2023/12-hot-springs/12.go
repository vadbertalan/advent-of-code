package main

import (
	"aoc/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const aocDay int = 12

const operationalChar = "."
const damagedChar = "#"
const unknownChar = "?"

type hotSpring struct {
	row   string
	order []int
}

func isValid(row string, counts []int) bool {
	if strings.Count(row, damagedChar) != utils.Sum(counts) {
		return false
	}

	if strings.Count(row, unknownChar) > 0 {
		return false
	}

	pattern := `\.*`

	for i, count := range counts {
		pattern += fmt.Sprintf(`#{%d}`, count)
		if i != len(counts)-1 {
			pattern += `\.+`
		}
	}

	compiledPattern := regexp.MustCompile(pattern)

	return compiledPattern.MatchString(row)
}

var possibilities int = 0
var str string = ""
var counts []int
var indices []int

func getIndices(in string, of string) (ret []int) {
	for i := 0; i < len(in); i++ {
		if string(in[i]) == of {
			ret = append(ret, i)
		}
	}
	return ret
}

var input = []string{operationalChar, damagedChar}

func backtracking(level int) {

	for _, c := range input {
		ind := indices[level]
		str = str[:ind] + strings.Replace(str[ind:], string(str[ind]), c, 1)

		if isValid(str, counts) {
			possibilities++
		} else if level < len(indices)-1 {
			backtracking(level + 1)
		}
	}
}

func first(lines []string) {
	fmt.Println("--- First ---")

	result := 0

	for _, line := range lines {
		fmt.Println(line)

		hotSprings, damagedGroupsStr := utils.SplitIn2(line, " ")
		damagedGroupsSplitStr := strings.Split(damagedGroupsStr, ",")

		damagedGroups := []int{}
		for _, n := range damagedGroupsSplitStr {
			nn, _ := strconv.Atoi(n)
			damagedGroups = append(damagedGroups, nn)
		}

		str = hotSprings
		counts = damagedGroups
		possibilities = 0
		indices = getIndices(str, unknownChar)

		if len(indices) > 0 {
			backtracking(0)
		}
		// fmt.Println(possibilities)
		result += possibilities

		// if k > 3 {
		// 	break
		// }
	}

	fmt.Println(result)
}

func second(lines []string) {
	fmt.Println("\n--- Second ---")

	result := 0

	for i, line := range lines {
		fmt.Println(i, line)

		hotSprings, damagedGroupsStr := utils.SplitIn2(line, " ")

		extendedHotSprings := hotSprings + unknownChar + hotSprings
		// if hotSprings[len(hotSprings) - 1] ==
		// extendedDamagedGroupsStr := damagedGroupsStr

		// for i := 0; i < 4; i++ {
		// extendedHotSprings += unknownChar + hotSprings
		// extendedDamagedGroupsStr += "," + damagedGroupsStr
		// }

		damagedGroupsSplitStr := strings.Split(damagedGroupsStr, ",")
		extendedDamagedGroupsSplitStr := strings.Split(damagedGroupsStr+","+damagedGroupsStr, ",")

		damagedGroups := []int{}
		for _, n := range damagedGroupsSplitStr {
			nn, _ := strconv.Atoi(n)
			damagedGroups = append(damagedGroups, nn)
		}

		extendedDamagedGroups := []int{}
		for _, n := range extendedDamagedGroupsSplitStr {
			nn, _ := strconv.Atoi(n)
			extendedDamagedGroups = append(extendedDamagedGroups, nn)
		}

		fmt.Println(extendedHotSprings, extendedDamagedGroups)

		str = hotSprings
		counts = damagedGroups
		possibilities = 0
		indices = getIndices(str, unknownChar)

		if len(indices) > 0 {
			backtracking(0)
		}

		first := possibilities
		fmt.Println("first", first)

		str = extendedHotSprings
		counts = extendedDamagedGroups
		possibilities = 0
		indices = getIndices(str, unknownChar)

		if len(indices) > 0 {
			backtracking(0)
		}

		second := possibilities

		x := second / first
		result += first * x * x * x * x

		fmt.Println("second", second)
		fmt.Println("---")
	}

	// fmt.Println(isValid(".##...###..", []int{1, 1, 3}))

	fmt.Println(result)
}

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt()

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	// first(lines)

	second(lines)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
