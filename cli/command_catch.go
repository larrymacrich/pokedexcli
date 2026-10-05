package cli

import (
	"fmt"
	"math/rand"
)

// commandCatch tries to succesfully catch a pokemon given a chance
func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		errMsg := fmt.Errorf("required parameter <pokemon_name> missing")
		return errMsg
	}

	pokemonResponse, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		errMsg := fmt.Errorf("something went wrong: %s", err)
		return errMsg
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonResponse.Name)
	if isCaught(pokemonResponse.BaseExp) {
		fmt.Printf("%s was caught!\n", pokemonResponse.Name)
		cfg.pokedex.Add(pokemonResponse.Name, pokemonResponse)
	} else {
		fmt.Printf("%s escaped!\n", pokemonResponse.Name)
	}

	return nil
}

// isCaught calculates success rate catching a pokemon
func isCaught(exp int) bool {
	randomNum := rand.Intn(400)
	return exp < randomNum
}
