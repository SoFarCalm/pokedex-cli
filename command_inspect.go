package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]
	//keys := []string{"Name", "Height", "Weight", "Stats", "Types"}

	if pokemon, exists := cfg.caughtPokemon[pokemonName]; exists {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, v := range pokemon.Stats {
			fmt.Printf("- %s: %d\n", v.Stat.Name, v.BaseStat)
		}
		fmt.Println("Types:")
		for _, v := range pokemon.Types {
			fmt.Printf("- %s\n", v.Type.Name)
		}
		return nil
	}

	return errors.New("You have not caught that pokemon")
}
