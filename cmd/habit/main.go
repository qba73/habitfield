package main

import (
	"bytes"
	"habit"
	"os"
)

func main() {
	userInput := os.Args[1:]
	writer := &bytes.Buffer{}
	habit.ProcessUserInput(userInput, writer)
}
