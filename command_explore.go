package main

import "fmt"

func commandExplore(cfg *config, area string) error {

	fmt.Println("Exploring", area, "...")

	pokeResp, err := cfg.pokeapiClient.ListPokemon(area)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon: ")
	for _, item := range pokeResp.PokemonEncounters {
		fmt.Println(" - ", item.Pokemon.Name)
	}

	return nil
}
