package handler

import (
	"fmt"
	"os"
	"regexp"
)

// HandleInput processes command-line input arguments and performs actions based on them.
// It handles options such as setting output file, color, and displaying ASCII art.
func HandleInput() {
    // Regular expressions to parse command-line arguments
	IsOutput := regexp.MustCompile(`^--output=((.*).txt)$`)
	Output := IsOutput.FindStringSubmatch(os.Args[1])
	IsColor := regexp.MustCompile(`^--color=(?i)(Red|Blue|Yellow|Green|Magenta|Cyan|Gray|orange)$`)
	Color := IsColor.FindStringSubmatch(os.Args[1])

    // Switch statement based on the number of command-line arguments
	switch len(os.Args) {
	case 5:
        // Handling various conditions for option validation and usage display
		if CheckFlag(os.Args[2]) || CheckFlag(os.Args[3]) || CheckFlag(os.Args[4]) || (Output == nil && Color == nil) || (Output != nil) {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> something")
			os.Exit(0)
		}
		Check(os.Args[3])
		cont := ReadFile(os.Args[4])
		AsciiArtColor(os.Args[3], cont)
	case 4:
        // Handling various conditions for option validation and usage display
		if CheckFlag(os.Args[2]) || CheckFlag(os.Args[3]) || (Output == nil && Color == nil) {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> something")
			os.Exit(0)
		} else if Output == nil {
			Check(os.Args[3])
			cont := ReadFile("standard")
			AsciiArtColor(os.Args[3], cont)
		} else {
			Check(os.Args[2])
			cont := ReadFile(os.Args[3])
			AsciiArtOutput(os.Args[2], cont)
		}

	case 3:
        // Handling various conditions for option validation and usage display
		if CheckFlag(os.Args[2])  {
            fmt.Println("Usage: go run. [OPTION] [STRING]\n\nEX: go run. --color=<color> <letters to be colored> something")
            os.Exit(0)
        }
		if Color != nil {
			cont := ReadFile("standard")
			AsciiArtColor(os.Args[2], cont)
		} else if Output != nil && !CheckFlag(os.Args[2]){
			Check(os.Args[2])
			cont := ReadFile("standard")
			AsciiArtOutput(os.Args[2], cont)
		} else {
			Check(os.Args[1])
			cont := ReadFile(os.Args[2])
			AsciiArt(os.Args[1], cont)
		}
	default:
        // Handling various conditions for option validation and usage display
		if Output == nil && Color == nil {
			Check(os.Args[1])
			cont := ReadFile("standard")
			AsciiArt(os.Args[1], cont)
		} else if Output != nil || Color != nil {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> something")
			os.Exit(0)
		}
	}
}

