package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you didn't provide location name")
	}
	locationName := args[0]

	locationArea, err := cfg.pokeapiClient.ExploreLocation(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\nFound Pokemon:\n", locationName)
	for _, pokemonEncounter := range locationArea.PokemonEncounters {
		fmt.Printf(" - %v\n", pokemonEncounter.Pokemon.Name)
	}

	return nil
}
