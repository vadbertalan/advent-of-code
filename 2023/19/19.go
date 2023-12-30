// https://adventofcode.com/2023/day/19

package main

import (
	"aoc/utils"
	"aoc/utils/collections"
	"aoc/utils/graph"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const aocDay int = 19

type part struct {
	x, m, a, s int
}

type wfRule struct {
	test    func(part) bool
	nextwf  string
	condStr string
}

type wf struct {
	name   string
	rules  []wfRule
	elsewf string
}

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func first(lines []string) {
	fmt.Println("--- First ---")

	result := 0

	partLinesStartAt := -1

	wfMap := map[string]wf{}
	for i, line := range lines {
		if line == "" {
			partLinesStartAt = i + 1
			break
		}
		wfName, rulesStr := utils.SplitIn2(line, "{")

		rulesSplit := strings.Split(rulesStr[:len(rulesStr)-1], ",")
		elsewf := rulesSplit[len(rulesSplit)-1]

		rules := []wfRule{}

		for _, ruleStr := range rulesSplit[:len(rulesSplit)-1] {
			condStr, nextwf := utils.SplitIn2(ruleStr, ":")
			var test func(part) bool
			if strings.Contains(condStr, "<") {
				charProp, operandStr := utils.SplitIn2(condStr, "<")
				operand, _ := strconv.Atoi(operandStr)
				test = func(part part) bool {
					switch charProp {
					case "x":
						return part.x < operand
					case "m":
						return part.m < operand
					case "a":
						return part.a < operand
					case "s":
						return part.s < operand
					default:
						panic("wrong char property found in rule")
					}
				}
			} else if strings.Contains(condStr, ">") {
				charProp, operandStr := utils.SplitIn2(condStr, ">")
				operand, _ := strconv.Atoi(operandStr)
				test = func(part part) bool {
					switch charProp {
					case "x":
						return part.x > operand
					case "m":
						return part.m > operand
					case "a":
						return part.a > operand
					case "s":
						return part.s > operand
					default:
						panic("wrong char property found in rule")
					}
				}
			} else {
				panic("Other operator than '>' or '<'")
			}

			rules = append(rules, wfRule{test: test, nextwf: nextwf, condStr: condStr})
		}

		wfMap[wfName] = wf{wfName, rules, elsewf}
	}

	parts := []part{}

	for _, line := range lines[partLinesStartAt:] {
		line = line[1 : len(line)-1]
		assignments := strings.Split(line, ",")
		part := part{}
		for _, ass := range assignments {
			charProp, valStr := utils.SplitIn2(ass, "=")
			value, _ := strconv.Atoi(valStr)
			switch charProp {
			case "x":
				part.x = value
			case "m":
				part.m = value
			case "a":
				part.a = value
			case "s":
				part.s = value
			default:
				panic("wrong char property found in assignment")
			}
		}
		parts = append(parts, part)
	}

	// Find accepted parts
	for _, part := range parts {
		wf := wfMap["in"]
	processPart:
		for {
			for _, rule := range wf.rules {
				if rule.test(part) {
					if rule.nextwf == "A" {
						result += part.x + part.m + part.a + part.s
						break processPart
					}
					if rule.nextwf == "R" {
						break processPart
					}
					wf = wfMap[rule.nextwf]
					continue processPart
				}
			}
			if wf.elsewf == "A" {
				result += part.x + part.m + part.a + part.s
				break processPart
			}
			if wf.elsewf == "R" {
				break processPart
			}
			wf = wfMap[wf.elsewf]
		}
	}

	fmt.Println(result)
}

// 306276 too low for p1
// 350678 good answer

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

const limit = 4000

type pathRule struct {
	name       string
	xmasRanges map[string][2]int
}

type node = graph.Node[pathRule]

type wfNode struct {
	parentNode *node
	value      pathRule
}

type rng struct {
	charProp string
	interval [2]int
}

func getValidRange(condStr string) (string, int, int) {
	var op string = ">"
	if strings.Contains(condStr, "<") {
		op = "<"
	}
	charProp, operandStr := utils.SplitIn2(condStr, op)
	operand, _ := strconv.Atoi(operandStr)

	if op == ">" {
		return charProp, operand + 1, limit
	}
	return charProp, 1, operand - 1
}

func getElseValidRange(condStr string) (string, int, int) {
	var op string = ">"
	if strings.Contains(condStr, "<") {
		op = "<"
	}
	charProp, operandStr := utils.SplitIn2(condStr, op)
	operand, _ := strconv.Atoi(operandStr)

	if op == ">" {
		return charProp, 1, operand
	}
	return charProp, operand, limit
}

func shrinkRange(toShrink [2]int, new [2]int) [2]int {
	min := utils.Max(toShrink[0], new[0])
	max := utils.Min(toShrink[1], new[1])
	return [2]int{min, max}
}

func shrinkXmasRanges(xmasRanges *map[string][2]int, rng rng) {
	switch rng.charProp {
	case "x":
		(*xmasRanges)["x"] = shrinkRange((*xmasRanges)["x"], rng.interval)
	case "m":
		(*xmasRanges)["m"] = shrinkRange((*xmasRanges)["m"], rng.interval)
	case "a":
		(*xmasRanges)["a"] = shrinkRange((*xmasRanges)["a"], rng.interval)
	case "s":
		(*xmasRanges)["s"] = shrinkRange((*xmasRanges)["s"], rng.interval)
	default:
		panic("wrong charProps while shrinking range")
	}
}

func second(lines []string) {
	fmt.Println("\n--- Second ---")

	result := 0

	wfMap := map[string]wf{}
	for _, line := range lines {
		if line == "" {
			break
		}
		wfName, rulesStr := utils.SplitIn2(line, "{")

		rulesSplit := strings.Split(rulesStr[:len(rulesStr)-1], ",")
		elsewf := rulesSplit[len(rulesSplit)-1]

		rules := []wfRule{}

		for _, ruleStr := range rulesSplit[:len(rulesSplit)-1] {
			condStr, nextwf := utils.SplitIn2(ruleStr, ":")
			var test func(part) bool
			if strings.Contains(condStr, "<") {
				charProp, operandStr := utils.SplitIn2(condStr, "<")
				operand, _ := strconv.Atoi(operandStr)
				test = func(part part) bool {
					switch charProp {
					case "x":
						return part.x < operand
					case "m":
						return part.m < operand
					case "a":
						return part.a < operand
					case "s":
						return part.s < operand
					default:
						panic("wrong char property found in rule")
					}
				}
			} else if strings.Contains(condStr, ">") {
				charProp, operandStr := utils.SplitIn2(condStr, ">")
				operand, _ := strconv.Atoi(operandStr)
				test = func(part part) bool {
					switch charProp {
					case "x":
						return part.x > operand
					case "m":
						return part.m > operand
					case "a":
						return part.a > operand
					case "s":
						return part.s > operand
					default:
						panic("wrong char property found in rule")
					}
				}
			} else {
				panic("Other operator than '>' or '<'")
			}

			rules = append(rules, wfRule{test: test, nextwf: nextwf, condStr: condStr})
		}

		wfMap[wfName] = wf{wfName, rules, elsewf}
	}

	rootName := "in"

	// Build tree
	var rootNode *node

	queue := collections.Queue[wfNode]{}
	queue.Append(wfNode{parentNode: nil, value: pathRule{name: rootName}})

	for !queue.IsEmpty() {
		wfn := queue.Pop()

		// A or R are leaves
		if wfn.value.name == "A" || wfn.value.name == "R" {
			leafNode := &node{Value: wfn.value, Parent: wfn.parentNode, Children: []*node{}}
			wfn.parentNode.Children = append(wfn.parentNode.Children, leafNode)
			continue
		}

		wf := wfMap[wfn.value.name]
		itNode := &node{Value: wfn.value, Parent: wfn.parentNode, Children: []*node{}}

		// Root node
		if rootNode == nil {
			rootNode = itNode
		} else {
			wfn.parentNode.Children = append(wfn.parentNode.Children, itNode)
		}

		// Keep account of xmas range modifications too while building tree
		xmasRanges := map[string][2]int{
			"x": {1, limit},
			"m": {1, limit},
			"a": {1, limit},
			"s": {1, limit},
		}
		for _, wfRule := range wf.rules {
			charProp, min, max := getValidRange(wfRule.condStr)
			copyXmasRanges := utils.ShallowCopyMap(xmasRanges)
			shrinkXmasRanges(&copyXmasRanges, rng{charProp: charProp, interval: [2]int{min, max}})
			queue.Append(wfNode{
				parentNode: itNode,
				value: pathRule{
					name:       wfRule.nextwf,
					xmasRanges: copyXmasRanges,
				},
			})

			// Prepping xmasRanges for new rule or for the last elseWf
			charProp, min, max = getElseValidRange(wfRule.condStr)
			shrinkXmasRanges(&xmasRanges, rng{charProp: charProp, interval: [2]int{min, max}})
		}

		// Else range is already up to date
		queue.Append(wfNode{parentNode: itNode, value: pathRule{name: wf.elsewf, xmasRanges: xmasRanges}})
	}

	tree := graph.Tree[pathRule]{Root: rootNode}

	allPaths := tree.FindAllPaths()
	acceptedPaths := utils.Filter[[]pathRule](allPaths, func(path []pathRule) bool { return path[len(path)-1].name == "A" })

	// Traverse accepted paths and
	for _, path := range acceptedPaths {
		xmasRanges := map[string][2]int{
			"x": {1, limit},
			"m": {1, limit},
			"a": {1, limit},
			"s": {1, limit},
		}

		// Start from 1 to skip starting node "in"
		for i := 1; i < len(path); i++ {
			pr := path[i]

			xmasRanges["x"] = shrinkRange(xmasRanges["x"], pr.xmasRanges["x"])
			xmasRanges["m"] = shrinkRange(xmasRanges["m"], pr.xmasRanges["m"])
			xmasRanges["a"] = shrinkRange(xmasRanges["a"], pr.xmasRanges["a"])
			xmasRanges["s"] = shrinkRange(xmasRanges["s"], pr.xmasRanges["s"])
		}

		// Summarize range
		pathPoss := (xmasRanges["x"][1] - xmasRanges["x"][0] + 1) * (xmasRanges["m"][1] - xmasRanges["m"][0] + 1) * (xmasRanges["a"][1] - xmasRanges["a"][0] + 1) * (xmasRanges["s"][1] - xmasRanges["s"][0] + 1)
		result += pathPoss
	}

	fmt.Println(result)
}

// Good answer: 124831893423809

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(2)

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	first(lines)

	second(lines)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
