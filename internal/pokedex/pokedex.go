package pokedex

import (
	"fmt"
	"sync"

	"github.com/larrymacrich/pokedexcli/internal/pokeapi"
)

type Pokedex struct {
	mu       sync.Mutex
	pokemons map[string]*pokeapi.PokemonResponse
}

func NewPokedex() *Pokedex {
	pokemons := make(map[string]*pokeapi.PokemonResponse)
	return &Pokedex{
		mu:       sync.Mutex{},
		pokemons: pokemons,
	}
}

func (p *Pokedex) Add(name string, pokemon *pokeapi.PokemonResponse) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.pokemons[name] = pokemon
}

func (p *Pokedex) Get(name string) (*pokeapi.PokemonResponse, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	entry, ok := p.pokemons[name]
	if !ok {
		return nil, ok
	}
	return entry, ok
}

func (p *Pokedex) PrintPokedex() {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range p.pokemons {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
}
