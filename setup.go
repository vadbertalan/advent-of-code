package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createGoFileFromSchema(src string, dst string, currentDay int) {
	// Read all content of src to data, may cause OOM for a large file.
	data, err := os.ReadFile(src)
	checkErr(err)

	dataStr := string(data)

	// Substitute current day variable dynamically
	dataStr = strings.Replace(dataStr, "aocDay int = 999", fmt.Sprintf("aocDay int = %d", currentDay), 1)
	// This is needed because, having main name would cause lint error
	dataStr = strings.Replace(dataStr, "DYNmain", "main", 1)

	data = []byte(dataStr)

	// Write data to dst
	err = os.WriteFile(dst, data, os.ModePerm)
	checkErr(err)
}

func main() {
	currentDay := time.Now().Day()

	fmt.Printf("Hello! ☀️\nSetting up Go workspace for day %d. GL! 🤙\n\n", currentDay)

	currentDayFolderName := fmt.Sprint(currentDay)

	// Pad the folder name in case the day is <= 9
	if len(currentDayFolderName) == 1 {
		currentDayFolderName = fmt.Sprintf("0%s", currentDayFolderName)
	}

	// Create dir for current day
	if err := os.Mkdir(currentDayFolderName, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created folder for today's problem, %s, make sure to navigate into it:\n\ncd %s\n\n", currentDayFolderName, currentDayFolderName)

	// Create files

	goFileName := fmt.Sprintf("%s/%d.go", currentDayFolderName, currentDay)
	inExampleFileName := fmt.Sprintf("%s/%d.exin", currentDayFolderName, currentDay)
	inFileName := fmt.Sprintf("%s/%d.in", currentDayFolderName, currentDay)

	// Create Go src file
	createGoFileFromSchema("schema.go", goFileName, currentDay)

	// Create input file for quick example
	err := os.WriteFile(inExampleFileName, []byte("\n"), os.ModePerm)
	checkErr(err)

	// Create input file
	// TODO: add current day's input with aocgen tool: https://github.com/timkelleher/aocgen
	err = os.WriteFile(inFileName, []byte("\n"), os.ModePerm)
	checkErr(err)

	fmt.Printf("Create files:\n- %s/%s\n- %s/%s\n- %s/%s\n\n", currentDayFolderName, goFileName, currentDayFolderName, inExampleFileName, currentDayFolderName, inFileName)
}
