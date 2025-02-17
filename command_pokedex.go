package main

import (
	"errors"
	"fmt"
)

func showPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex) == 0 {
		return errors.New("didn't catch any pokemons yet")
	}

	for _, pokemon := range cfg.pokedex {
		fmt.Println("  -", pokemon.Name)
	}
	return nil
}
