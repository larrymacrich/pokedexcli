package main

func main() {
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	cfg := config{
		commands: getCliCommands(),
		next:     &baseURL,
	}
	startRepl(&cfg)
}
