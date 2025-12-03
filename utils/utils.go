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

func WriteToFile(path string, content string) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err := f.WriteString(content); err != nil {
		panic(err)
	}
}

func AppendLineToFile(path string, line string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.WriteString(line + "\n"); err != nil {
		panic(err)
	}
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
func GetInputFileExt(inputFileCount int) string {
	useRealInputP := flag.Bool("r", false, "Specify if you want to run the solution against the real personalized input, which should be put in file `XX.in` beforehand. By default the example provided in the AoC problem description is used.")

	useExPs := []*bool{}
	if inputFileCount > 1 {
		for i := 0; i < inputFileCount-1; i++ {
			a := fmt.Sprintf("e%d", i+2)
			useExInputI := flag.Bool(a, false, fmt.Sprintf("Specify if you want to run the solution against example input number %d, which should be put in file `XX.exin%d`. By default the simpler example provided in the AoC problem description is used.", i+2, i+2))
			useExPs = append(useExPs, useExInputI)
		}
	}

	flag.Parse()

	var inputFileExtension string
	if *useRealInputP {
		inputFileExtension = "in"
	} else {
		useEx := -1
		for i, useExP := range useExPs {
			if *useExP {
				useEx = i + 2
			}
		}

		if useEx != -1 {
			fmt.Printf("Using example %d input", useEx)
			fmt.Println("-------------------")
			inputFileExtension = "exin" + fmt.Sprint(useEx)
		} else {
			fmt.Println("Using example input")
			fmt.Println("-------------------")
			inputFileExtension = "exin"
		}
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

// TODO: rename to MapValues
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

func AllTrue(arr []bool) bool {
	for _, v := range arr {
		if !v {
			return false
		}
	}
	return true
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

func Map[T any, U any](ss []T, transform func(T) U) (ret []U) {
	for _, s := range ss {
		ret = append(ret, transform(s))
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

func Some[T any](ss []T, isValid func(T) bool) bool {
	for _, s := range ss {
		if isValid(s) {
			return true
		}
	}
	return false
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

// Least Common Multiple (LCM) via GCD for an array
func LCMArr(ints []int) int {
	ret := 1
	for _, v := range ints {
		ret = LCM(ret, v)
	}
	return ret
}

func ParseInts(str string, sep string) []int {
	nrStrList := strings.Split(str, sep)
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

func Find[T any](arr []T, predicate func(T) bool) (item *T, index int) {
	for i := 0; i < len(arr); i++ {
		if predicate(arr[i]) {
			return &arr[i], i
		}
	}
	return nil, -1
}

func RemoveOnIndexOrderPreserved[T any](arr []T, index int) []T {
	return append(arr[:index], arr[index+1:]...)
}

func RemoveOnIndex[T any](arr []T, index int) []T {
	arr[index] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}

func RemoveItemFromArray[T comparable](arr []T, item T) []T {
	for i := 0; i < len(arr); i++ {
		if arr[i] == item {
			return RemoveOnIndex(arr, i)
		}
	}
	return arr
}

func Sum(arr []int) (sum int) {
	for _, v := range arr {
		sum += v
	}
	return sum
}

func GetIndices(in string, of string) (ret []int) {
	for i := 0; i < len(in); i++ {
		if string(in[i]) == of {
			ret = append(ret, i)
		}
	}
	return ret
}

func ShallowCopyMap[K comparable, V any](m map[K]V) (newMap map[K]V) {
	newMap = map[K]V{}
	for k, v := range m {
		newMap[k] = v
	}
	return newMap
}

func CloneArray[T any](arr []T) []T {
	newArr := make([]T, len(arr))
	copy(newArr, arr)
	return newArr
}

func Abs(a int) int {
	return int(math.Abs(float64(a)))
}

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func FilterDuplicates[T comparable](arr []T) []T {
	keys := make(map[T]bool)
	list := []T{}
	for _, entry := range arr {
		if _, ok := keys[entry]; !ok {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func GroupBy[T any, K comparable](arr []T, keySelector func(T) K) map[K][]T {
	grouped := map[K][]T{}
	for _, item := range arr {
		key := keySelector(item)
		grouped[key] = append(grouped[key], item)
	}
	return grouped
}

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
