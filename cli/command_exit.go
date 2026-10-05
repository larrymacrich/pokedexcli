package cli

import (
	"fmt"
	"os"
)

// commandExit prints a farewell message and exits the program.
func commandExit(cfg *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
