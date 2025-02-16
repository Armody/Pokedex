package main

import (
	"errors"
	"fmt"
)

func inspectPokemon(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must input a pokemon name")
	}

	pokemonName := args[0]
	pokemon, exists := cfg.pokedex[pokemonName]
	if !exists {
		fmt.Printf("%s isn't in your pokedex", pokemonName)
		return nil
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, types := range pokemon.Types {
		fmt.Printf("  -%s\n", types.Type.Name)
	}

	return nil
}
