// https://adventofcode.com/2023/day/12

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

//  ____             _    _____               _      ____            _     _
// | __ )  __ _  ___| | _|_   _| __ __ _  ___| | __ |  _ \ __ _ _ __| |_  / |
// |  _ \ / _` |/ __| |/ / | || '__/ _` |/ __| |/ / | |_) / _` | '__| __| | |
// | |_) | (_| | (__|   <  | || | | (_| | (__|   <  |  __/ (_| | |  | |_  | |
// |____/ \__,_|\___|_|\_\ |_||_|  \__,_|\___|_|\_\ |_|   \__,_|_|   \__| |_|

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

func first_backtrack(lines []string) {
	fmt.Println("--- First with backtrack ---")

	result := 0

	for _, line := range lines {
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
		result += possibilities
	}

	fmt.Println(result)
}

//  ____                          _
// |  _ \ ___  ___ _   _ _ __ ___(_)_   _____
// | |_) / _ \/ __| | | | '__/ __| \ \ / / _ \
// |  _ <  __/ (__| |_| | |  \__ \ |\ V /  __/
// |_| \_\___|\___|\__,_|_|  |___/_| \_/ \___|

var cache map[string]int = make(map[string]int)

func getCacheKey(cfg string, nums []int) string {
	return fmt.Sprintf("%s-%v", cfg, nums)
}

func count(cfg string, nums []int) int {
	if cfg == "" {
		if len(nums) == 0 {
			return 1
		}
		return 0
	}

	if len(nums) == 0 {
		if !strings.Contains(cfg, damagedChar) {
			return 1
		}
		return 0
	}

	// If solution of subtask present in cache, use it instead
	key := getCacheKey(cfg, nums)
	cachedResult, ok := cache[key]
	if ok {
		return cachedResult
	}

	result := 0

	// Consider the unknown spring ('?') as an operational spring ('.').
	if string(cfg[0]) == operationalChar || string(cfg[0]) == unknownChar {
		result += count(cfg[1:], nums)
	}

	// Consider the unknown spring ('?') as a damaged spring ('#').
	if string(cfg[0]) == damagedChar || string(cfg[0]) == unknownChar {
		// If there are enough springs to cover the current number and there
		// are no operational springs in that range.
		if nums[0] <= len(cfg) && !strings.Contains(cfg[:nums[0]], operationalChar) {
			// If the range is exactly at the end of the springs line.
			if len(cfg) == nums[0] {
				result += count("", nums[1:])
				// Otherwise see that the following spring to damaged interval is not a
				// damaged spring. An unknown spring ('?') is allowed, as it can be considered
				// an operational one ('.').
			} else if string(cfg[nums[0]]) != damagedChar {
				// Skip the section and take the next number.
				result += count(cfg[nums[0]+1:], nums[1:])
			}
		}
	}

	cache[key] = result
	return result
}

func first(lines []string) {
	fmt.Println("\n--- First ---")

	result := 0

	for _, line := range lines {
		hotSprings, damagedGroupsStr := utils.SplitIn2(line, " ")
		damagedGroupsSplitStr := strings.Split(damagedGroupsStr, ",")

		damagedGroups := []int{}
		for _, n := range damagedGroupsSplitStr {
			nn, _ := strconv.Atoi(n)
			damagedGroups = append(damagedGroups, nn)
		}
		result += count(hotSprings, damagedGroups)
	}

	fmt.Println(result)
}

// Your puzzle answer was 7173.

func second(lines []string) {
	fmt.Println("\n--- Second ---")

	result := 0

	for _, line := range lines {
		hotSprings, damagedGroupsStr := utils.SplitIn2(line, " ")

		extendedHotSprings := hotSprings + unknownChar + hotSprings + unknownChar + hotSprings + unknownChar + hotSprings + unknownChar + hotSprings

		damagedGroupsSplitStr := strings.Split(damagedGroupsStr, ",")
		extendedDamagedGroupsSplitStr := strings.Split(damagedGroupsStr+","+damagedGroupsStr+","+damagedGroupsStr+","+damagedGroupsStr+","+damagedGroupsStr, ",")

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

		// fmt.Println(extendedHotSprings, extendedDamagedGroups)

		c := count(extendedHotSprings, extendedDamagedGroups)
		// fmt.Println(c)
		// fmt.Println()
		result += c
	}

	fmt.Println(result)
}

// Your puzzle answer was 29826669191291.

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(2)

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	first_backtrack(lines)

	first(lines)

	second(lines)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
