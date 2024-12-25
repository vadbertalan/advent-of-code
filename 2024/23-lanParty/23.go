// https://adventofcode.com/2024/day/23

package main

import (
	"aoc/utils"
	"aoc/utils/collections"
	"aoc/utils/graph"
	"fmt"
	"sort"
	"strings"
	"time"
)

const aocDay int = 23

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func isCycle(g *graph.UGraph[any], n1, n2, n3 string) bool {
	return g.HasEdge(n1, n2) && g.HasEdge(n2, n3) && g.HasEdge(n3, n1)
}

func First(lines []string) (strigifiedResult string) {
	result := 0

	g := graph.NewUGraph[any]()

	for _, line := range lines {
		a, b := utils.SplitIn2(line, "-")
		g.AddNode(a, nil)
		g.AddNode(b, nil)
		g.AddEdge(a, b)
	}

	nodes := g.GetNodes()
	for i := 0; i < len(nodes)-2; i++ {
		for j := i + 1; j < len(nodes)-1; j++ {
			for k := j + 1; k < len(nodes); k++ {
				isAnyStartingWithT := nodes[i][0] == 't' || nodes[j][0] == 't' || nodes[k][0] == 't'
				if isAnyStartingWithT && isCycle(g, nodes[i], nodes[j], nodes[k]) {
					result++
				}
			}
		}
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// correct: 1149, example 7

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func stringifySortedCircle(circle []string) string {
	ret := ""
	sort.Strings(circle)
	for _, node := range circle {
		ret += node + ","
	}
	return ret[:len(ret)-1]
}

func normalize(bigCircles [][]string) [][]string {
	set := collections.NewSet[string]()
	for _, circle := range bigCircles {
		set.Add(stringifySortedCircle(circle))
	}
	ret := [][]string{}
	for _, circleStr := range set.GetValues() {
		ret = append(ret, strings.Split(circleStr, ","))
	}
	return ret
}

func Second(lines []string) (strigifiedResult string) {
	result := ""

	g := graph.NewUGraph[any]()

	for _, line := range lines {
		a, b := utils.SplitIn2(line, "-")
		g.AddNode(a, nil)
		g.AddNode(b, nil)
		g.AddEdge(a, b)
	}

	circles := [][3]string{}

	nodes := g.GetNodes()
	for i := 0; i < len(nodes)-2; i++ {
		for j := i + 1; j < len(nodes)-1; j++ {
			for k := j + 1; k < len(nodes); k++ {
				if isCycle(g, nodes[i], nodes[j], nodes[k]) {
					circles = append(circles, [3]string{nodes[i], nodes[j], nodes[k]})
				}
			}
		}
	}

	bigCircles := [][]string{}

	for _, circle := range circles {
		bigCircles = append(bigCircles, []string{circle[0], circle[1], circle[2]})
	}

	for len(bigCircles) > 0 {
		newBigCircles := [][]string{}
		for _, existingBigCircle := range bigCircles {
			for _, node := range g.Neighbors[existingBigCircle[0]] {
				if utils.Contains(existingBigCircle, node) {
					continue
				}
				if utils.Every(existingBigCircle, func(member string) bool {
					return g.HasEdge(node, member)
				}) {
					newBigCircles = append(newBigCircles, append(existingBigCircle, node))
				}
			}
		}

		if len(newBigCircles) > 0 {
			result = stringifySortedCircle(newBigCircles[0])
		}

		bigCircles = normalize(newBigCircles)
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// incorrect: 13
// correct: as,co,do,kh,km,mc,np,nt,un,uq,wc,wz,yo. example: co,de,ka,ta

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
