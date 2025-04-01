package main

import (
	"os"
	"fmt"
	"log"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No file provided")
	}

	// Get the file path from the command line arguments
	filePath := os.Args[1]

	// Check if the file path is provided
	if filePath == "" {
		log.Fatal("No file provided")
	}

	// Get the file contents
	data, err := readFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	// Print the file contents
	fmt.Println(string(data))
}