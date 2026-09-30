package main

import (
	"time"

	"github.com/larrymacrich/pokedexcli/internal/pokeapi"
)

func main() {
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	cacheInterval := 15 * time.Second
	cfg := config{
		commands:      getCliCommands(),
		next:          &baseURL,
		cacheInterval: cacheInterval,
		pokeapiClient: pokeapi.NewClient(cacheInterval),
	}
	startRepl(&cfg)
}
