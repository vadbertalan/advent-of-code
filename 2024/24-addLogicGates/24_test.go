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
	expected := "53258032898766"

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

func TestFirstWithExample2Input(t *testing.T) {
	expected := "2024"

	lines := getExample2InputLines()
	result := First(lines)

	if result != expected {
		t.Errorf("First() = %v; want %v", result, expected)
	}
}

func TestSecond(t *testing.T) {
	expected := "gbs,hwq,thm,wrm,wss,z08,z22,z29"

	lines := getInputLines()
	result := Second(lines)

	if result != expected {
		t.Errorf("Second() = %v; want %v", result, expected)
	}
}
