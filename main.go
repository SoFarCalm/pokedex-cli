package main

import (
	"time"

	"github.com/SoFarCalm/pokedex-cli/internal/pokeapi"
	"github.com/SoFarCalm/pokedex-cli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokecache := pokecache.NewCache(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokeapiCache:  pokecache,
	}

	startRepl(cfg)
}
