package cli

// commandPokedex prints a list of all caught pokemon to the console
func commandPokedex(cfg *config, args []string) error {
	cfg.pokedex.PrintPokedex()
	return nil
}
