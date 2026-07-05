package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) < 1 {
		return errors.New("You have not caught any pokemon, try using the catching some")
	}

	fmt.Println("Your Pokedex:")
	for _, v := range cfg.caughtPokemon {
		fmt.Printf("- %s\n", v.Name)
	}
	fmt.Println()
	return nil
}
