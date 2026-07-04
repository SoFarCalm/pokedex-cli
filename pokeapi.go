package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PokeLocation struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getPokeLocation(url string) (PokeLocation, error) {
	// Create GET request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeLocation{}, fmt.Errorf("failed to create request: %w", err)
	}

	//Create HTTP client and perform the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return PokeLocation{}, fmt.Errorf("failed to perform request: %w", err)
	}
	defer response.Body.Close()

	// Decode the JSON response
	var location PokeLocation
	if err := json.NewDecoder(response.Body).Decode(&location); err != nil {
		return PokeLocation{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return location, nil
}
