package cli

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

// getCliCommands returns the commands supported by the CLI.
// TODO: change type to get an orderd list
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
			description: ("Attempting to catch a pokemon"),
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: ("Printing the name, height, weight, stats and type(s) of the pokemon to the console"),
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: ("Prints a list of all caught pokemon to the console"),
			callback:    commandPokedex,
		},
	}
	return commands
}
