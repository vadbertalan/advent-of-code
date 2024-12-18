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
	expected := "334"

	lines := getInputLines()
	result := First(lines, 70, 70, 1024)

	if result != expected {
		t.Errorf("First() = %v; want %v", result, expected)
	}
}

func TestFirstWithExampleInput(t *testing.T) {
	expected := "22"

	lines := getExampleInputLines()
	result := First(lines, 6, 6, 12)

	if result != expected {
		t.Errorf("First() = %v; want %v", result, expected)
	}
}

func TestSecond(t *testing.T) {
	expected := "20,12"

	lines := getInputLines()
	result := Second(lines, 70, 70, 1024)

	if result != expected {
		t.Errorf("Second() = %v; want %v", result, expected)
	}
}

func TestSecondWithExampleInput(t *testing.T) {
	expected := "6,1"

	lines := getExampleInputLines()
	result := Second(lines, 6, 6, 12)

	if result != expected {
		t.Errorf("Second() = %v; want %v", result, expected)
	}
}
