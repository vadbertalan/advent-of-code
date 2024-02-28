// https://adventofcode.com/2023/day/22

package main

import (
	"aoc/utils"
	"aoc/utils/coordinate"
	"aoc/utils/matrix"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const aocDay int = 22

type coord = coordinate.Coord
type mat = matrix.Matrix[string]

type crd2 struct {
	x, y int
}

type crd3 struct {
	x, y, z int
}

func calcIntersectionCoord3(P11, P12, P21, P22 crd3) *crd3 {
	// First line equation
	a1 := P12.x - P11.x
	b1 := P12.y - P11.y
	c1 := P12.z - P11.z

	// Second line equation
	a2 := P22.x - P21.x
	b2 := P22.y - P21.y
	// c2 := P22.z - P21.z

	// Calc s and t based on first line equation:
	s1 := (P21.y - P11.y - ((P21.x-P11.x)/a1)*b1) / ((a2*b1)/a1 - b2)
	// s2 := (P22.z - P11.z - ((P22.x-P11.x)/a1)*c1) / ((a2*c1)/a1 - c2)

	t1 := (P21.x-P11.x)/a1 + (s1*a2)/a1

	x := P11.x + t1*a1
	y := P11.y + t1*b1
	z := P11.z + t1*c1

	return &crd3{x, y, z}
}

func findIntersection(p1, p2, q1, q2 crd2) (crd2, error) {
	a1 := p2.y - p1.y
	b1 := p1.x - p2.x
	c1 := a1*p1.x - b1*p1.y

	a2 := q2.y - q1.y
	b2 := q1.x - q2.x
	c2 := a2*q1.x - b2*q1.y

	// Use Cramer system to solve equation system

	determinant := a1*b2 - a2*b1

	if determinant == 0 {
		return crd2{}, fmt.Errorf("Lines are skew, no unique intersection point")
	}

	x := (c1*b2 - c2*b1) / determinant
	y := (a1*c2 - a2*c1) / determinant

	return crd2{x, y}, nil
}

type brick = [2]crd3

func parseBricks(lines []string) []brick {
	bricks := []brick{}
	for _, line := range lines {
		c1str, c2str := utils.SplitIn2(line, "~")
		c1 := strings.Split(c1str, ",")
		c2 := strings.Split(c2str, ",")

		x1, _ := strconv.Atoi(c1[0])
		y1, _ := strconv.Atoi(c1[1])
		z1, _ := strconv.Atoi(c1[2])

		x2, _ := strconv.Atoi(c2[0])
		y2, _ := strconv.Atoi(c2[1])
		z2, _ := strconv.Atoi(c2[2])

		bricks = append(bricks, brick{crd3{x1, y1, z1}, crd3{x2, y2, z2}})
	}
	return bricks
}

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func First(lines []string) (strigifiedResult string) {
	fmt.Println("--- First ---")

	bricks := parseBricks(lines)

	for _, br := range bricks {
		bricksBelow := utils.Filter[brick](bricks, func(otherBr [2]crd3) bool {
			return otherBr
		})
	}

	fmt.Println(findIntersection(crd2{1, 0}, crd2{1, 2}, crd2{0, 0}, crd2{2, 0}))

	result := 0
	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func Second(lines []string) (strigifiedResult string) {
	fmt.Println("\n--- Second ---")

	result := 0

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(1)

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	result := First(lines)
	fmt.Println(result)

	result = Second(lines)
	fmt.Println(result)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
