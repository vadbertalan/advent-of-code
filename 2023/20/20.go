// https://adventofcode.com/2023/day/20

package main

import (
	"aoc/utils"
	"aoc/utils/collections"
	"aoc/utils/graph"
	"fmt"
	"strings"
	"time"
)

const aocDay int = 20

const broadcasterNode = "broadcaster"

type moduleType = int

const (
	flipflop moduleType = iota
	conjunction
	broadcaster
	output
)

type pulseType = string

const (
	low  pulseType = "low"
	high           = "high"
)

func parseLine(line string) (fromNode string, modType moduleType, toNodes []string) {
	fromNode, toStr := utils.SplitIn2(line, " -> ")

	for _, toNode := range strings.Split(toStr, ", ") {
		toNodes = append(toNodes, toNode)
	}

	if fromNode == "broadcaster" {
		modType = broadcaster
	} else if fromNode[0] == '%' {
		modType = flipflop
		fromNode = fromNode[1:]
	} else if fromNode[0] == '&' {
		modType = conjunction
		fromNode = fromNode[1:]
	} else {
		modType = output
	}

	return fromNode, modType, toNodes
}

func parseLines(lines []string) (dgraph *graph.DGraph[interface{}], ffStates map[string]bool, conjStates map[string]map[string]pulseType) {
	dgraph = graph.NewDGraph[interface{}]()

	ffStates = map[string]bool{}
	conjStates = map[string]map[string]pulseType{}

	conjNodes := []string{}

	for _, line := range lines {
		fromNode, modType, toNodes := parseLine(line)

		if modType == flipflop {
			ffStates[fromNode] = false
		} else if modType == conjunction {
			conjStates[fromNode] = map[string]pulseType{}
		}

		dgraph.AddNode(fromNode, nil)
		for _, toNode := range toNodes {
			dgraph.AddNode(toNode, nil)
			dgraph.AddEdge(fromNode, toNode)
		}

		if modType == conjunction {
			conjNodes = append(conjNodes, fromNode)
		}
	}

	// Get inc nodes of conj states to init inc states
	for _, conjNode := range conjNodes {
		for _, incNode := range dgraph.GetIncomingNodesOf(conjNode) {
			conjStates[conjNode][incNode] = low
		}
	}

	return dgraph, ffStates, conjStates
}

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

type pulse struct {
	from     string
	to       string
	pulseTyp pulseType
}

func First(lines []string) (strigifiedResult string) {
	fmt.Println("--- First ---")

	// Parse graph and init data structures
	dgraph, ffStates, conjStates := parseLines(lines)

	// Solve
	pulseCounts := map[pulseType]int{
		low:  0,
		high: 0,
	}

	const buttonPressCount = 1000

	i := buttonPressCount
	for i > 0 {
		// fmt.Println(fmt.Sprintf("\n---- Button press %d ---\n", buttonPressCount-i+1))

		// button -low-> broadcaster counts too
		pulseCounts[low]++

		q := collections.NewQueue[pulse]()
		for _, broadcasterNeighbor := range dgraph.Neighbors[broadcasterNode] {
			q.Append(pulse{from: broadcasterNode, to: broadcasterNeighbor, pulseTyp: low})
		}

		for !q.IsEmpty() {

			p := q.Pop()

			// fmt.Printf("%s -%s-> %s\n", p.from, p.pulseTyp, p.to)

			pulseCounts[p.pulseTyp]++

			// Transmit correct pulse
			var pulseTypeToSend pulseType

			// Determine pulse type to send
			wasTurnedOn, isFlipFlop := ffStates[p.to]
			if isFlipFlop {
				if p.pulseTyp == high {
					continue
				}
				ffStates[p.to] = !wasTurnedOn

				if wasTurnedOn {
					pulseTypeToSend = low
				} else {
					pulseTypeToSend = high
				}
			} else {
				// If not flipflop, then check whether it's a conjunctive module?
				_, isConj := conjStates[p.to]

				if isConj {
					// Update memory of conj module first
					conjStates[p.to][p.from] = p.pulseTyp

					// If every remembered pulse is high, then send low, otherwise high
					if utils.Every(utils.Values(conjStates[p.to]), func(typ pulseType) bool { return typ == high }) {
						pulseTypeToSend = low
					} else {
						pulseTypeToSend = high
					}
				} else {
					// To-node must be output node, so nothing should be done
					continue
				}
			}

			for _, neighbor := range dgraph.Neighbors[p.to] {
				q.Append(pulse{from: p.to, to: neighbor, pulseTyp: pulseTypeToSend})
			}
		}

		i--
	}

	result := pulseCounts[low] * pulseCounts[high]
	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// Your puzzle answer was 834323022.

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

const solutionModuleName = "rx"

func Second(lines []string) (strigifiedResult string) {
	fmt.Println("\n--- Second ---")

	// Parse graph and init data structures
	dgraph, ffStates, conjStates := parseLines(lines)

	// Solve

	// Determine the sole input module for the solution module
	incNodesToSol := dgraph.GetIncomingNodesOf(solutionModuleName)
	if len(incNodesToSol) != 1 {
		panic("Only one node should feed into the solution module")
	}
	var feedToSol string = incNodesToSol[0]

	// Assert that it's a conjunctive module
	_, isConjMod := conjStates[feedToSol]
	if !isConjMod {
		panic("Input module to solution should be conjuctive module")
	}

	// Get input modules to the feed module (feed = input to solution)
	cycleCounts := map[string]int{}
	incNodesToFeed := dgraph.GetIncomingNodesOf(feedToSol)

	// Count of the time the watchlisted modules emit a high pulse
	seenCountMap := map[string]int{}
	for _, inc := range incNodesToFeed {
		seenCountMap[inc] = 0
	}

	buttonPressCount := 0

	result := -1

buttonPressing:
	for {
		buttonPressCount++

		// fmt.Println(fmt.Sprintf("\n---- Button press %d ---\n", buttonPressCount))

		q := collections.NewQueue[pulse]()
		for _, broadcasterNeighbor := range dgraph.Neighbors[broadcasterNode] {
			q.Append(pulse{from: broadcasterNode, to: broadcasterNeighbor, pulseTyp: low})
		}

		for !q.IsEmpty() {

			p := q.Pop()

			// fmt.Printf("%s -%s-> %s\n", p.from, p.pulseTyp, p.to)

			// Could also check utils.Contains(incNodesToFeed, p.from), but the feed input modules have only one output
			if p.to == feedToSol && p.pulseTyp == high {
				seenCountMap[p.from] += 1

				_, alreadyPresent := cycleCounts[p.from]
				if !alreadyPresent {
					cycleCounts[p.from] = buttonPressCount
				} else {
					if buttonPressCount != seenCountMap[p.from]*cycleCounts[p.from] {
						panic("Incorrect button press count, not multiple of cycle count!")
					}
				}

				// If every input to the feed was seen at least once
				seenCounts := utils.Values(seenCountMap)
				if utils.Every(seenCounts, func(count int) bool { return count > 0 }) {
					result = utils.LCMArr(utils.Values(cycleCounts))
					break buttonPressing
				}
			}

			// Transmit forward the correct pulse:

			// 1. Determine pulse type to send
			var pulseTypeToSend pulseType
			wasTurnedOn, isFlipFlop := ffStates[p.to]
			if isFlipFlop {
				if p.pulseTyp == high {
					continue
				}
				ffStates[p.to] = !wasTurnedOn

				if wasTurnedOn {
					pulseTypeToSend = low
				} else {
					pulseTypeToSend = high
				}
			} else {
				// If not flipflop, then check whether it's a conjunctive module?
				_, isConj := conjStates[p.to]

				if isConj {
					// Update memory of conj module first
					conjStates[p.to][p.from] = p.pulseTyp

					// If every remembered pulse is high, then send low, otherwise high
					if utils.Every(utils.Values(conjStates[p.to]), func(typ pulseType) bool { return typ == high }) {
						pulseTypeToSend = low
					} else {
						pulseTypeToSend = high
					}
				} else {
					// To-node must be output node, so nothing should be done
					continue
				}
			}

			// 2. Send pulses
			for _, neighbor := range dgraph.Neighbors[p.to] {
				q.Append(pulse{from: p.to, to: neighbor, pulseTyp: pulseTypeToSend})
			}
		}
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(2)

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	result := First(lines)
	fmt.Println(result)

	result = Second(lines)
	fmt.Println(result)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
