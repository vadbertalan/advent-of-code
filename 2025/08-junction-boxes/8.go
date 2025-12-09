// https://adventofcode.com/2025/day/8

package main

import (
	"aoc/utils"
	"aoc/utils/coordinate"
	"aoc/utils/graph"
	"flag"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/rameshputalapattu/heapq"
)

const aocDay int = 8

type coord3 = coordinate.Coord3

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

type edge struct {
	a, b     coord3
	distance float64
}

func less(a, b edge) bool {
	return a.distance < b.distance
}

func parseCoord3IntoNodeString(c coord3) string {
	return fmt.Sprintf("%d,%d,%d", c.X, c.Y, c.Z)
}

func parseNodeStringIntoCoord3(s string) coord3 {
	splitStr := strings.Split(s, ",")
	return coord3{X: utils.Atoi(splitStr[0]), Y: utils.Atoi(splitStr[1]), Z: utils.Atoi(splitStr[2])}
}

func First(lines []string, edgeCount int) (strigifiedResult string) {
	result := 0

	pq := heapq.NewPQWithItems([]edge{}, less)
	udg := graph.NewUGraph[string]()

	for _, line := range lines {
		newCoord3 := parseNodeStringIntoCoord3(line)

		if len(udg.GetNodes()) == 0 {
			udg.AddNode(line, nil)
			continue
		}

		for _, existingNodeStr := range udg.GetNodes() {
			existingNodeCoord3 := parseNodeStringIntoCoord3(existingNodeStr)
			distance := newCoord3.EuclideanDist3(existingNodeCoord3)
			pq.Push(edge{a: newCoord3, b: existingNodeCoord3, distance: distance})
		}
		udg.AddNode(line, nil)
	}

	for i := 0; i < edgeCount; i++ {
		if pq.Len() == 0 {
			panic("Not enough edges in PQ")
		}
		popped := pq.Pop()
		udg.AddEdge(parseCoord3IntoNodeString(popped.a), parseCoord3IntoNodeString(popped.b))
	}

	_, components := udg.CountComponents(lines[0])

	componentLengths := []int{}
	for _, component := range components {
		componentLengths = append(componentLengths, len(component))
	}
	sort.Slice(componentLengths, func(i, j int) bool {
		return componentLengths[i] > componentLengths[j]
	})

	result = componentLengths[0] * componentLengths[1] * componentLengths[2]
	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// 0 too low -> blind try
// Your puzzle answer was 98696

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func Second(lines []string) (strigifiedResult string) {
	result := 0

	pq := heapq.NewPQWithItems([]edge{}, less)
	udg := graph.NewUGraph[string]()

	for _, line := range lines {
		newCoord3 := parseNodeStringIntoCoord3(line)

		if len(udg.GetNodes()) == 0 {
			udg.AddNode(line, nil)
			continue
		}

		for _, existingNodeStr := range udg.GetNodes() {
			existingNodeCoord3 := parseNodeStringIntoCoord3(existingNodeStr)
			distance := newCoord3.EuclideanDist3(existingNodeCoord3)
			pq.Push(edge{a: newCoord3, b: existingNodeCoord3, distance: distance})
		}
		udg.AddNode(line, nil)
	}

	// TODO use Kruskal's algorithm to build MST
	componentsCount, _ := udg.CountComponents(lines[0])
	lastXProduct := -1
	for componentsCount > 1 {
		if pq.Len() == 0 {
			panic("Not enough edges in PQ")
		}
		popped := pq.Pop()
		udg.AddEdge(parseCoord3IntoNodeString(popped.a), parseCoord3IntoNodeString(popped.b))
		lastXProduct = popped.a.X * popped.b.X
		componentsCount, _ = udg.CountComponents(lines[0])
	}

	result = lastXProduct
	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// Your puzzle answer was 2245203960

func main() {
	startTime := time.Now()

	edgeCountP := flag.Int("ec", 10, "Specify if you want to modify the edge count. Example edge count is 10")

	inputFileExtension := utils.GetInputFileExt(1)

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	fmt.Println("--- First ---")
	result := First(lines, *edgeCountP)
	fmt.Println(result)

	fmt.Println("\n--- Second ---")
	result = Second(lines)
	fmt.Println(result)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
