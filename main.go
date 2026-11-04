package main

import (
	"time"

	"github.com/ppieniazek/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Minute)
	pokedex := make(map[string]pokeapi.Pokemon)
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeClient,
		pokedex:       pokedex,
	}
	startRepl(cfg)
}
