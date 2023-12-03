package utils

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
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

// Splits the string in 2 and returns the resulting 2 strings
func SplitIn2(str string, sep string) (string, string) {
	split := strings.Split(str, sep)
	return split[0], split[1]
}

// Returns the input file extension based on the command line args.
func GetInputFileExt() string {
	useRealInputP := flag.Bool("r", false, "Specify if you want to run the solution against the real personalized input. By default the example provided in the AoC problem description is used.")
	flag.Parse()

	var inputFileExtension string
	if *useRealInputP {
		inputFileExtension = "in"
	} else {
		fmt.Println("Using example input")
		fmt.Println("-------------------")
		inputFileExtension = "exin"
	}

	return inputFileExtension
}
