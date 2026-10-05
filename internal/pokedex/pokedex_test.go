package pokedex

import (
	"testing"
	"time"

	"github.com/larrymacrich/pokedexcli/internal/pokeapi"
)

func TestAddAndGet(t *testing.T) {
	const cacheInterval = 5 * time.Second
	const waitTime = 200 * time.Millisecond
	pokeapiClient := pokeapi.NewClient(cacheInterval)
	pokemonResponse, err := pokeapiClient.GetPokemon("pikachu")
	if err != nil {
		t.Fatalf("GetPokemon failed: %v", err)
	}
	time.Sleep(waitTime)
	p := NewPokedex()
	p.Add("pikachu", pokemonResponse)

	pokemonResponse, ok := p.Get("pikachu")
	if !ok {
		t.Fatalf("expected pikachu to be found")
	}
	if pokemonResponse.Name != "pikachu" {
		t.Fatalf("expected name 'pikachu', got %s", pokemonResponse.Name)
	}
}

func TestGetMissing(t *testing.T) {
	p := NewPokedex()
	_, ok := p.Get("nonexistent")
	if ok {
		t.Fatalf("expected nonexistent pokemon to not be found")
	}
}
