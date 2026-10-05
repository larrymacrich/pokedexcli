package cli

import (
	"time"

	"github.com/larrymacrich/pokedexcli/internal/pokeapi"
	"github.com/larrymacrich/pokedexcli/internal/pokedex"
)

type config struct {
	commands      map[string]cliCommand
	next          *string
	previous      *string
	cacheInterval time.Duration
	pokeapiClient *pokeapi.Client
	pokedex       *pokedex.Pokedex
}
