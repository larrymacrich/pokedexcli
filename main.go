package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	commands := getCliCommands()
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

		// Check for valid commands
		command, exists := commands[input[0]]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		// Check callback for errors
		err := command.callback()
		if err != nil {
			fmt.Println(err)
		}

	}
	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
