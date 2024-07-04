package handler

import "strings"

// SplitCont splits the content and argument into segments for further processing.
// It splits the content based on double newline characters and the argument based on "\\n" escape sequence.
// Returns three slices: SplitFile (content segments), SplitArg (argument segments), and result (empty 2D slice for storing processed data).
func SplitCont(arg string, cont string) ([]string, []string, [][]string) {
	
	// Split the content into segments based on double newline characters
	SplitFile := strings.Split(cont, "\n\n")
	
	// Split the argument into segments based on "\\n" escape sequence
	SplitArg := strings.Split(arg, "\\n")
	
	// Initialize an empty 2D slice to store processed data
	result := make([][]string, len(SplitArg))
	for mak := 0; mak < len(SplitArg); mak++ {
		result[mak] = make([]string, 8)
	}
	return SplitFile, SplitArg, result
}
