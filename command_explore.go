package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) < 1 {
		return errors.New("not enough arguments")
	}
	locationName := args[0]

	locationResp, err := cfg.pokeapiClient.GetLocation(locationName)
	if err != nil {
		return err
	}

	if len(locationResp.PokemonEncounters) < 1 {
		return errors.New("Did not found any pokemons in the location")
	}

	fmt.Println("Exploring ", locationName, "...")
	fmt.Println("Found Pokemon:")
	for _, encounter := range locationResp.PokemonEncounters {
		fmt.Println(" - ", encounter.Pokemon.Name)
	}
	return nil
}
