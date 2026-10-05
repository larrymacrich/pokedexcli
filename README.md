# Pokedex CLI

A command-line Pokédex and interactive REPL built in Go. This tool interacts with the [PokéAPI](https://pokeapi.co/) to allow users to explore locations, encounter wild Pokémon, attempt to catch them, and inspect their stats and types.

To minimize external network calls and respect PokéAPI rate limits, the application includes a custom in-memory caching system with automatic stale-entry cleanup.

## Features

- **Interactive REPL**: Prompt-driven interface for navigating Pokémon locations and managing your Pokédex.
- **Location Exploration**: Paginated traversal through map locations and exploration of specific areas.
- **Catch & Inspect Mechanics**: Attempt to catch wild Pokémon with catch rates determined by base experience; inspect caught Pokémon details.
- **Custom Cache**: In-memory caching layer (`pokecache`) with configurable time-to-live (TTL) and background reap loops.

## Commands

- `help`: Displays a list of available commands and their descriptions.
- `map`: Displays the next 20 location areas.
- `mapb`: Displays the previous 20 location areas.
- `explore <area_name>`: Lists all Pokémon located in a given area.
- `catch <pokemon_name>`: Attempts to catch a Pokémon and add it to your Pokédex.
- `inspect <pokemon_name>`: Prints stats, types, height, and weight for a caught Pokémon.
- `pokedex`: Lists all the Pokémon you have successfully caught.
- `exit`: Exits the Pokédex CLI.

## Installation & Setup

### Prerequisites

- [Go](https://go.dev/dl/) (version 1.20 or newer recommended)
- Git

### Build and Run

1. Clone the repository:
   ```bash
   git clone https://github.com/<your-username>/<your-repo-name>.git
   cd <your-repo-name>

Run the application directly:

go run .

Alternatively, build and install the binary:

go build -o pokedexcli
./pokedexcli

Running Tests
To run the internal unit tests (including cache TTL validation):

go test ./...


Remember to replace `<your-username>` and `<your-repo-name>` with your actual GitHub details. If you implemented any additional extensions, you can list them under the **Features** section.


