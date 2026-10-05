package pokeapi

type PokemonResponse struct {
	Name    string        `json:"name"`
	BaseExp int           `json:"base_experience"`
	Height  int           `json:"height"`
	Weight  int           `json:"weight"`
	Stats   []PokemonStat `json:"stats"`
	Types   []PokemonType `json:"types"`
}

type PokemonStat struct {
	BaseStat int  `json:"base_stat"`
	Stat     Stat `json:"stat"`
}

type Stat struct {
	Name string `json:"name"`
}

type PokemonType struct {
	PType PType `json:"type"`
}

type PType struct {
	Name string `json:"name"`
}
