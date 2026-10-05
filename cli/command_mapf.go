package cli

import "fmt"

// commandMap prints the names of the next 20 location areas
func commandMapf(cfg *config, args []string) error {
	if cfg.next == nil {
		fmt.Println("you're on the last page")
		return nil
	}
	return fetchAndDisplayLocations(cfg, *cfg.next)
}
