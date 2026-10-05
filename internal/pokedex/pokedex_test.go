package pokedex

import "testing"

func TestAddAndGet(t *testing.T) {
	p := NewPokedex()
	p.Add("pikachu", 112)

	name, ok := p.Get("pikachu")
	if !ok {
		t.Fatalf("expected pikachu to be found")
	}
	if name != "pikachu" {
		t.Fatalf("expected name 'pikachu', got %s", name)
	}
}

func TestGetMissing(t *testing.T) {
	p := NewPokedex()
	_, ok := p.Get("nonexistent")
	if ok {
		t.Fatalf("expected nonexistent pokemon to not be found")
	}
}
