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

func getExampleInputLines(exin int) []string {
	if exin == 0 {
		return utils.ReadLines(fmt.Sprintf("%d.exin", aocDay))
	}
	lines := utils.ReadLines(fmt.Sprintf("%d.exin%d", aocDay, exin))
	return lines
}

func TestFirst(t *testing.T) {
	expected := "1446158"

	lines := getInputLines()
	result := First(lines)

	if result != expected {
		t.Errorf("First() = %v; want %v", result, expected)
	}
}

func TestFirstWithExampleInput(t *testing.T) {
	expected := "2028"

	lines := getExampleInputLines(0)
	result := First(lines)

	if result != expected {
		t.Errorf("First() = %v; want %v", result, expected)
	}
}

func TestFirstWithExample2Input(t *testing.T) {
	expected := "10092"

	lines := getExampleInputLines(2)
	result := First(lines)

	if result != expected {
		t.Errorf("First() = %v; want %v", result, expected)
	}
}

func TestSecond(t *testing.T) {
	expected := "1446175"

	lines := getInputLines()
	result := Second(lines)

	if result != expected {
		t.Errorf("Second() = %v; want %v", result, expected)
	}
}

func TestSecondWithExample2Input(t *testing.T) {
	expected := "9021"

	lines := getExampleInputLines(2)
	result := Second(lines)

	if result != expected {
		t.Errorf("Second() = %v; want %v", result, expected)
	}
}

func TestSecondWithExample3Input(t *testing.T) {
	expected := "1115"

	lines := getExampleInputLines(3)
	result := Second(lines)

	if result != expected {
		t.Errorf("Second() = %v; want %v", result, expected)
	}
}

func TestSecondWithExample4Input(t *testing.T) {
	expected := "618"

	lines := getExampleInputLines(4)
	result := Second(lines)

	if result != expected {
		t.Errorf("Second() = %v; want %v", result, expected)
	}
}
