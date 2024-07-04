package handler

import (
	"fmt"
	"os"
	"strings"
)

// AsciiArt generates ASCII art based on the provided arguments.
// It splits the content and argument, and prints the ASCII art.
func AsciiArt(arg string, cont string) {
	// Split the content and argument
	SplitFile, SplitArg, result := SplitCont(arg, cont)
	// Handle new line insertion if required
	if NewLineCheck(os.Args[1]) {
		for i := 0; i < len(os.Args[1])/2; i++ {
			fmt.Println()
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
				// Append characters to result
				if SplitCharacters[l] == "" {
					result[i][l] += "      "
				}
				result[i][l] += SplitCharacters[l]
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
