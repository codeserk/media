package main

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

var path = "/Users/codeserk/Downloads/ol_dump_editions_2025-04-30.txt.gz"

func main() {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	gzReader, err := gzip.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}
	defer gzReader.Close()

	scanner := bufio.NewScanner(gzReader)
	i := 0
	var titles []string
	for i < 1 && scanner.Scan() {
		newTitle := processLine(scanner.Text())
		titles = append(titles, newTitle)
		i++
	}

	fmt.Printf("titles %v", len(titles))
	for _, title := range titles {
		if strings.Contains(strings.ToLower(title), "hamlet") {
			fmt.Printf("Found title containing 'hamlet': %s\n", title)
		}
	}
}

func processLine(line string) string {
	fmt.Printf("%v", line)

	// Split the line into columns
	columns := strings.Split(line, "\t")
	if len(columns) != 5 {
		log.Fatalf("Invalid line format: %s", line)
	}

	// Extract the JSON part
	jsonPart := columns[4]

	// Unmarshal the JSON part into the Work struct
	var work Work
	err := json.Unmarshal([]byte(jsonPart), &work)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Print the extracted information
	// fmt.Printf("Type: %s\n", columns[0])
	// fmt.Printf("ID: %s\n", columns[1])
	// fmt.Printf("Version: %s\n", columns[2])
	// fmt.Printf("Date: %s\n", columns[3])
	// fmt.Printf("Title: %s\n", work.Title)
	// fmt.Printf("Created: %s\n", work.Created.Value)
	// fmt.Printf("Last Modified: %s\n", work.LastModified.Value)
	// fmt.Printf("Key: %s\n", work.Title)
	// fmt.Printf("Authors: %v\n", work.Authors)

	return work.Title
}

type Work struct {
	Title string `json:"title"`
	// Created      DateTime     `json:"created"`
	// LastModified DateTime     `json:"last_modified"`
	// Key          string       `json:"key"`
	// Authors      []AuthorRole `json:"authors"`
}

type DateTime struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type AuthorRole struct {
	Type   string `json:"type"`
	Author Author `json:"author"`
}

type Author struct {
	Key string `json:"key"`
}
