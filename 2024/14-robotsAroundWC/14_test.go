package main

import (
	"aoc/utils"
	"fmt"
	"testing"
)

func getInputLines() []string {
	lines := utils.ReadLines(fmt.Sprintf("%d.in", aocDay))
	return lines
}

func getExampleInputLines() []string {
	lines := utils.ReadLines(fmt.Sprintf("%d.exin", aocDay))
	return lines
}

func TestFirst(t *testing.T) {
	expected := "228410028"

	lines := getInputLines()
	result := First(lines, 103, 101)

	if result != expected {
		t.Errorf("First() = %v; want %v", result, expected)
	}
}

func TestFirstWithExampleInput(t *testing.T) {
	expected := "12"

	lines := getExampleInputLines()
	result := First(lines, 7, 11)

	if result != expected {
		t.Errorf("First() = %v; want %v", result, expected)
	}
}

// No test for second, because drawing exercise
