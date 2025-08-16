package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]

	location, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location.Name)

	fmt.Println("Found Pokemon: ")
	for _, item := range location.PokemonEncounters {
		fmt.Println(" - ", item.Pokemon.Name)
	}

	return nil
}
