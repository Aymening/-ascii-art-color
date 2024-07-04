package handler

import (
	"bufio"
	"fmt"
	"os"
)

// CheckFileLines checks if the number of lines in the specified file matches the expected count.
// It opens the file and counts the number of lines. If the count matches the expected value (855),
// it returns true; otherwise, it prints an error message and exits the program, returning false.
func CheckFileLines(filename string) bool {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("use this Banner (standard or shadow or benso or thinkertoy)")
		os.Exit(0)
	}
	// Create a scanner to read lines from the file
	scan := bufio.NewScanner(file)
	con := 0
	// Count the number of lines in the file
	for scan.Scan() {
		con++
	}
	// Compare the line count with the expected value (855)
	return con == 855 

}
