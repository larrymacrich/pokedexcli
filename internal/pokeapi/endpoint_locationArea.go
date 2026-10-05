package pokeapi

type LocationAreaResponse struct {
	Id                int                `json:"id"`
	Name              string             `json:"name"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon PokemonArea `json:"pokemon"`
}

type PokemonArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
