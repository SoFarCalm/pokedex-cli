package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("%s has appeared!\n", pokemon.Name)
	time.Sleep(1 * time.Second)
	if pokemon.Name == "pikachu" {
		fmt.Println()
		fmt.Println("(\\o^ - ^o)")
		fmt.Println()
	}

	time.Sleep(1 * time.Second)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	time.Sleep(2 * time.Second)
	fmt.Printf("Pokeball moves back and forth...\n")
	time.Sleep(4 * time.Second)

	res := rand.Intn(pokemon.BaseExperience)
	if res > 40 {
		fmt.Printf("%s escaped pokeball and got away...darn\n", pokemon.Name)
		return nil
	}

	fmt.Printf("Yay! you caught a %s! Pokemon added to pokedex\n", pokemon.Name)
	cfg.caughtPokemon[pokemon.Name] = pokemon

	return nil
}
