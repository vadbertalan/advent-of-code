package utils

import (
	"bufio"
	"os"
	"strconv"
)

// Read line by line into memory.
// All file contents is stores in lines[]
func ReadLines(path string) (lines []string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// Converts input param strings array to an array of ints by parsing ints
func ConvertToInts(strs []string) (ints []int) {
	for _, strval := range strs {
		intval, err := strconv.Atoi(strval)
		if err != nil {
			panic(err)
		}
		ints = append(ints, intval)
	}
	return ints
}
