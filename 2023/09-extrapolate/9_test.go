package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"testing"
)

func getLines() []string {
	lines := utils.ReadLines(fmt.Sprintf("%d.in", aocDay))
	return lines
}

func TestFirst(t *testing.T) {
	expected := "1708206096"
	lines := getLines()

	result := First(lines)

	if strconv.Itoa(result) != expected {
		t.Errorf("First() = %v; want %v", result, expected)
	}
}

func TestSecond(t *testing.T) {
	expected := "1050"
	lines := getLines()

	result := Second(lines)

	if strconv.Itoa(result) != expected {
		t.Errorf("Second() = %v; want %v", result, expected)
	}
}
