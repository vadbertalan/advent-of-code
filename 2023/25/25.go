// https://adventofcode.com/2023/day/25

package main

import (
	"aoc/utils"
	"aoc/utils/graph"
	"fmt"
	"slices"
	"strings"
	"time"
)

const aocDay int = 25

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func getMostUsedEdge(g graph.UGraph) (edgeKey string) {
	// Count init egde frequencies counter map
	edges := g.GetEdges()
	edgeFreqMap := map[string]int{}
	for _, edge := range edges {
		edgeFreqMap[graph.GetEdgeKey(edge[0], edge[1])] = 0
	}

	// TraverseBFS starting from each node and count visited edges
	for node := range g.Neighbors {
		_, edges := g.TraverseBFS(node)

		for _, edgeStr := range edges {
			edgeFreqMap[edgeStr]++
		}
	}

	type edgeFreqPair struct {
		node1, node2 string
		edgeKey      string
		freq         int
	}
	edgeFreqPairs := []edgeFreqPair{}
	for edgeKey, freq := range edgeFreqMap {
		node1, node2 := graph.ParseEdgeKey(edgeKey)
		edgeFreqPairs = append(edgeFreqPairs, edgeFreqPair{node1, node2, edgeKey, freq})
	}
	slices.SortFunc(edgeFreqPairs, func(ef1, ef2 edgeFreqPair) int {
		if ef1.freq == ef2.freq {
			return 0
		}
		if ef1.freq < ef2.freq {
			return 1
		}
		return -1
	})

	return edgeFreqPairs[0].edgeKey
}

func First(lines []string) (strigifiedResult string) {
	fmt.Println("--- First ---")

	result := 0

	g := graph.NewUGraph()

	var firstNode string

	// Build graph
	for i, line := range lines {
		node1, nodesStr := utils.SplitIn2(line, ":")
		if i == 0 {
			firstNode = node1
		}
		neighbors := strings.Fields(nodesStr)
		for _, neighbor := range neighbors {
			g.AddNode(node1)
			g.AddNode(neighbor)
			g.AddEdge(node1, neighbor)
		}
	}

	mostUsedEdge := getMostUsedEdge(*g)
	g.RemoveEdge(mostUsedEdge)

	secondMostUsedEdge := getMostUsedEdge(*g)
	g.RemoveEdge(secondMostUsedEdge)

	thirdMostUsedEdge := getMostUsedEdge(*g)
	g.RemoveEdge(thirdMostUsedEdge)

	path, _ := g.TraverseBFS(firstNode)
	result = (len(g.Neighbors) - len(path)) * len(path)

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// 5812 too low
// 527790 good

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt()

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	result := First(lines)
	fmt.Println(result)

	// No part 2, as this was the last exercise this year

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
