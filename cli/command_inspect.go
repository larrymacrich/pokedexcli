package cli

import (
	"fmt"

	"github.com/larrymacrich/pokedexcli/internal/pokeapi"
)

// commandInspect prints the name, height, weight, stats and type(s) of the pokemon
func commandInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		errMsg := fmt.Errorf("required parameter <pokemon_name> missing")
		return errMsg
	}
	pokemon := args[0]
	pokemonResponse, ok := cfg.pokedex.Get(pokemon)
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	printPokemon(pokemonResponse)
	return nil

}

func printPokemon(pokemon *pokeapi.PokemonResponse) {
	fmt.Printf(`Name: %s
Height: %v
Weight: %v
`,
		pokemon.Name,
		pokemon.Height,
		pokemon.Weight)

	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf("  - %s: %v\n", s.Stat.Name, s.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.PType.Name)
	}

}
