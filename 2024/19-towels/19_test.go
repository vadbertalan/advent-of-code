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
	expected := "298"
	
	lines := getInputLines()
	result := First(lines)

	if result != expected {
		t.Errorf("First() = %v; want %v", result, expected)
	}
}

func TestFirstWithExampleInput(t *testing.T) {
	expected := "6"
	
	lines := getExampleInputLines()
	result := First(lines)

	if result != expected {
		t.Errorf("First() = %v; want %v", result, expected)
	}
}

func TestSecond(t *testing.T) {
	expected := "572248688842069"
	
	lines := getInputLines()
	result := Second(lines)

	if result != expected {
		t.Errorf("Second() = %v; want %v", result, expected)
	}
}

func TestSecondWithExampleInput(t *testing.T) {
	expected := "16"
	
	lines := getExampleInputLines()
	result := Second(lines)

	if result != expected {
		t.Errorf("Second() = %v; want %v", result, expected)
	}
}
