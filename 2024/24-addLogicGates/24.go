// https://adventofcode.com/2024/day/24

package main

import (
	"aoc/utils-go"
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
)

const aocDay int = 24

type wire struct {
	name          string
	value         *bool
	inputForGates []*gate
	outputForGate *gate
}

type operation int

const (
	xor operation = iota
	or
	and
)

type gate struct {
	op         operation
	inputWires [2]*wire
	outputWire *wire
}

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func First(lines []string) (strigifiedResult string) {
	_, emptyLineIndex := utils.Find(lines, func(line string) bool { return line == "" })

	wireMap := make(map[string]*wire)
	gates := []*gate{}

	// Parse initial input wires
	for _, line := range lines[:emptyLineIndex] {
		wireName, valueStr := utils.SplitIn2(line, ": ")

		boolValue := valueStr == "1"
		w := wire{name: wireName, value: &boolValue}
		wireMap[wireName] = &w
	}

	// Parse gates
	for _, line := range lines[emptyLineIndex+1:] {
		strParts := strings.Split(line, " ")

		opStr := strParts[1]
		var op operation
		switch opStr {
		case "AND":
			op = and
		case "OR":
			op = or
		case "XOR":
			op = xor
		default:
			panic("Unknown operation")
		}

		g := gate{op: op}

		inputWire1Name := strParts[0]
		inputWire2Name := strParts[2]
		outputWireName := strParts[4]

		if w, ok := wireMap[inputWire1Name]; !ok {
			newWire := wire{name: inputWire1Name}
			newWire.inputForGates = append(newWire.inputForGates, &g)
			wireMap[inputWire1Name] = &newWire
		} else {
			w.inputForGates = append(w.inputForGates, &g)
		}

		if w, ok := wireMap[inputWire2Name]; !ok {
			newWire := wire{name: inputWire2Name}
			newWire.inputForGates = append(newWire.inputForGates, &g)
			wireMap[inputWire2Name] = &newWire
		} else {
			w.inputForGates = append(w.inputForGates, &g)
		}

		if w, ok := wireMap[outputWireName]; !ok {
			newWire := wire{name: outputWireName}
			newWire.outputForGate = &g
			wireMap[outputWireName] = &newWire
		} else {
			w.outputForGate = &g
		}

		g.inputWires[0] = wireMap[inputWire1Name]
		g.inputWires[1] = wireMap[inputWire2Name]
		g.outputWire = wireMap[outputWireName]

		gates = append(gates, &g)
	}

	// Solve: make sure every output wire
	for utils.Some(gates, func(g *gate) bool {
		return g.outputWire.value == nil
	}) {
		for _, g := range gates {
			// fmt.Printf("%v %v %v -> %v\n", g.inputWires[0].name, g.op, g.inputWires[1].name, g.outputWire.name)

			if g.inputWires[0].value != nil && g.inputWires[1].value != nil && g.outputWire.value == nil {
				var opResult bool
				switch g.op {
				case xor:
					opResult = *g.inputWires[0].value != *g.inputWires[1].value
				case or:
					opResult = *g.inputWires[0].value || *g.inputWires[1].value
				case and:
					opResult = *g.inputWires[0].value && *g.inputWires[1].value
				}

				g.outputWire.value = &opResult
			}
		}
	}

	// Calculate result: look at the zXY wire values
	result := 0

	zNr := 0
	zStr := fmt.Sprintf("z%02d", zNr)
	zValue := wireMap[zStr].value
	for zValue != nil {
		if *zValue {
			result += int(math.Pow(2, float64(zNr)))
		}

		zNr++
		zStr = fmt.Sprintf("z%02d", zNr)
		if w, ok := wireMap[zStr]; ok {
			zValue = w.value
		} else {
			zValue = nil
		}
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// correct: 53258032898766, example1: 4, example2: 2024

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

func solveGate(g *gate) {
	var opResult bool
	switch g.op {
	case xor:
		opResult = *g.inputWires[0].value != *g.inputWires[1].value
	case or:
		opResult = *g.inputWires[0].value || *g.inputWires[1].value
	case and:
		opResult = *g.inputWires[0].value && *g.inputWires[1].value
	}

	g.outputWire.value = &opResult
}

func Second(lines []string) (strigifiedResult string) {
	_, emptyLineIndex := utils.Find(lines, func(line string) bool { return line == "" })

	wireMap := make(map[string]*wire)
	gates := []*gate{}

	// Parse initial input wires
	for _, line := range lines[:emptyLineIndex] {
		wireName, valueStr := utils.SplitIn2(line, ": ")

		boolValue := valueStr == "1"
		w := wire{name: wireName, value: &boolValue}
		wireMap[wireName] = &w
	}

	// Parse gates
	for _, line := range lines[emptyLineIndex+1:] {
		strParts := strings.Split(line, " ")

		opStr := strParts[1]
		var op operation
		switch opStr {
		case "AND":
			op = and
		case "OR":
			op = or
		case "XOR":
			op = xor
		default:
			panic("Unknown operation")
		}

		g := gate{op: op}

		inputWire1Name := strParts[0]
		inputWire2Name := strParts[2]
		outputWireName := strParts[4]

		if w, ok := wireMap[inputWire1Name]; !ok {
			newWire := wire{name: inputWire1Name}
			newWire.inputForGates = append(newWire.inputForGates, &g)
			wireMap[inputWire1Name] = &newWire
		} else {
			w.inputForGates = append(w.inputForGates, &g)
		}

		if w, ok := wireMap[inputWire2Name]; !ok {
			newWire := wire{name: inputWire2Name}
			newWire.inputForGates = append(newWire.inputForGates, &g)
			wireMap[inputWire2Name] = &newWire
		} else {
			w.inputForGates = append(w.inputForGates, &g)
		}

		if w, ok := wireMap[outputWireName]; !ok {
			newWire := wire{name: outputWireName}
			newWire.outputForGate = &g
			wireMap[outputWireName] = &newWire
		} else {
			w.outputForGate = &g
		}

		g.inputWires[0] = wireMap[inputWire1Name]
		g.inputWires[1] = wireMap[inputWire2Name]
		g.outputWire = wireMap[outputWireName]

		gates = append(gates, &g)
	}

	// gatesInOrderOfSolving := []*gate{}

	// Solve: make sure every output wire
	for utils.Some(gates, func(g *gate) bool {
		return g.outputWire.value == nil
	}) {
		for _, g := range gates {
			// fmt.Printf("%v %v %v -> %v\n", g.inputWires[0].name, g.op, g.inputWires[1].name, g.outputWire.name)

			if g.inputWires[0].value != nil && g.inputWires[1].value != nil && g.outputWire.value == nil {
				solveGate(g)
				// gatesInOrderOfSolving = append(gatesInOrderOfSolving, g)
			}
		}
	}

	// hard coded wireSwaps
	wireSwaps := [][2]*wire{
		{wireMap["wss"], wireMap["wrm"]},
		{wireMap["z08"], wireMap["thm"]},
		{wireMap["z29"], wireMap["gbs"]},
		{wireMap["hwq"], wireMap["z22"]},
	}

	// for _, wireSwap := range wireSwaps {
	// 	tempWire := wireSwap[0].outputForGate.outputWire
	// 	wireSwap[0].outputForGate.outputWire = wireSwap[1].outputForGate.outputWire
	// 	wireSwap[1].outputForGate.outputWire = tempWire

	// 	solveGate(wireSwap[0].outputForGate)
	// 	solveGate(wireSwap[1].outputForGate)

	// 	// Resolve system
	// 	for _, g := range gatesInOrderOfSolving {
	// 		solveGate(g)
	// 	}
	// }

	// f, _ := os.Create("logic_gates.dot")
	// defer f.Close()

	// ff := func(op operation) string {
	// 	switch op {
	// 	case xor:
	// 		return "XOR"
	// 	case or:
	// 		return "OR"
	// 	case and:
	// 		return "AND"
	// 	}
	// 	return ""
	// }

	// vv := func(v *bool) string {
	// 	if v == nil {
	// 		return "nil"
	// 	}
	// 	if *v {
	// 		return "1"
	// 	}
	// 	return "0"
	// }

	// shouldBeBin := "0111001101000100010001000101111111110110000011"

	// fmt.Fprintln(f, "digraph LogicGates {")
	// for _, g := range gatesInOrderOfSolving {
	// 	indStr := g.outputWire.name[1:]
	// 	ind, _ := strconv.Atoi(indStr)
	// 	isProblematicGate := false

	// 	// example wrm
	// 	if g.op == xor &&
	// 		(g.inputWires[0].name[0] == 'x' && g.inputWires[1].name[0] == 'y' || g.inputWires[0].name[0] == 'y' && g.inputWires[1].name[0] == 'x') &&
	// 		len(g.outputWire.inputForGates) != 2 &&
	// 		g.outputWire.name != "z00" {
	// 		isProblematicGate = true
	// 	} else if g.op == xor &&
	// 		(g.inputWires[0].name[0] == 'x' && g.inputWires[1].name[0] == 'y' || g.inputWires[0].name[0] == 'y' && g.inputWires[1].name[0] == 'x') &&
	// 		g.outputWire.name != "z00" {

	// 		xorGate, _ := utils.Find(g.outputWire.inputForGates, func(g *gate) bool {
	// 			return g.op == xor
	// 		})
	// 		if xorGate == nil {
	// 			panic("no xor gate, but should be")
	// 		}
	// 		if !strings.HasPrefix((*xorGate).outputWire.name, "z") && (*xorGate).outputWire.name != "z00" {
	// 			isProblematicGate = true
	// 		}
	// 	}

	// 	// example ?
	// 	if g.op == xor &&
	// 		(g.inputWires[0].name[0] == 'x' && g.inputWires[1].name[0] == 'y' || g.inputWires[0].name[0] == 'y' && g.inputWires[1].name[0] == 'x') &&
	// 		len(g.outputWire.inputForGates) != 2 &&
	// 		g.outputWire.name != "z00" &&
	// 		len(g.outputWire.inputForGates) == 2 &&
	// 		// next gates should have and and xor
	// 		!(g.outputWire.inputForGates[0].op == and && g.outputWire.inputForGates[1].op == xor || g.outputWire.inputForGates[0].op == xor && g.outputWire.inputForGates[1].op == and) {
	// 		isProblematicGate = true
	// 	}

	// 	// example wss
	// 	if g.op == and &&
	// 		(g.inputWires[0].name[0] == 'x' && g.inputWires[1].name[0] == 'y' || g.inputWires[0].name[0] == 'y' && g.inputWires[1].name[0] == 'x') &&
	// 		len(g.outputWire.inputForGates) != 1 &&
	// 		(g.inputWires[0].name != "x00" && g.inputWires[1].name != "y00" && g.inputWires[0].name != "y00" && g.inputWires[1].name != "x00") {

	// 		isProblematicGate = true

	// 	} else if g.op == and && // example ?
	// 		(g.inputWires[0].name[0] == 'x' && g.inputWires[1].name[0] == 'y' || g.inputWires[0].name[0] == 'y' && g.inputWires[1].name[0] == 'x') &&
	// 		len(g.outputWire.inputForGates) != 1 &&
	// 		(g.inputWires[0].name != "x00" && g.inputWires[1].name != "y00" && g.inputWires[0].name != "y00" && g.inputWires[1].name != "x00") &&
	// 		// AND next gates should have OR
	// 		(g.outputWire.inputForGates[0].op != or) {

	// 		isProblematicGate = true
	// 	} else if g.op == and && // example ?
	// 		(g.inputWires[0].name[0] == 'x' && g.inputWires[1].name[0] == 'y' || g.inputWires[0].name[0] == 'y' && g.inputWires[1].name[0] == 'x') &&
	// 		len(g.outputWire.inputForGates) != 1 &&
	// 		(g.inputWires[0].name != "x00" && g.inputWires[1].name != "y00" && g.inputWires[0].name != "y00" && g.inputWires[1].name != "x00") &&
	// 		// AND next gates should have OR
	// 		(g.outputWire.inputForGates[0].op != or) &&
	// 		len(g.outputWire.inputForGates[0].outputWire.inputForGates) != 2 &&
	// 		// next next gates should have XOR and AND
	// 		!(g.outputWire.inputForGates[0].outputWire.inputForGates[0].op == xor && g.outputWire.inputForGates[0].outputWire.inputForGates[1].op == and || g.outputWire.inputForGates[0].outputWire.inputForGates[0].op == and && g.outputWire.inputForGates[0].outputWire.inputForGates[1].op == xor) {

	// 		isProblematicGate = true
	// 	}

	// 	if strings.HasPrefix(g.outputWire.name, "z") && g.op != xor && g.outputWire.name != "z45" {
	// 		isProblematicGate = true
	// 	}

	// 	if strings.HasPrefix(g.outputWire.name, "z") {
	// 		fmt.Fprintf(f, "  %s [style=filled, fillcolor=lightblue]\n", g.outputWire.name)
	// 		if g.outputWire.value != nil {
	// 			if vv(g.outputWire.value) != string(shouldBeBin[ind]) {
	// 				fmt.Fprintf(f, "  %s [fontcolor=red, fontname=\"Helvetica-Bold\"]\n", g.outputWire.name)
	// 			}
	// 			fmt.Fprintf(f, "  %s [xlabel=\"%s|%s\"]\n", g.outputWire.name, vv(g.outputWire.value), string(shouldBeBin[ind]))
	// 		} else {
	// 			panic("output wire has no value")
	// 		}
	// 	}
	// 	if strings.HasPrefix(g.inputWires[0].name, "x") {
	// 		fmt.Fprintf(f, "  %s [style=filled, fillcolor=lightgreen]\n", g.inputWires[0].name)
	// 	}
	// 	if strings.HasPrefix(g.inputWires[1].name, "x") {
	// 		fmt.Fprintf(f, "  %s [style=filled, fillcolor=lightgreen]\n", g.inputWires[1].name)
	// 	}
	// 	if strings.HasPrefix(g.inputWires[0].name, "y") {
	// 		fmt.Fprintf(f, "  %s [style=filled, fillcolor=lightgreen]\n", g.inputWires[0].name)
	// 	}
	// 	if strings.HasPrefix(g.inputWires[1].name, "y") {
	// 		fmt.Fprintf(f, "  %s [style=filled, fillcolor=lightgreen]\n", g.inputWires[1].name)
	// 	}
	// 	if isProblematicGate {
	// 		fmt.Fprintf(f, "  %s [style=filled, fillcolor=darkred]\n", g.outputWire.name)
	// 	}
	// 	fmt.Fprintf(f, "  %s -> %s [label=\"%s\"]\n", g.inputWires[0].name, g.outputWire.name, ff(g.op))
	// 	fmt.Fprintf(f, "  %s -> %s [label=\"%s\"]\n", g.inputWires[1].name, g.outputWire.name, ff(g.op))

	// 	if g.inputWires[0].value != nil {
	// 		fmt.Fprintf(f, "  %s [xlabel=\"%s\"]\n", g.inputWires[0].name, vv(g.inputWires[0].value))
	// 	}
	// 	if g.inputWires[1].value != nil {
	// 		fmt.Fprintf(f, "  %s [xlabel=\"%s\"]\n", g.inputWires[1].name, vv(g.inputWires[1].value))
	// 	}
	// }

	// fmt.Fprintln(f, "}")

	outputWires := []string{}
	for _, wireSwap := range wireSwaps {
		for _, w := range wireSwap {
			outputWires = append(outputWires, w.name)
		}
	}
	sort.Strings(outputWires)

	result := strings.Join(outputWires, ",")
	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// correct: gbs,hwq,thm,wrm,wss,z08,z22,z29

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
