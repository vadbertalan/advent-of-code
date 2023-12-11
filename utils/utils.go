package utils

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
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
	useEx2InputP := flag.Bool("e2", false, "Specify if you want to run the solution against a second example input, which should be put in file `XX.exin2`. By default the simpler example provided in the AoC problem description is used.")
	useRealInputP := flag.Bool("r", false, "Specify if you want to run the solution against the real personalized input, which should be put in file `XX.in` beforehand. By default the example provided in the AoC problem description is used.")
	flag.Parse()

	var inputFileExtension string
	if *useRealInputP {
		inputFileExtension = "in"
	} else if *useEx2InputP {
		fmt.Println("Using example 2 input")
		fmt.Println("-------------------")
		inputFileExtension = "exin2"

	} else {
		fmt.Println("Using example input")
		fmt.Println("-------------------")
		inputFileExtension = "exin"
	}

	return inputFileExtension
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func SolveMasodfoku(a, b, c float64) (x1, x2 float64) {
	if a == 0 {
		panic("a cannot be 0")
	}

	delta := b*b - 4*a*c

	if delta < 0 {
		panic("Equation has complex solutions")
	}

	x1 = (-b + math.Sqrt(delta)) / (2 * a)
	x2 = (-b - math.Sqrt(delta)) / (2 * a)

	return x1, x2
}

func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, len(m))
	i := 0
	for k := range m {
		r[i] = k
		i++
	}
	return r
}

func Values[M ~map[K]V, K comparable, V any](m M) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

func EqualArr[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Filter[T any](ss []T, isRemaining func(T) bool) (ret []T) {
	for _, s := range ss {
		if isRemaining(s) {
			ret = append(ret, s)
		}
	}
	return
}

func Every[T any](ss []T, isValid func(T) bool) bool {
	for _, s := range ss {
		if !isValid(s) {
			return false
		}
	}
	return true
}

// Greatest Common Divisor (GCD) - Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// Least Common Multiple (LCM) via GCD
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func ParseInts(str string) []int {
	nrStrList := strings.Fields(str)
	nrs := []int{}
	for _, str := range nrStrList {
		nr, _ := strconv.Atoi(str)
		nrs = append(nrs, nr)
	}
	return nrs
}

func Contains[T comparable](arr []T, x T) bool {
	for _, item := range arr {
		if item == x {
			return true
		}
	}
	return false
}
