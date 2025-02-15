package main

import (
	"time"

	"github.com/Armody/Pokedex/internal/pokeapi"
	"github.com/Armody/Pokedex/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokecache.NewCache(10 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokeCache:     pokeCache,
	}

	startRepl(cfg)
}
