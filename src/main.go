package main

import (
	"bufio"
	"fmt"
	"os"
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

		// Print back the user's input
		fmt.Println(input)
	}
}
