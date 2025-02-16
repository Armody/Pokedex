package main

import (
	"errors"
	"fmt"
)

func exploreArea(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	locationName := args[0]
	locationResp, err := cfg.pokeapiClient.GetLocation(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationName)
	fmt.Println("Found Pokemons: ")

	for _, encounters := range locationResp.PokemonEncounters {
		fmt.Println(encounters.Pokemon.Name)
	}
	return nil
}
