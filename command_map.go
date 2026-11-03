package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	locationAreas, err := cfg.pokeapiClient.GetLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationAreas.Next
	cfg.prevLocationsURL = locationAreas.Previous

	for _, location := range locationAreas.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return fmt.Errorf("you're on the first page")
	}
	locationArea, err := cfg.pokeapiClient.GetLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationArea.Next
	cfg.prevLocationsURL = locationArea.Previous

	for _, location := range locationArea.Results {
		fmt.Println(location.Name)
	}
	return nil
}
