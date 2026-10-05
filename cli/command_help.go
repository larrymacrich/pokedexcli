package cli

import "fmt"

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
