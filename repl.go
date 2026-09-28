package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokeapi "github.com/larrymacrich/pokedexcli/internal"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	commands map[string]cliCommand
	next     *string
	previous *string
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
		"map": {
			name: "map",
			description: ("Displays the names of location areas\n" +
				"Each subsequent call displays the next 20 areas"),
			callback: commandMap,
		},
		"mapb": {
			name: "map",
			description: ("Displays the names of location areas\n" +
				"Each subsequent call displays the previous 20 areas"),
			callback: commandMapb,
		},
	}
	return commands
}

// commandExit prints a farewell message and exits the program.
func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// commandHelp prints usage information for available commands.
func commandHelp(cfg *config) error {
	message := `Welcome to the Pokedex!
Usage:
`
	fmt.Println(message)
	commands := cfg.commands
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

// commandMap prints the names of the next 20 location areas
func commandMap(cfg *config) error {
	if cfg.next == nil {
		fmt.Println("you're on the last page")
		return nil
	}
	return fetchAndDisplayLocations(cfg, *cfg.next)
}

// commandMapb prints the names of the previous 20 location areas
func commandMapb(cfg *config) error {
	if cfg.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	return fetchAndDisplayLocations(cfg, *cfg.previous)
}

func fetchAndDisplayLocations(cfg *config, url string) error {
	locationAreasResponse, err := pokeapi.GetLocationAreas(url)
	if err != nil {
		errMsg := fmt.Errorf("something went wrong: %s", err)
		return errMsg
	}

	cfg.next = locationAreasResponse.Next
	cfg.previous = locationAreasResponse.Previous

	for _, areaResult := range locationAreasResponse.Results {
		fmt.Println(areaResult.Name)
	}
	return nil
}

// startRepl loops infinitly through CLI commands
func startRepl(cfg *config) {
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
		command, exists := cfg.commands[input[0]]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		// Check callback for errors
		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}

	}
	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}

// cleanInput normalizes input and splits it into words.
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
