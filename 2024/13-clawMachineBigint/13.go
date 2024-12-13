// https://adventofcode.com/2024/day/13

package main

import (
	"aoc/utils"
	"fmt"
	"math/big"
	"time"
)

const aocDay int = 13

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func First(lines []string) (strigifiedResult string) {
	result := 0.0

	var i, ax, ay, bx, by, px, py int

	for i < len(lines) {
		fmt.Sscanf(lines[i], "Button A: X+%d, Y+%d", &ax, &ay)
		fmt.Sscanf(lines[i+1], "Button B: X+%d, Y+%d", &bx, &by)
		fmt.Sscanf(lines[i+2], "Prize: X=%d, Y=%d", &px, &py)

		bPressCount := float64((ax*py - ay*px) / (by*ax - bx*ay))
		aPressCount := (float64(py) - float64(by)*bPressCount) / float64(ay)

		isAPressCountInteger := aPressCount == float64(int(aPressCount))
		isBPressCountInteger := bPressCount == float64(int(bPressCount))

		if isAPressCountInteger && isBPressCountInteger && aPressCount <= 100 && bPressCount <= 100 {
			result += aPressCount*3 + bPressCount
		}

		i += 4
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// correct: 31623

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func Second(lines []string) (strigifiedResult string) {
	result := new(big.Int)

	var i int
	var ax, ay, bx, by, px, py big.Int

	for i < len(lines) {
		fmt.Sscanf(lines[i], "Button A: X+%d, Y+%d", &ax, &ay)
		fmt.Sscanf(lines[i+1], "Button B: X+%d, Y+%d", &bx, &by)
		fmt.Sscanf(lines[i+2], "Prize: X=%d, Y=%d", &px, &py)

		px.Add(&px, big.NewInt(10000000000000))
		py.Add(&py, big.NewInt(10000000000000))

		axpy := new(big.Int).Mul(&ax, &py)
		aypx := new(big.Int).Mul(&ay, &px)
		axpy.Sub(axpy, aypx)

		byax := new(big.Int).Mul(&by, &ax)
		bxay := new(big.Int).Mul(&bx, &ay)
		byax.Sub(byax, bxay)

		bPressCount := new(big.Float).Quo(new(big.Float).SetInt(axpy), new(big.Float).SetInt(byax))
		aPressCount := new(big.Float).Quo(new(big.Float).Sub(new(big.Float).SetInt(&py), new(big.Float).Mul(new(big.Float).SetInt(&by), bPressCount)), new(big.Float).SetInt(&ay))

		isAPressCountInteger := aPressCount.IsInt()
		isBPressCountInteger := bPressCount.IsInt()

		if isAPressCountInteger && isBPressCountInteger {
			aPressCountInt, _ := aPressCount.Int64()
			bPressCountInt, _ := bPressCount.Int64()
			result.Add(result, new(big.Int).Mul(new(big.Int).SetInt64(aPressCountInt), new(big.Int).SetInt64(3)))
			result.Add(result, new(big.Int).SetInt64(bPressCountInt))
		}

		i += 4
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// too high 93863741768430 -> need bigint
// correct 93209116744825

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
