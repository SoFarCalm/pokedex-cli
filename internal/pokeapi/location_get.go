package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// Explore Locations
func (c *Client) ExploreLocations(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	if cachedData, exists := c.cache.Get(url); exists {
		exploreResp := Location{}
		err := json.Unmarshal(cachedData, &exploreResp)
		if err != nil {
			return Location{}, err
		}

		return exploreResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	exploreResp := Location{}
	err = json.Unmarshal(data, &exploreResp)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, data)

	return exploreResp, nil
}
