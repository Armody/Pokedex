package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Printf(`Welcome to the Pokedex!
Usage:

`)
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func catch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]

	if _, exists := cfg.pokedex[pokemonName]; exists {
		fmt.Printf("%s already in pokedex\n", pokemonName)
		return nil
	}

	pokemonResp, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	if rand.Intn(pokemonResp.BaseExperience) > 20 {
		fmt.Printf("%s escaped...\n", pokemonName)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonName)
	cfg.pokedex[pokemonName] = pokemonResp

	return nil
}

func commandMapf(cfg *config, args ...string) error {

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

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

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
