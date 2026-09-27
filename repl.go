package main

import (
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// getCliCommands returns the commands supported by the CLI.
func getCliCommands() map[string]cliCommand {
	var commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
	return commands
}

// commandExit prints a farewell message and exits the program.
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// commandHelp prints usage information for available commands.
func commandHelp() error {
	message := `Welcome to the Pokedex!
Usage:
`
	fmt.Println(message)
	commands := getCliCommands()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

// cleanInput normalizes input and splits it into words.
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
