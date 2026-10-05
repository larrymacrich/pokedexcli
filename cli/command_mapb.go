package cli

import (
	"fmt"
)

// commandMapb prints the names of the previous 20 location areas
func commandMapb(cfg *config, args []string) error {
	if cfg.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	return fetchAndDisplayLocations(cfg, *cfg.previous)
}
