// https://adventofcode.com/2024/day/21

package main

import (
	"aoc/utils-go"
	"aoc/utils-go/coordinate"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/mowshon/iterium"
)

const aocDay int = 21

type coord = coordinate.Coord

var numpadMap = map[string]coord{
	"1": {Row: 2, Col: 0}, "2": {Row: 2, Col: 1}, "3": {Row: 2, Col: 2}, "4": {Row: 1, Col: 0}, "5": {Row: 1, Col: 1}, "6": {Row: 1, Col: 2}, "7": {Row: 0, Col: 0}, "8": {Row: 0, Col: 1}, "9": {Row: 0, Col: 2}, "0": {Row: 3, Col: 1}, "A": {Row: 3, Col: 2},
}

var dirpadMap = map[string]coord{
	"^": {Row: 0, Col: 1}, "v": {Row: 1, Col: 1}, "<": {Row: 1, Col: 0}, ">": {Row: 1, Col: 2}, "A": {Row: 0, Col: 2},
}

func getAllThePaths(sourceCoord, destinationCoord, gapCoord coord) []string {
	drow := destinationCoord.Row - sourceCoord.Row
	dcol := destinationCoord.Col - sourceCoord.Col

	path := ""
	if drow > 0 {
		for i := 0; i < drow; i++ {
			path += "v"
		}
	} else if drow < 0 {
		for i := 0; i < utils.Abs(drow); i++ {
			path += "^"
		}
	}

	if dcol > 0 {
		for i := 0; i < dcol; i++ {
			path += ">"
		}
	} else if dcol < 0 {
		for i := 0; i < utils.Abs(dcol); i++ {
			path += "<"
		}
	}

	permutations := iterium.Permutations[string](strings.Split(path, ""), len(path))
	allThePaths, err := permutations.Slice()
	if err != nil {
		panic(err)
	}

	filteredPaths := utils.FilterDuplicates(utils.Map(allThePaths, func(path []string) string {
		return strings.Join(path, "")
	}))

	returnablePaths := []string{}
	for _, path := range filteredPaths {
		itCoord := sourceCoord
		pathCorrect := true
		for _, dir := range path {
			switch dir {
			case 'v':
				itCoord.Row++
			case '^':
				itCoord.Row--
			case '>':
				itCoord.Col++
			case '<':
				itCoord.Col--
			}
			if itCoord.IsEqual(gapCoord) {
				pathCorrect = false
			}
		}
		if pathCorrect {
			returnablePaths = append(returnablePaths, path+"A")
		}
	}

	return utils.FilterDuplicates(returnablePaths)
}

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func getMinSequence(currentPad, targetPad string, depth int, padMap map[string]coord, gapCoord coord) string {
	currentCoord := padMap[currentPad]
	targetCoord := padMap[targetPad]

	paths := getAllThePaths(currentCoord, targetCoord, gapCoord)

	minPath := ""
	minLen := math.MaxInt64
	if depth > 0 {
		for _, path := range paths {
			subseq := getNextSequences(path, "A", depth)
			if len(subseq) < minLen {
				minPath = subseq
				minLen = len(subseq)
			}
		}
	} else {
		for _, path := range paths {
			if len(path) < minLen {
				minPath = path
				minLen = len(path)
			}
		}
	}

	return minPath
}

func getNextSequences(dirpadInput string, currentPad string, depth int) string {
	sequence := ""
	for _, dir := range dirpadInput {
		sequence += getMinSequence(currentPad, string(dir), depth-1, dirpadMap, coord{Row: 0, Col: 0})
		currentPad = string(dir)
	}
	return sequence
}

func First(lines []string) (strigifiedResult string) {
	result := 0

	for _, line := range lines {
		s3 := ""
		prevNum := "A"
		for _, nrChar := range line {
			currentNum := string(nrChar)

			s3 += getMinSequence(prevNum, currentNum, 2, numpadMap, coord{Row: 3, Col: 0})
			prevNum = currentNum
		}

		numPartOrNrStr, err := strconv.Atoi(line[:len(line)-1])
		if err != nil {
			panic(err)
		}

		result += (len(s3) * numPartOrNrStr)
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// 219108 too high
// 217184 too high
// 215216 too high
// correct: 212488, example 126384

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

var cache = map[string]int{}

func getCacheKey(currentPad, targetPad string, depth int) string {
	return currentPad + "," + targetPad + "," + strconv.Itoa(depth)
}

func getMinSequence2(currentPad, targetPad string, depth int, padMap map[string]coord, gapCoord coord) int {
	if value, ok := cache[getCacheKey(currentPad, targetPad, depth)]; ok {
		return value
	}

	currentCoord := padMap[currentPad]
	targetCoord := padMap[targetPad]

	paths := getAllThePaths(currentCoord, targetCoord, gapCoord)

	minLen := math.MaxInt64
	if depth > 0 {
		for _, path := range paths {
			subseqLen := getNextSequences2(path, "A", depth)
			if subseqLen < minLen {
				minLen = subseqLen
			}
		}
	} else {
		for _, path := range paths {
			if len(path) < minLen {
				minLen = len(path)
			}
		}
	}

	cache[getCacheKey(currentPad, targetPad, depth)] = minLen

	return minLen
}

func getNextSequences2(dirpadInput string, currentPad string, depth int) int {
	sequenceLen := 0
	for _, dir := range dirpadInput {
		sequenceLen += getMinSequence2(currentPad, string(dir), depth-1, dirpadMap, coord{Row: 0, Col: 0})
		currentPad = string(dir)
	}
	return sequenceLen
}

func Second(lines []string) (strigifiedResult string) {
	result := 0

	for _, line := range lines {
		s3len := 0
		prevNum := "A"
		for _, nrChar := range line {
			currentNum := string(nrChar)

			s3len += getMinSequence2(prevNum, currentNum, 25, numpadMap, coord{Row: 3, Col: 0})
			prevNum = currentNum
		}

		numPartOrNrStr, err := strconv.Atoi(line[:len(line)-1])
		if err != nil {
			panic(err)
		}

		result += (s3len * numPartOrNrStr)
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// correct: 258263972600402, no example give

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
