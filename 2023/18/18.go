// https://adventofcode.com/2023/day/18

package main

import (
	"aoc/utils"
	"aoc/utils/coordinate"
	"aoc/utils/direction"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const aocDay int = 18

type coord = coordinate.Coord
type dir = direction.Direction

type offset struct {
	dir       direction.Direction
	rowOffset int
	colOffset int
}

func getOffsetsArray() []offset {
	return []offset{
		{direction.Up, -1, 0},
		{direction.Right, 0, 1},
		{direction.Down, 1, 0},
		{direction.Left, 0, -1},
	}
}

type coordMap map[string]bool

func (cm coordMap) contains(c coord) bool {
	_, ok := cm[fmt.Sprintf("%d-%d", c.Row, c.Col)]
	return ok
}

func (cm *coordMap) add(c coord) {
	(*cm)[fmt.Sprintf("%d-%d", c.Row, c.Col)] = true
}

func buildCoordMap(coords []coord) coordMap {
	cm := &coordMap{}
	for _, c := range coords {
		cm.add(c)
	}
	return *cm
}

var cm, seen *coordMap

func trav(c coord) (count int) {
	seen.add(c)

	count = 1

	offsets := getOffsetsArray()
	neighborCount := len(offsets)
	for i := 0; i < neighborCount; i++ {
		xx := c.Row + offsets[i].rowOffset
		yy := c.Col + offsets[i].colOffset
		nc := coord{Row: xx, Col: yy}
		if !cm.contains(nc) && !seen.contains(nc) {
			count += trav(nc)
		}
	}

	return count
}

func first(lines []string) {
	fmt.Println("--- First ---")

	result := 0

	cur := coord{Row: 0, Col: 0}
	coords := []coord{cur}

	n, m := 1, 1
	nd, md := 1, 1

	for _, line := range lines {
		fields := strings.Fields(line)
		d := fields[0]
		c, _ := strconv.Atoi(fields[1])

		dr, dc := 0, 0
		switch d {
		case "R":
			{
				dr, dc = 0, 1
				md += c
				if md > m {
					m = md
				}
			}
		case "D":
			{
				dr, dc = 1, 0
				nd += c
				if nd > n {
					n = nd
				}
			}
		case "L":
			{
				dr, dc = 0, -1
				md -= c
			}
		case "U":
			{
				dr, dc = -1, 0
				nd -= c
			}
		}

		for i := 0; i < c; i++ {
			cur.Row += dr
			cur.Col += dc
			coords = append(coords, cur)
		}

	}

	result += len(coords) - 1

	cm2 := buildCoordMap(coords)
	cm = &cm2

	seen = &coordMap{}

	// You may need to change these coords not to to get stack overflow (for ex in case of .exin2 input)
	inside := trav(coord{Row: 1, Col: 1})

	fmt.Println(result + inside)
}

func second(lines []string) {
	fmt.Println("\n--- Second ---")

	result := 0

	cur := coord{Row: 0, Col: 0}
	coords := []coord{cur}

	dmap := map[string]string{
		"0": "R",
		"1": "D",
		"2": "L",
		"3": "U",
	}

	// boundaries
	b := 0

	for _, line := range lines {
		fields := strings.Fields(line)
		hex := fields[2]

		c, _ := strconv.ParseInt(hex[2:7], 16, 64)
		d := dmap[string(hex[7])]

		dr, dc := 0, 0
		switch d {
		case "R":
			{
				dr, dc = 0, 1
			}
		case "D":
			{
				dr, dc = 1, 0
			}
		case "L":
			{
				dr, dc = 0, -1
			}
		case "U":
			{
				dr, dc = -1, 0
			}
		}

		for i := 0; i < int(c); i++ {
			cur.Row += dr
			cur.Col += dc
		}
		coords = append(coords, cur)

		// count number of points in the wall in total
		b += int(c)
	}

	// dropping last coord as it is the same as the first one
	coords = coords[:len(coords)-1]

	// Calculating area with Shoelace formulae https://en.wikipedia.org/wiki/Shoelace_formula
	// See `Other formulas` section. Absolute value is needed because of the order of the vertices.
	A := int(utils.CalcAreaShoelace(coords))

	// Pick's theorem https://en.wikipedia.org/wiki/Pick%27s_theorem
	i := A - b/2 + 1

	// re-adding boundary count to the inner #'s
	result = i + b
	fmt.Println(result)
}

func main() {
	// Can't be set to more than 2000000000 bytes
	// debug.SetMaxStack(2000000000)

	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(1)

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	first(lines)

	second(lines)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
