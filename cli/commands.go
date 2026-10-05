package cli

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
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
		"mapf": {
			name: "mapf",
			description: ("Displays the names of location areas\n" +
				"Each subsequent call displays the next 20 areas"),
			callback: commandMapf,
		},
		"mapb": {
			name: "mapb",
			description: ("Displays the names of location areas\n" +
				"Each subsequent call displays the previous 20 areas"),
			callback: commandMapb,
		},
		"explore": {
			name:        "explore <area_name>",
			description: ("Displays the names of pokemon found in a given area"),
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: ("Attempting to catch the given pokemon"),
			callback:    commandCatch,
		},
	}
	return commands
}
