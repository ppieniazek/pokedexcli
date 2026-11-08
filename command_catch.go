package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you have to provide pokemon name")
	}
	pokemonName := args[0]
	if pokemon, exists := cfg.pokedex[pokemonName]; exists {
		return fmt.Errorf("%v is already in your pokedex", pokemon.Name)
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)
	const threshold = 40
	if rand.IntN(pokemon.BaseExperience) < threshold {
		fmt.Printf("%v was caught!\n", pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command.")
		cfg.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%v escaped!\n", pokemon.Name)
	}
	return nil
}
