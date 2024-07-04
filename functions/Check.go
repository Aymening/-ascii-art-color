package handler

import (
	"fmt"
	"os"
)

// Check checks if the provided string contains valid characters within the ASCII range [32, 126].
// If any character falls outside this range, it prints an error message and exits the program.
func Check(str string) {
	for i := 0; i < len(str); i++ {
		if str[i] < 32 || str[i] > 126 {
			fmt.Println("PROGRAM : Invalid characters")
			os.Exit(0)
		}
	}
}
