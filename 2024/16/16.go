// https://adventofcode.com/2024/day/16

package main

import (
	"aoc/utils"
	"aoc/utils/coordinate"
	"aoc/utils/direction"
	"aoc/utils/matrix"
	"fmt"
	"time"

	"github.com/rameshputalapattu/heapq"
)

const aocDay int = 16

type coord = coordinate.Coord
type mat = matrix.Matrix[string]
type dir = direction.Direction

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func calcRotationCost(facingDir, newDir direction.Direction) int {
	if facingDir == newDir {
		return 0
	}
	oneTurnNeighbors := []direction.Direction{coordinate.GetClockwise90DegreeDirection(facingDir), coordinate.GetCounterClockwise90DegreeDirection(facingDir)}
	if oneTurnNeighbors[0] == newDir || oneTurnNeighbors[1] == newDir {
		return 1000
	}
	return 2000
}

type node struct {
	coord coord
	dir   dir
	dist  int
}

func less(a, b node) bool {
	return a.dist < b.dist
}

func dijkstraFromStartToEndCoord(m mat, startCoord coord, endCoord coord) int {
	pq := heapq.NewPQWithItems([]node{}, less)
	pq.Push(node{coord: startCoord, dir: direction.Right, dist: 0})

	seen := coordinate.TravelMap{}

	for pq.Len() > 0 {
		item := pq.Pop()

		if item.coord == endCoord {
			return item.dist
		}

		if seen.ContainsCoordAndDir(item.coord, item.dir) {
			continue
		}

		seen.Add(item.coord, item.dir)

		offsets := []coordinate.DirOffset{coordinate.GetOffsetForDir(coordinate.GetClockwise90DegreeDirection(item.dir)), coordinate.GetOffsetForDir(item.dir), coordinate.GetOffsetForDir(coordinate.GetCounterClockwise90DegreeDirection(item.dir))}
		for _, offset := range offsets {
			newCoord := item.coord.GetNewCoord(offset.Dir)

			if !m.IsValidCoord(newCoord) || m.At(newCoord) == "#" {
				continue
			}

			rotationCost := calcRotationCost(item.dir, offset.Dir)
			dist := 1 + rotationCost

			pq.Push(node{coord: newCoord, dir: offset.Dir, dist: item.dist + dist})
		}
	}

	return -1
}

func First(lines []string) (strigifiedResult string) {
	m, startCoord := matrix.ParseStringMatrixAndGetStartingPoint(lines, func(value string) bool {
		return value == "S"
	})

	var endCoord coord
	m.ForEach(func(i, j int, value string) {
		if value == "E" {
			endCoord = coord{Row: i, Col: j}
		}
	})

	result := dijkstraFromStartToEndCoord(m, *startCoord, endCoord)

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// 85440 too high
// correct: 85432, e1: 7036, e2: 11048 -> needed to add dir to seen array

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func less2(a, b node2) bool {
	return a.dist < b.dist
}

type node2 struct {
	coord coord
	dir   dir
	dist  int
	path  []coord
}

func getCoordsOfAllBestPaths(m mat, startCoord coord, endCoord coord, shortestDist int) (resultCoords coordinate.CoordMap) {
	resultCoords = coordinate.CoordMap{}

	pq := heapq.NewPQWithItems([]node2{}, less2)
	pq.Push(node2{coord: startCoord, dir: direction.Right, dist: 0, path: []coord{startCoord}})

	seenCosts := coordinate.TravelDirCostMap{}

	for pq.Len() > 0 {
		item := pq.Pop()

		if item.coord == endCoord {
			// if the best paths are found already (they will be found first), stop the function
			if shortestDist != item.dist {
				return resultCoords
			}
			for _, coord := range item.path {
				resultCoords.Add(coord)
			}
		}

		if seenCosts.ContainsCoordDirCost(item.coord, item.dir) && seenCosts.Get(item.coord, item.dir) != item.dist {
			continue
		}

		seenCosts.Add(item.coord, item.dir, item.dist)

		offsets := []coordinate.DirOffset{coordinate.GetOffsetForDir(coordinate.GetClockwise90DegreeDirection(item.dir)), coordinate.GetOffsetForDir(item.dir), coordinate.GetOffsetForDir(coordinate.GetCounterClockwise90DegreeDirection(item.dir))}
		for _, offset := range offsets {
			newCoord := item.coord.GetNewCoord(offset.Dir)

			if !m.IsValidCoord(newCoord) || m.At(newCoord) == "#" {
				continue
			}

			rotationCost := calcRotationCost(item.dir, offset.Dir)
			dist := 1 + rotationCost

			pathToNewCoord := append(item.path, newCoord)

			pq.Push(node2{coord: newCoord, dir: offset.Dir, dist: item.dist + dist, path: utils.CloneArray(pathToNewCoord)})
		}
	}

	return resultCoords
}

func Second(lines []string) (strigifiedResult string) {
	m, startCoord := matrix.ParseStringMatrixAndGetStartingPoint(lines, func(value string) bool {
		return value == "S"
	})

	var endCoord coord
	m.ForEach(func(i, j int, value string) {
		if value == "E" {
			endCoord = coord{Row: i, Col: j}
		}
	})

	shortestDist := dijkstraFromStartToEndCoord(m, *startCoord, endCoord)
	coords := getCoordsOfAllBestPaths(m, *startCoord, endCoord, shortestDist)
	result := len(coords)

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// correct: 465, e1: 45, e2: 64 -> needed to add dir to seen array

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(2)

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
