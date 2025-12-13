package main

import (
	"aoc/utils-go"
	"fmt"
	"testing"
)

func getInputLines() []string {
	lines := utils.ReadLines(fmt.Sprintf("%d.in", aocDay))
	return lines
}

func getExampleInputLines(exNr int) []string {
	if exNr == 0 {
		return utils.ReadLines(fmt.Sprintf("%d.exin", aocDay))
	}
	lines := utils.ReadLines(fmt.Sprintf("%d.exin%d", aocDay, exNr))
	return lines
}

func TestFirst(t *testing.T) {
	expected := "1573474"

	lines := getInputLines()
	result := First(lines)

	if result != expected {
		t.Errorf("First() = %v; want %v", result, expected)
	}
}

func TestFirstWithExample1Input(t *testing.T) {
	expected := "140"

	lines := getExampleInputLines(0)
	result := First(lines)

	if result != expected {
		t.Errorf("First() = %v; want %v", result, expected)
	}
}
func TestFirstWithExample2Input(t *testing.T) {
	expected := "772"

	lines := getExampleInputLines(2)
	result := First(lines)

	if result != expected {
		t.Errorf("First() = %v; want %v", result, expected)
	}
}

func TestFirstWithExample3Input(t *testing.T) {
	expected := "1930"

	lines := getExampleInputLines(3)
	result := First(lines)

	if result != expected {
		t.Errorf("First() = %v; want %v", result, expected)
	}
}

func TestFirstWithExample4Input(t *testing.T) {
	expected := "692"

	lines := getExampleInputLines(4)
	result := First(lines)

	if result != expected {
		t.Errorf("First() = %v; want %v", result, expected)
	}
}

func TestSecond(t *testing.T) {
	expected := "966476"

	lines := getInputLines()
	result := Second(lines)

	if result != expected {
		t.Errorf("Second() = %v; want %v", result, expected)
	}
}

func TestSecondWithExampleInput(t *testing.T) {
	expected := "80"

	lines := getExampleInputLines(0)
	result := Second(lines)

	if result != expected {
		t.Errorf("Second() = %v; want %v", result, expected)
	}
}

func TestSecondWithExample2Input(t *testing.T) {
	expected := "436"

	lines := getExampleInputLines(2)
	result := Second(lines)

	if result != expected {
		t.Errorf("Second() = %v; want %v", result, expected)
	}
}

func TestSecondWithExample3Input(t *testing.T) {
	expected := "1206"

	lines := getExampleInputLines(3)
	result := Second(lines)

	if result != expected {
		t.Errorf("Second() = %v; want %v", result, expected)
	}
}

func TestSecondWithExample4Input(t *testing.T) {
	expected := "236"

	lines := getExampleInputLines(4)
	result := Second(lines)

	if result != expected {
		t.Errorf("Second() = %v; want %v", result, expected)
	}
}

func TestSecondWithExample5Input(t *testing.T) {
	expected := "368"

	lines := getExampleInputLines(5)
	result := Second(lines)

	if result != expected {
		t.Errorf("Second() = %v; want %v", result, expected)
	}
}
