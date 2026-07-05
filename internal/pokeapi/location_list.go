package pokeapi

// This file defines a method named ListLocations that lets the client fetch a paginated list of location areas from the PokeAPI.

// What it does:

// It builds the request URL using the base API address plus the /location-area endpoint.
// If a page URL is provided, it uses that instead of the default endpoint, which supports pagination.
// It creates an HTTP GET request.
// It sends the request through the client’s configured HTTP client.
// It reads the response body and decodes the JSON into a RespShallowLocations structure.
// Why it is useful:

// It abstracts the API call into a simple method.
// It allows the application to retrieve location data without manually handling HTTP requests or JSON decoding.
// The pagination support makes it suitable for walking through large result sets.
// In short, this file is responsible for fetching and parsing location-area data from the PokeAPI.

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {

	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if cachedData, exists := c.cache.Get(url); exists {
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(cachedData, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, data)

	return locationsResp, nil
}
