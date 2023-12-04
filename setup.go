package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
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

func WebInput(year, day int) []byte {
	// Fetch from Advent of Code website
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	httpClient := http.Client{Timeout: time.Duration(3) * time.Second}

	session := &http.Cookie{
		Name:   "session",
		Value:  os.Getenv("AOC_SESSION"),
		MaxAge: 0,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error with creating the HTTP request object: %s", err.Error())
	}
	req.AddCookie(session)

	req.Header.Set("User-Agent", "vadbertalan@yahoo.com")

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(fmt.Sprintf("Error: %s", err))
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic("Error received from server: (possibly could not authenticate with AOC server)")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("Error: %s", err))
	}
	return body
}

func main() {
	now := time.Now()
	currentYear := now.Year()
	currentDay := now.Day()

	fmt.Printf("Hello! ☀️\nSetting up Go workspace for day %d. GL! 🤙\n\n", currentDay)

	currentDayFolderName := fmt.Sprint(currentDay)

	// Pad the folder name in case the day is <= 9
	if len(currentDayFolderName) == 1 {
		currentDayFolderName = fmt.Sprintf("0%s", currentDayFolderName)
	}

	newFolder := fmt.Sprintf("%d/%s", currentYear, currentDayFolderName)

	// Create dir for current day
	if err := os.Mkdir(newFolder, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created folder for today's problem, %s, make sure to navigate into it:\n\ncd %s\n\n", currentDayFolderName, newFolder)

	// Create files

	goFileName := fmt.Sprintf("%s/%d.go", newFolder, currentDay)
	inExampleFileName := fmt.Sprintf("%s/%d.exin", newFolder, currentDay)
	inFileName := fmt.Sprintf("%s/%d.in", newFolder, currentDay)

	// Create Go src file
	createGoFileFromSchema("schema.go", goFileName, currentDay)

	// Create input file for quick example
	err := os.WriteFile(inExampleFileName, []byte("\n"), os.ModePerm)
	checkErr(err)

	// Create input file
	input := WebInput(currentYear, currentDay)
	err = os.WriteFile(inFileName, []byte(fmt.Sprintf("%s", input)), os.ModePerm)
	checkErr(err)

	fmt.Printf("Created files:\n- %s\n- %s\n- %s\n\n", goFileName, inExampleFileName, inFileName)
}
