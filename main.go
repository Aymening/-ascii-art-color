package main

import (
	"fmt"
	"os"

	"handler/functions"
)

func main() {
	if len(os.Args) > 5 ||  len(os.Args) == 1 {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> something")
		return
	}
	handler.HandleInput()
}
