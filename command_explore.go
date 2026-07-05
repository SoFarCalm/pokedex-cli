package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	//add logic to explore locations here
	if len(args) != 1 {
		return errors.New("You must provide a location name to explore. Usage: explore <location_name>")
	}

	locationName := args[0]
	exploreResp, err := cfg.pokeapiClient.ExploreLocations(locationName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", exploreResp.Name)
	fmt.Println("Found Pokemon: ")
	for _, encounter := range exploreResp.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
