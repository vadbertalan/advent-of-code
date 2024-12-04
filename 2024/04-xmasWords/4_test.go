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

func getExample2InputLines() []string {
	lines := utils.ReadLines(fmt.Sprintf("%d.exin2", aocDay))
	return lines
}

func getExample3InputLines() []string {
	lines := utils.ReadLines(fmt.Sprintf("%d.exin3", aocDay))
	return lines
}

func getExample4InputLines() []string {
	lines := utils.ReadLines(fmt.Sprintf("%d.exin4", aocDay))
	return lines
}

func TestFirst(t *testing.T) {
	expected := "2551"

	lines := getInputLines()
	result := First(lines)

	if result != expected {
		t.Errorf("First() = %v; want %v", result, expected)
	}
}

func TestFirstWithExampleInput(t *testing.T) {
	expected := "4"

	lines := getExampleInputLines()
	result := First(lines)

	if result != expected {
		t.Errorf("First() = %v; want %v", result, expected)
	}
}

func TestFirstWithExampleInput2(t *testing.T) {
	expected := "18"

	lines := getExample2InputLines()
	result := First(lines)

	if result != expected {
		t.Errorf("First() = %v; want %v", result, expected)
	}
}

func TestFirstWithExampleInput3(t *testing.T) {
	expected := "0"

	lines := getExample3InputLines()
	result := First(lines)

	if result != expected {
		t.Errorf("First() = %v; want %v", result, expected)
	}
}

func TestFirstWithExampleInput4(t *testing.T) {
	expected := "0"

	lines := getExample4InputLines()
	result := First(lines)

	if result != expected {
		t.Errorf("First() = %v; want %v", result, expected)
	}
}

// ============== SECOND PART ==============

func TestSecond(t *testing.T) {
	expected := "1985"

	lines := getInputLines()
	result := Second(lines)

	if result != expected {
		t.Errorf("Second() = %v; want %v", result, expected)
	}
}

func TestSecondWithExampleInput(t *testing.T) {
	expected := "0"

	lines := getExampleInputLines()
	result := Second(lines)

	if result != expected {
		t.Errorf("Second() = %v; want %v", result, expected)
	}
}

func TestSecondWithExampleInput2(t *testing.T) {
	expected := "9"

	lines := getExample2InputLines()
	result := Second(lines)

	if result != expected {
		t.Errorf("Second() = %v; want %v", result, expected)
	}
}

func TestSecondWithExampleInput3(t *testing.T) {
	expected := "1"

	lines := getExample3InputLines()
	result := Second(lines)

	if result != expected {
		t.Errorf("Second() = %v; want %v", result, expected)
	}
}

func TestSecondWithExampleInput4(t *testing.T) {
	expected := "9"

	lines := getExample4InputLines()
	result := Second(lines)

	if result != expected {
		t.Errorf("Second() = %v; want %v", result, expected)
	}
}
