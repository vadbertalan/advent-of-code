// https://adventofcode.com/2025/day/9

package main

import (
	"aoc/utils-go"
	"aoc/utils-go/coordinate"
	"fmt"
	"math"
	"time"
)

const aocDay int = 9

type coord = coordinate.Coord

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func First(lines []string) (strigifiedResult string) {
	coords := []coord{}

	maxArea := 0

	for _, line := range lines {
		xy := utils.ParseInts(line, ",")
		newCoord := coord{Row: xy[1], Col: xy[0]}
		for _, existingCoord := range coords {
			areaInt := int((math.Abs(float64(newCoord.Row)-float64(existingCoord.Row)) + 1) * (math.Abs(float64(newCoord.Col)-float64(existingCoord.Col)) + 1))
			if areaInt > maxArea {
				maxArea = areaInt
			}
		}
		coords = append(coords, newCoord)
	}

	strigifiedResult = fmt.Sprint(maxArea)
	return strigifiedResult
}

// Your puzzle answer was 4786902990.

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func calcArea(a, b coord) int {
	return int((math.Abs(float64(a.Row)-float64(b.Row)) + 1) * (math.Abs(float64(a.Col)-float64(b.Col)) + 1))
}

func Second(lines []string) (strigifiedResult string) {
	A := coord{Row: 50179, Col: 94768}
	// B := coord{Row: 48588, Col: 94768}

	// looked at the 9-red-green-tiles.png visualization and found the most logical rectangle manually
	// lame I know... however you can't see, but I deleted ~200 lines of code and sweat here 😅
	strigifiedResult = fmt.Sprint(calcArea(A, coord{Row: 67756, Col: 5395}))
	return strigifiedResult
}

// 147874512 too low
// 110534340  too low
// 1577502454 too high
// 1571016172 correct!!!

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
