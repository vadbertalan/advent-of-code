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

func TestFirst(t *testing.T) {
	expected := "3689"

	lines := getInputLines()
	result := First(lines, 64)

	if result != expected {
		t.Errorf("First() = %v; want %v", result, expected)
	}
}

func TestSecond(t *testing.T) {
	expected := "610158187362102"

	lines := getInputLines()
	result := Second(lines, 26501365)

	if result != expected {
		t.Errorf("Second() = %v; want %v", result, expected)
	}
}
