// https://adventofcode.com/2025/day/5

package main

import (
	"aoc/utils"
	"aoc/utils/collections"
	"fmt"
	"time"
)

const aocDay int = 5

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func First(lines []string) (strigifiedResult string) {
	result := 0

	ranges := [][2]int{}

	startOfAvailableIDs := 0
	for i, line := range lines {
		if line == "" {
			startOfAvailableIDs = i + 1
			break
		}
		startStr, endStr := utils.SplitIn2(line, "-")
		start := utils.Atoi(startStr)
		end := utils.Atoi(endStr)
		ranges = append(ranges, [2]int{start, end})
	}

	for i := startOfAvailableIDs; i < len(lines); i++ {
		id := utils.Atoi(lines[i])
		isFresh := false
	outerFor:
		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				isFresh = true
				break outerFor
			}
		}
		if isFresh {
			result++
		}
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// Your puzzle answer was 664.

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

type Range struct {
	start int
	end   int
}

func Second(lines []string) (strigifiedResult string) {
	result := 0

	mergedRanges := []*Range{}

	newRangesQ := collections.NewQueue[Range]()

	for _, line := range lines {
		if line == "" {
			break
		}
		startStr, endStr := utils.SplitIn2(line, "-")
		start := utils.Atoi(startStr)
		end := utils.Atoi(endStr)

		if len(mergedRanges) == 0 {
			mergedRanges = append(mergedRanges, &Range{start: start, end: end})
			continue
		}

		mergeHappened := false
		for i, mr := range mergedRanges {
			// Invalidated range check
			if mr == nil {
				continue
			}

			// No overlap case
			if end < mr.start || start > mr.end {
				continue
			}

			newRangesQ.Append(Range{start: min(start, mr.start), end: max(end, mr.end)})
			mergedRanges[i] = nil // Invalidate merged range for GC optimization
			mergeHappened = true
			break
		}

		if !mergeHappened {
			mergedRanges = append(mergedRanges, &Range{start: start, end: end})
		}
	}

	for !newRangesQ.IsEmpty() {
		newRange := newRangesQ.Pop()

		mergeHappened := false
		for i, mr := range mergedRanges {
			// Invalidated range check
			if mr == nil {
				continue
			}

			// No overlap case
			if newRange.end < mr.start || newRange.start > mr.end {
				continue
			}

			newRangesQ.Append(Range{start: min(newRange.start, mr.start), end: max(newRange.end, mr.end)})
			mergedRanges[i] = nil // Invalidate merged range for GC optimization
			mergeHappened = true
			break
		}

		if !mergeHappened {
			mergedRanges = append(mergedRanges, &newRange)
		}
	}

	for _, r := range mergedRanges {
		if r == nil {
			continue
		}
		result += (r.end - r.start + 1)
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// Your puzzle answer was 350780324308385.

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
