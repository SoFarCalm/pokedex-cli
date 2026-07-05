package pokeapi

// This file defines a small HTTP client wrapper for the PokeAPI package.

// What it contains:

// A package named pokeapi, which suggests this code is meant to encapsulate API-related networking logic.
// A Client struct with one field:
// httpClient of type http.Client
// This stores the underlying Go standard library HTTP client used for making requests.
// A constructor function:
// NewClient(timeout time.Duration) Client
// This creates and returns a new Client instance.
// It initializes the embedded http.Client with a configurable timeout.
// Why it exists:

// It provides a reusable, typed wrapper around Go’s HTTP client.
// The timeout can be set when the client is created, which is useful for avoiding hanging requests when the API is slow or unavailable.
// In short, this file is a foundation for making PokeAPI requests in a controlled and reusable way.

import (
	"net/http"
	"time"

	"github.com/SoFarCalm/pokedex-cli/internal/pokecache"
)

// Client -
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// NewClient -
func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
