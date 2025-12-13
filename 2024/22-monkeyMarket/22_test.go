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

func getExampleInputLines() []string {
	lines := utils.ReadLines(fmt.Sprintf("%d.exin", aocDay))
	return lines
}

func getExample2InputLines() []string {
	lines := utils.ReadLines(fmt.Sprintf("%d.exin2", aocDay))
	return lines
}

func TestFirst(t *testing.T) {
	expected := "16299144133"

	lines := getInputLines()
	result := First(lines)

	if result != expected {
		t.Errorf("First() = %v; want %v", result, expected)
	}
}

func TestFirstWithExampleInput(t *testing.T) {
	expected := "37327623"

	lines := getExampleInputLines()
	result := First(lines)

	if result != expected {
		t.Errorf("First() = %v; want %v", result, expected)
	}
}

func TestSecond(t *testing.T) {
	expected := "1896"

	lines := getInputLines()
	result := Second(lines)

	if result != expected {
		t.Errorf("Second() = %v; want %v", result, expected)
	}
}

func TestSecondWithExample2Input(t *testing.T) {
	expected := "23"

	lines := getExample2InputLines()
	result := Second(lines)

	if result != expected {
		t.Errorf("Second() = %v; want %v", result, expected)
	}
}
