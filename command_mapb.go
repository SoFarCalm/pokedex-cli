package main

import (
	"fmt"
)

func commandMapb(cfg *config) error {
	location, err := getPokeLocation(cfg.previous)
	if err != nil {
		return fmt.Errorf("failed to get location: %w", err)
	}

	//Update the config with the next and previous location URLs
	cfg.next = location.Next
	cfg.previous = location.Previous

	for _, results := range location.Results {
		fmt.Println(results.Name)
	}

	return nil
}
