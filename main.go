package main

import (
	"time"

	"github.com/OferRavid/pokedexcli/internal/pokecache"
)

func main() {
	cache := pokecache.NewCache(5 * time.Second)
	startRepl(&config{}, &cache)
}
