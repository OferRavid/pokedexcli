package main

import (
	"time"

	"github.com/OferRavid/pokedexcli/internal/pokeapi"
)

/*
main initializes the application by creating a PokeAPI client with specific timeouts,
sets up the configuration struct, and starts the REPL (Read-Eval-Print) Loop.
*/
func main() {
	// Create a new PokeAPI client with a 5-second HTTP timeout and a 5-minute cache duration.
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)

	// Initialize the application configuration, including caches for caught Pokemon
	// and a reference to the PokeAPI client.
	cfg := &config{
		caughtPokemon:      map[string]pokeapi.Pokemon{},
		caughtPokemonCount: map[string]int{},
		pokeapiClient:      pokeClient,
	}

	// Start the REPL to process user commands.
	startRepl(cfg)
}
