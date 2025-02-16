package main

import (
	"fmt"
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
