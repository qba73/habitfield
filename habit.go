package habit

import (
	"fmt"
	"io"
)

func ProcessUserInput(userInput []string, writer io.Writer) (int, error) {
	if userInput[0] == "habit" {
		if len(userInput) > 1 {
		for _, habit := range userInput[1:] {
			return fmt.Fprintf(writer, "%s", habit)
	}
}
	return fmt.Fprintf(writer, "Habit is a command line tool for tracking habits. To get started, type 'habit help'")
	}
	return fmt.Fprintf(writer, "%s is not a habit command", userInput[0])
}
