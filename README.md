# Advent of Code repo - Bertalan Vad

Hello and welcome! In this repo I gather all my solutions to all the AoC editions.

[[_TOC_]]

# Setup

## Session for aocd package

Session token needed for the aoc packages to work: `export AOC_SESSION=ABCD1324`, get that from browser cookies. (More info: https://github.com/wimglenn/advent-of-code-wim/issues/1)

## Solving with Python

The package also introspects the current day from the path/name of the script that contains the solution, so be careful with namings: `10_1_asd.py` is not good, use `10_a_asd.py` instead.

## Solving with Go

### Start to solve a day's exercise

You can setup an initial workspace for Go for the current day by running `./setupgo.sh` and then navigate into the created directory. Be sure that the file has the right permissions: `chmod 777 setupgo.sh`.

#### Example Usage of `setupgo.sh`

To set up the workspace for a specific year and day, you can use the following command:

```sh
./setupgo.sh -y 2023 -d 1
```

If you want to set up the workspace for the current year and day, simply run (works only if the month is December):

```sh
./setupgo.sh
```

### Running with the framework

The Go framework uses command line flags to determine which input file to use. Here are the available flags:

- `-r`: Use the real personalized input file (`XX.in`). By default (if you omit this flag), the example provided in the AoC problem description is used.
- `-e2`, `-e3`, ..., `-eN`: Use example input number 2, 3, ..., N (`XX.exin2`, `XX.exin3`, ..., `XX.exinN`). By default, the simpler example provided in the AoC problem description is used.

#### Example Usage

To run the solution against the real input file:

```sh
go run 25.go -r
```

To run the solution against the second example input file:

```sh
go run 25.go -e2
```

To run the solution against the default example input file:

```sh
go run 25.go
```

### Result-test generation

This utility generates a Go test file for a specific Advent of Code challenge day and year, with provided answers for the first and second parts of the challenge, including example answers.

#### Usage: 

`go run utils/gen/gentest.go <year> <day> <first_example_answer> <second_example_answer> <first_answer> <second_answer>`

#### Example of usage:

`go run utils/gen/gentest.go 2023 01 123 456 12345 67890`

The generated test file will be created in a directory named after the specified year, with the filename format `<day>\_test.go`. The test file will contain the provided answers embedded in a template.
