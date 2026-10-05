package cli

import "fmt"

// commandExplore prints the available pokemon in a given area
func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		errMsg := fmt.Errorf("required parameter <area_name> missing")
		return errMsg
	}

	locationAreaResponse, err := cfg.pokeapiClient.GetLocationArea(args[0])
	// TODO: add proper error handling, e.g.
	// 404 Not found => invalid Pokemon name
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
