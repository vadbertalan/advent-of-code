// https://adventofcode.com/2024/day/11

package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"time"
)

const aocDay int = 11

var dp = make(map[int]map[int]int)

func transform_count(nr int, blinksLeft int) int {
	if _, ok := dp[nr]; ok {
		if val, ok := dp[nr][blinksLeft]; ok {
			return val
		}
	} else {
		dp[nr] = make(map[int]int)
	}

	if blinksLeft == 0 {
		return 1
	}

	nrStr := strconv.Itoa(nr)

	if nr == 0 {
		return transform_count(1, blinksLeft-1)
	} else if len(nrStr)%2 == 0 {
		nr1str := nrStr[:len(nrStr)/2]
		nr2str := nrStr[len(nrStr)/2:]
		nr1, _ := strconv.Atoi(nr1str)
		nr2, _ := strconv.Atoi(nr2str)
		nr1result := transform_count(nr1, blinksLeft-1)
		nr2result := transform_count(nr2, blinksLeft-1)
		dp[nr1][blinksLeft-1] = nr1result
		dp[nr2][blinksLeft-1] = nr2result
		dp[nr][blinksLeft] = nr1result + nr2result
		return nr1result + nr2result
	} else {
		return transform_count(nr*2024, blinksLeft-1)
	}
}

func solve(lines []string, blinksLeft int) (strigifiedResult string) {
	result := 0

	ints := utils.ParseInts(lines[0], " ")

	for _, nr := range ints {
		result += transform_count(nr, blinksLeft)
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func First(lines []string) (strigifiedResult string) {
	return solve(lines, 25)
}

// correct 198075, example 55312

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func Second(lines []string) (strigifiedResult string) {
	return solve(lines, 75)
}

// 65601038650482 not correct - wrong input, lame, 1min cooldown, nice
// 235571309320764 correct, example, no result provided on site, but I trust my program output: 65601038650482

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
