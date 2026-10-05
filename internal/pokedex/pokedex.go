package pokedex

import "sync"

type Pokedex struct {
	mu       sync.Mutex
	pokemons map[string]PokedexEntry
}

type PokedexEntry struct {
	name string
	exp  int
}

func NewPokedex() *Pokedex {
	pokemons := make(map[string]PokedexEntry)
	return &Pokedex{
		mu:       sync.Mutex{},
		pokemons: pokemons,
	}
}

func (p *Pokedex) Add(name string, exp int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	entry := PokedexEntry{
		name: name,
		exp:  exp,
	}
	p.pokemons[name] = entry
}

func (p *Pokedex) Get(name string) (string, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	entry, ok := p.pokemons[name]
	if !ok {
		return "", ok
	}
	return entry.name, ok
}
