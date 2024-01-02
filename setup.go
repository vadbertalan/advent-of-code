package main

import (
	"flag"
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

func createGoFileFromSchema(src string, dst string, currentYear, currentDay int) {
	// Read all content of src to data, may cause OOM for a large file.
	data, err := os.ReadFile(src)
	checkErr(err)

	dataStr := string(data)

	// Substitute current day variable dynamically
	dataStr = strings.Replace(dataStr, "aocDay int = 999", fmt.Sprintf("aocDay int = %d", currentDay), 1)
	dataStr = strings.Replace(dataStr, "yyyy/day/dd", fmt.Sprintf("%d/day/%d", currentYear, currentDay), 1)
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

	yearInputP := flag.Int("y", currentYear, "Specify to tell the bootstrapper script the year in which you want to generate files for a day.")
	dayInputP := flag.Int("d", currentDay, "Specify to tell the bootstrapper script the day you want to generate files for.")
	flag.Parse()

	fmt.Printf("Hello, AoC warrior! ☀️\nSetting up Go workspace for day %d, %d. GL! 🤙\n\n", *dayInputP, *yearInputP)

	currentDayFolderName := fmt.Sprint(*dayInputP)

	// Pad the folder name in case the day is <= 9
	if len(currentDayFolderName) == 1 {
		currentDayFolderName = fmt.Sprintf("0%s", currentDayFolderName)
	}

	// Create dir for year if it does not exist yet
	newFolderForYear := fmt.Sprint(*yearInputP)
	if err := os.MkdirAll(newFolderForYear, os.ModePerm); err != nil {
		log.Fatal("Error while creating folder for year: ", err)
	}

	// Create dir for current day. If dir for day already exists, fails.
	newFolderForDay := fmt.Sprintf("%s/%s", newFolderForYear, currentDayFolderName)
	if err := os.Mkdir(newFolderForDay, os.ModePerm); err != nil {
		log.Fatal("Error while creating folder for day: ", err)
	}

	fmt.Printf("Created folder for problem %s, make sure to navigate into it:\n\ncd %s\n\n", currentDayFolderName, newFolderForDay)

	// Create files

	goFileName := fmt.Sprintf("%s/%d.go", newFolderForDay, *dayInputP)
	inExampleFileName := fmt.Sprintf("%s/%d.exin", newFolderForDay, *dayInputP)
	inFileName := fmt.Sprintf("%s/%d.in", newFolderForDay, *dayInputP)

	// Create Go src file
	createGoFileFromSchema("schema.go", goFileName, *yearInputP, *dayInputP)

	// Create input file for quick example
	err := os.WriteFile(inExampleFileName, []byte("\n"), os.ModePerm)
	checkErr(err)

	// Create input file
	input := WebInput(*yearInputP, *dayInputP)
	err = os.WriteFile(inFileName, []byte(fmt.Sprintf("%s", input)), os.ModePerm)
	checkErr(err)

	fmt.Printf("Created files:\n- %s\n- %s\n- %s\n\n", goFileName, inExampleFileName, inFileName)
}
