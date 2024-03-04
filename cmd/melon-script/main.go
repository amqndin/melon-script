package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for true {
		fmt.Print("> ") // Print the prompt

		// Read user input until newline
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			break // Exit loop on error
		}

		// Remove trailing newline character
		// input = strings.TrimSuffix(input, "\n")

		// Print back the user's input
		fmt.Println(input)
	}
}
