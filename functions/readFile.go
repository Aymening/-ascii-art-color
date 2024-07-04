package handler

import (
	"fmt"
	"os"
	"strings"
)

// ReadStandardFile reads the content of a standard file located in the "Banner" directory.
// It takes the file name as input, appends ".txt" extension, and reads the file's content.
// If the file does not exist or cannot be read, it prints an error message and exits the program.
// Returns the content of the file as a string.
func ReadFile(FileName string) string {
	// Construct file path
	file := "Banner/" + FileName + ".txt"

	// Check if the file contains any invalid lines
	if !CheckFileLines(file) {
		os.Exit(0)
	}

	// Read the content of the file
	cont, err := os.ReadFile(file)
	if err != nil {
		// Print error message and exit if file cannot be read
		fmt.Println("PROGRAM : ", err)
		os.Exit(0)
	}

	// Replace carriage return characters with empty string
	modf := strings.NewReplacer(
		"\r", "",
	)
	result := modf.Replace(string(cont))

	// Return the content of the file
	return string(result)
}
