package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
)

const testTemplate = `package main

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
	expected := "{{.FirstAnswer}}"
	
	lines := getInputLines()
	result := First(lines)

	if result != expected {
		t.Errorf("First() = %v; want %v", result, expected)
	}
}

func TestFirstExample(t *testing.T) {
	expected := "{{.FirstAnswerExample}}"
	
	lines := getExampleInputLines()
	result := First(lines)

	if result != expected {
		t.Errorf("First() = %v; want %v", result, expected)
	}
}

func TestSecond(t *testing.T) {
	expected := "{{.SecondAnswer}}"
	
	lines := getInputLines()
	result := Second(lines)

	if result != expected {
		t.Errorf("Second() = %v; want %v", result, expected)
	}
}

func TestSecondExample(t *testing.T) {
	expected := "{{.SecondAnswerExample}}"
	
	lines := getExampleInputLines()
	result := Second(lines)

	if result != expected {
		t.Errorf("Second() = %v; want %v", result, expected)
	}
}
`

type Answers struct {
	FirstAnswerExample  string
	SecondAnswerExample string
	FirstAnswer         string
	SecondAnswer        string
}

// This utility generates a Go test file
// for a specific Advent of Code challenge day and year, with provided answers for the
// first and second parts of the challenge, including example answers.
//
// Usage:
//
//	gentest <year> <day> <first_answer_example> <second_answer_example> <first_answer> <second_answer>
//
// Example:
//
//	gentest 2023 01 123 456 12345 67890
//
// The generated test file will be created in a directory named after the specified year,
// with the filename format <day>_test.go. The test file will contain the provided answers
// embedded in a template.
func main() {
	if len(os.Args) != 7 {
		fmt.Println("Usage: gentest <year> <day> <first_example_answer> <first_example_answer> <first_answer> <second_answer>")
		return
	}

	year := os.Args[1]
	day := os.Args[2]
	answers := Answers{
		FirstAnswerExample:  os.Args[3],
		SecondAnswerExample: os.Args[4],
		FirstAnswer:         os.Args[5],
		SecondAnswer:        os.Args[6],
	}

	tmpl, err := template.New("test").Parse(testTemplate)
	if err != nil {
		fmt.Println("Error creating template:", err)
		return
	}

	yearDir := filepath.Join(year)
	// yearDir := filepath.Join("..", "..", year)
	dayDir := ""

	dayInt, err := strconv.Atoi(day)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid day:", err)
		os.Exit(1)
	}

	err = filepath.Walk(filepath.Join(yearDir), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != yearDir {
			dirDay, err := strconv.Atoi(filepath.Base(path)[:2])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error finding day directory:", err)
				return nil // skip directories that don't start with a number
			}

			if dirDay == dayInt {
				dayDir = path
				return filepath.SkipDir
			}
		}
		return nil
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error finding day directory:", err)
		os.Exit(1)
	}
	if dayDir == "" {
		fmt.Fprintln(os.Stderr, "Day directory not found")
		os.Exit(1)
	}

	filePath := filepath.Join(dayDir, fmt.Sprintf("%s_test.go", day))

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	err = tmpl.Execute(file, answers)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	fmt.Println("Test file generated successfully.", filePath)
}
