// https://adventofcode.com/2025/day/11

package main

import (
	"aoc/utils"
	"aoc/utils/graph"
	"fmt"
	"slices"
	"strings"
	"time"
)

const aocDay int = 11

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func First(lines []string) (strigifiedResult string) {
	dg := graph.NewDGraph[string]()

	for _, line := range lines {
		node, nbsStr := utils.SplitIn2(line, ": ")
		nbs := strings.Split(nbsStr, " ")
		dg.AddNode(node, nil)
		for _, n := range nbs {
			dg.AddEdgeUpsert(node, n)
		}
	}

	result := len(dg.GetAllPaths("you", "out"))
	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// Your puzzle answer was 643.

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

type pathCacheKey struct {
	From     string
	To       string
	CurDepth int
}

func getAllPaths2Inner(g *graph.DGraph[string], from, to string, curDepth, maxDepth int, cache map[pathCacheKey][][]string) (paths [][]string) {
	key := pathCacheKey{From: from, To: to, CurDepth: curDepth}
	if res, ok := cache[key]; ok {
		return res
	}

	if curDepth > maxDepth {
		cache[key] = [][]string{} // Cache empty slice
		return [][]string{}
	}

	if from == to {
		res := [][]string{{to}}
		cache[key] = res
		return res
	}

	var currentPaths [][]string
	for _, nb := range g.Neighbors[from] {
		// Recursively find paths from neighbor to target
		subPaths := getAllPaths2Inner(g, nb, to, curDepth+1, maxDepth, cache)
		for _, path := range subPaths {
			// Prepend 'nb' to each sub-path to form a complete path from 'from'
			currentPaths = append(currentPaths, slices.Concat([]string{nb}, path))
		}
	}

	cache[key] = currentPaths
	return currentPaths
}

func getAllPaths2(g *graph.DGraph[string], from, to string, maxDepth int) (paths [][]string) {
	// Initialize cache for this specific call to getAllPaths2
	cache := make(map[pathCacheKey][][]string)

	intermediatePaths := getAllPaths2Inner(g, from, to, 0, maxDepth, cache)

	var finalPaths [][]string
	for _, path := range intermediatePaths {
		// Only prepend 'from' if the path is not empty
		// (e.g., if 'from' is 'to' and path is just 'to')
		if len(path) > 0 && path[0] == from {
			finalPaths = append(finalPaths, path)
		} else {
			finalPaths = append(finalPaths, slices.Concat([]string{from}, path))
		}
	}

	return finalPaths
}

func Second(lines []string) (strigifiedResult string) {
	dg := graph.NewDGraph[string]()

	for _, line := range lines {
		node, nbsStr := utils.SplitIn2(line, ": ")
		nbs := strings.Split(nbsStr, " ")
		dg.AddNode(node, nil)
		for _, n := range nbs {
			dg.AddEdgeUpsert(node, n)
		}
	}

	// Notice that it's a tree and a state machine, arrows flow only in one direction

	result1 := len(getAllPaths2(dg, "svr", "fft", 11)) // 6351
	result2 := len(getAllPaths2(dg, "fft", "dac", 16)) // 6788852
	result3 := len(getAllPaths2(dg, "dac", "out", 10)) // 9676
	result := result1 * result2 * result3
	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// correct result 417190406827152 -> slow, ~19s

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
