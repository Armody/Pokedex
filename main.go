package main

import (
	"time"

	"github.com/Armody/Pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       make(map[string]pokeapi.Pokemon),
		experience:    1,
	}

	startRepl(cfg)
}
