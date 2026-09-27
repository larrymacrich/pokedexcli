package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Reading from console
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		fmt.Printf("Your command was: %s\n", input[0])
	}
	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
