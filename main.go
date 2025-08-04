package main

import (
	"pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	pokeClient, err := pokeapi.NewClient(5 * time.Second)
	if err != nil {
		return
	}

	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
