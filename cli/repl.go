package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/larrymacrich/pokedexcli/internal/pokeapi"
	"github.com/larrymacrich/pokedexcli/internal/pokedex"
)

// startRepl loops infinitly through CLI commands
func StartRepl() {
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	cacheInterval := 15 * time.Second
	cfg := config{
		commands:      getCliCommands(),
		next:          &baseURL,
		cacheInterval: cacheInterval,
		pokeapiClient: pokeapi.NewClient(cacheInterval),
		pokedex:       pokedex.NewPokedex(),
	}
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
		err := command.callback(&cfg, input[1:])
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
