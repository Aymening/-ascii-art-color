package handler

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// AsciiArtColor generates ASCII art with colored letters based on the provided arguments.
// It splits the content and argument, searches for highlighted positions,
// and applies color to the corresponding characters. Finally, it prints the colored ASCII art.
func AsciiArtColor(arg string, cont string) {
	// Split the content and argument
	SplitFile, SplitArg, result := SplitCont(arg, cont)
	if NewLineCheck(arg) {
		for i := 0; i < len(arg)/2; i++ {
			fmt.Println()
		}
		return
	}
	
	// Regular expression to check the color flag
	CheckFlag := regexp.MustCompile(`^--color=(?i)(red|blue|yellow|green|magenta|cyan|gray|orange)$`)
	Flag := CheckFlag.FindStringSubmatch(os.Args[1])
	
	// Print usage and exit if color flag is invalid
	if Flag == nil {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> something")
		os.Exit(0)
	}
	
	// Iterate over each argument segment
	for i := 0; i < len(SplitArg); i++ {
		// Find positions to highlight
		highlightPositions := SearchSub(SplitArg[i], os.Args[2])
		// Iterate over each line of ASCII art
		for l := 0; l < 8; l++ {
			// Iterate over each character in the argument
			for j := 0; j < len(SplitArg[i]); j++ {
				// Split characters of ASCII art
				SplitCharacters := strings.Split(string(SplitFile[SplitArg[i][j]-32]), "\n")
				// Apply color if character position is highlighted
				if IsHere(j, highlightPositions) {
					if SplitCharacters[l] == "" {
						result[i][l] += "      "
					}
					Color(SplitCharacters, strings.ToLower(Flag[1]))
				}
				// Append characters to result
				if SplitCharacters[l] == "" {
					result[i][l] += "      "
				} else {
					result[i][l] += SplitCharacters[l]
				}
			}
			// Print the result line by line
			if result[i][l] != "" {
				fmt.Println(result[i][l])
			} else {
				fmt.Println()
				break
			}
		}
	}
}

