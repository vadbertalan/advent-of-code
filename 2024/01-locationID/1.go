// https://adventofcode.com/2024/day/1

package main

import (
	"aoc/utils-go"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

const aocDay int = 1

func First(lines []string) string {
	fmt.Println("--- First ---")

	var list1, list2 []int

	for _, line := range lines {
		split := strings.Split(line, "   ")

		nr1, _ := strconv.Atoi(split[0])
		nr2, _ := strconv.Atoi(split[1])

		list1 = append(list1, nr1)
		list2 = append(list2, nr2)
	}

	sortedList1 := append([]int(nil), list1...)
	sort.Ints(sortedList1)

	sortedList2 := append([]int(nil), list2...)
	sort.Ints(sortedList2)

	result := 0
	for i := 0; i < len(sortedList1); i++ {
		result += int(math.Abs(float64(sortedList1[i] - sortedList2[i])))
	}

	return fmt.Sprint(result)
}

// correct 2166959, example 11

func Second(lines []string) string {
	fmt.Println("\n--- Second ---")

	var list1 []int
	map2 := make(map[int]int)

	for _, line := range lines {
		split := strings.Split(line, "   ")

		nr1, _ := strconv.Atoi(split[0])
		nr2, _ := strconv.Atoi(split[1])

		list1 = append(list1, nr1)
		map2[nr2]++
	}

	result := 0
	for _, nr := range list1 {
		if count, ok := map2[nr]; ok {
			result += nr * count
		}
	}

	return fmt.Sprint(result)
}

// correct 23741109, example 31

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(1)
	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	fmt.Println(First(lines))
	fmt.Println(Second(lines))

	fmt.Printf("\n✨ Finished in %.3f seconds\n", time.Since(startTime).Seconds())
}
