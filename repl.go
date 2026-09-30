package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	pokeapi "github.com/larrymacrich/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
	parameters  map[string]string
}

type config struct {
	commands      map[string]cliCommand
	next          *string
	previous      *string
	cacheInterval time.Duration
	pokeapiClient *pokeapi.Client
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
		"explore": {
			name:        "explore",
			description: ("Displays the names of pokemon found in a given area"),
			callback:    commandExplore,
		},
	}
	return commands
}

// commandExit prints a farewell message and exits the program.
func commandExit(cfg *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// commandHelp prints usage information for available commands.
func commandHelp(cfg *config, args []string) error {
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
func commandMap(cfg *config, args []string) error {
	if cfg.next == nil {
		fmt.Println("you're on the last page")
		return nil
	}
	return fetchAndDisplayLocations(cfg, *cfg.next)
}

// commandMapb prints the names of the previous 20 location areas
func commandMapb(cfg *config, args []string) error {
	if cfg.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	return fetchAndDisplayLocations(cfg, *cfg.previous)
}

func fetchAndDisplayLocations(cfg *config, url string) error {
	locationAreasResponse, err := cfg.pokeapiClient.GetLocationAreas(url)
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

// commandExplore prints the available pokemon in a given area
func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		errMsg := fmt.Errorf("required parameter <area_name> missing")
		return errMsg
	}

	locationAreaResponse, err := cfg.pokeapiClient.GetLocationArea(args[0])
	if err != nil {
		errMsg := fmt.Errorf("something went wrong: %s", err)
		return errMsg
	}

	fmt.Printf("Exploring %s...\nFound Pokemon:\n", locationAreaResponse.Name)
	for _, pokemonEncounter := range locationAreaResponse.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemonEncounter.Pokemon.Name)
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

		// Check for valid commands and parameters
		command, exists := cfg.commands[input[0]]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		// Check callback for errors
		err := command.callback(cfg, input[1:])
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
