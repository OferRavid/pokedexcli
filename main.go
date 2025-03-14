package main

import (
	"time"

	"github.com/OferRavid/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		caughtPokemon:      map[string]pokeapi.Pokemon{},
		caughtPokemonCount: map[string]int{},
		pokeapiClient:      pokeClient,
	}

	startRepl(cfg)
}
