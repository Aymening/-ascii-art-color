package handler

import "strings"

// NewLineCheck checks if the provided string consists only of newline characters.
// It replaces escape sequences representing newlines with actual newline characters,
// then counts the occurrences of newline characters in the string.
// Returns true if the string contains only newline characters, otherwise false.
func NewLineCheck(str string) bool {
	cont := 0
	result := strings.ReplaceAll(str, "\\n", "\n")
	for i := 0; i < len(result); i++ {
		if result[i] == '\n' {
			cont++
		}
	}
	return cont == len(result) 
}
