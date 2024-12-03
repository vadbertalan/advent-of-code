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

func TestFirst(t *testing.T) {
	expected := "161289189"

	lines := getInputLines()
	result := First(lines)

	if result != expected {
		t.Errorf("First() = %v; want %v", result, expected)
	}
}

func TestFirstExample(t *testing.T) {
	expected := "161"

	lines := getExampleInputLines()
	result := First(lines)

	if result != expected {
		t.Errorf("First() = %v; want %v", result, expected)
	}
}

func TestSecond(t *testing.T) {
	expected := "83595109"

	lines := getInputLines()
	result := Second(lines)

	if result != expected {
		t.Errorf("Second() = %v; want %v", result, expected)
	}
}

func TestSecondExample(t *testing.T) {
	expected := "48"

	lines := getExample2InputLines()
	result := Second(lines)

	if result != expected {
		t.Errorf("Second() = %v; want %v", result, expected)
	}
}
