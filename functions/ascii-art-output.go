package handler

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// AsciiArtOutput generates ASCII art and writes it to a file based on the provided arguments.
// It splits the content and argument, creates or truncates the output file, and writes the ASCII art to it.
func AsciiArtOutput(arg string, cont string) {
	SplitFile, SplitArg, result := SplitCont(arg, cont)
	
	// Regular expression to check the output file flag
	CheckFlag := regexp.MustCompile(`^--output=((.*).txt)$`)
	Flag := CheckFlag.FindStringSubmatch(os.Args[1])
	
	// Print usage and exit if output file flag is invalid
	if Flag == nil {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	}

	// Create or truncate the output file
	file, err := os.Create(Flag[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer file.Close()
	
	// Handle new line insertion if required
	if NewLineCheck(os.Args[2]) {
		for i := 0; i < len(os.Args[2])/2; i++ {
			file.WriteString(string("\n"))
		}
		return
	}
	
	// Iterate over each argument segment
	for i := 0; i < len(SplitArg); i++ {
		// Iterate over each line of ASCII art
		for l := 0; l < 8; l++ {
			// Iterate over each character in the argument
			for j := 0; j < len(SplitArg[i]); j++ {
				// Split characters of ASCII art
				SplitCharacters := strings.Split(string(SplitFile[SplitArg[i][j]-32]), "\n")
				if SplitCharacters[l] == "" {
					result[i][l] += "      "
				}
				result[i][l] += SplitCharacters[l]
			}
			// Write the result line by line to the output file
			if result[i][l] != "" {
				file.WriteString(string(result[i][l] + "\n"))
			} else {
				file.WriteString(string("\n"))
				break
			}
		}
	}
}
