package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func catch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]
	pokemonResp, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	if rand.Intn(pokemonResp.BaseExperience) > cfg.experience {
		fmt.Printf("%s escaped...\n", pokemonName)

		cfg.experience++
		fmt.Println("You gain 1 xp")
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonName)
	cfg.pokedex[pokemonName] = pokemonResp

	expGain := int(float32(pokemonResp.BaseExperience) / 100.0 * 5)
	cfg.experience += expGain
	fmt.Printf("You gain %v xp\n", expGain)
	return nil
}
