// https://adventofcode.com/2024/day/9

package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"time"
)

const aocDay int = 9

//  ____            _     _
// |  _ \ __ _ _ __| |_  / |
// | |_) / _` | '__| __| | |
// |  __/ (_| | |  | |_  | |
// |_|   \__,_|_|   \__| |_|

func First(lines []string) (strigifiedResult string) {
	result := 0

	line := lines[0]

	memory := []string{}

	id := 0
	for i := 0; i < len(line); i++ {
		digit, _ := strconv.Atoi(string(line[i]))
		for j := 0; j < digit; j++ {
			var strToAppend string
			if i%2 == 0 {
				strToAppend = strconv.Itoa(id)
			} else {
				strToAppend = "."
			}
			memory = append(memory, strToAppend)
		}

		if i%2 == 1 {
			id++
		}
	}

	endIndex := len(memory) - 1
	for i := 0; i < len(memory) && i < endIndex; i++ {
		if memory[i] == "." {
			for endIndex > i && memory[endIndex] == "." {
				endIndex--
			}

			memory[i] = memory[endIndex]
			memory[endIndex] = "."
			endIndex--
		}
	}

	for i := 0; i <= endIndex; i++ {
		nr, _ := strconv.Atoi(memory[i])
		result += i * nr
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// 90165373410 not correct
// correct: 6323641412437, example: 1928

//  ____            _     ____
// |  _ \ __ _ _ __| |_  |___ \
// | |_) / _` | '__| __|   __) |
// |  __/ (_| | |  | |_   / __/
// |_|   \__,_|_|   \__| |_____|

type file struct {
	fileId     int
	spaceAfter int
	fileLength int
	startIndex int
	nextFileId int
	prevFileId int
}

func (f file) String() string {
	return fmt.Sprintf("fileId: %d, spaceAfter: %d, fileLength: %d, startIndex: %d, nextFileId: %d, prevFileId: %d", f.fileId, f.spaceAfter, f.fileLength, f.startIndex, f.nextFileId, f.prevFileId)
}

func Second(lines []string) (strigifiedResult string) {
	result := 0

	line := lines[0]

	memory := make(map[int]*file)

	fileCount := 0
	memoryIndex := 0
	for i := 0; i < len(line); i++ {
		digit, _ := strconv.Atoi(string(line[i]))
		var fileToAppend file
		if i%2 == 0 {
			fileToAppend = file{fileCount, 0, digit, memoryIndex, -1, -1}

			if fileCount > 0 {
				memory[fileCount-1].nextFileId = fileCount
				fileToAppend.prevFileId = fileCount - 1
			}

			memory[fileCount] = &fileToAppend

			fileCount++
		} else {
			memory[fileCount-1].spaceAfter = digit
		}
		memoryIndex += digit
	}

	firstFileId := 0

	// do not move first file
	for rightmostFileId := fileCount - 1; rightmostFileId >= 1; rightmostFileId-- {
		leftId := firstFileId
		moved := false
		for !moved && leftId != rightmostFileId {
			leftFile := memory[leftId]
			rightFile := memory[rightmostFileId]

			if leftFile.spaceAfter >= rightFile.fileLength {
				// update linkings
				if _, ok := memory[rightFile.prevFileId]; ok {
					memory[rightFile.prevFileId].nextFileId = rightFile.nextFileId

					if _, ok := memory[rightFile.nextFileId]; ok {
						memory[rightFile.nextFileId].prevFileId = rightFile.prevFileId
					}

					memory[rightFile.prevFileId].spaceAfter += rightFile.fileLength + rightFile.spaceAfter
				}
				rightFile.prevFileId = leftFile.fileId

				// update props
				rightFile.spaceAfter = leftFile.spaceAfter - rightFile.fileLength
				rightFile.startIndex = leftFile.startIndex + leftFile.fileLength
				leftFile.spaceAfter = 0

				// update some more linkings (order matters)
				rightFile.nextFileId = leftFile.nextFileId

				if _, ok := memory[leftFile.nextFileId]; ok {
					memory[leftFile.nextFileId].prevFileId = rightmostFileId
				}
				leftFile.nextFileId = rightmostFileId

				leftId = rightmostFileId
				moved = true
			} else {
				leftId = leftFile.nextFileId
			}
		}
	}

	processedFiles := 0
	currentFileId := 0
	index := 0
	for processedFiles < fileCount {
		currentFile := memory[currentFileId]
		for i := 0; i < currentFile.fileLength; i++ {
			result += index * currentFileId
			index++
		}
		for i := 0; i < currentFile.spaceAfter; i++ {
			index++
		}
		currentFileId = currentFile.nextFileId
		processedFiles++
	}

	strigifiedResult = fmt.Sprint(result)
	return strigifiedResult
}

// 85613407812 too low
// correct: 6351801932670, example 2858

func main() {
	startTime := time.Now()

	inputFileExtension := utils.GetInputFileExt(1)

	lines := utils.ReadLines(fmt.Sprintf("%d.%s", aocDay, inputFileExtension))

	fmt.Println("--- First ---")
	result := First(lines)
	fmt.Println(result)

	fmt.Println("\n--- Second ---")
	result = Second(lines)
	fmt.Println(result)

	programDuration := time.Since(startTime).Seconds()

	fmt.Printf("\n✨ Finished in %.3f seconds\n", programDuration)
}
