package pokeapi

import (
	"encoding/json/v2"
	"fmt"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	if val, ok := c.cache.Get(url); ok {
		var pokemon Pokemon
		if err := json.Unmarshal(val, &pokemon); err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return Pokemon{}, fmt.Errorf("pokemon not found (status: %s)", res.Status)
	}

	var pokemon Pokemon
	if err := json.UnmarshalRead(res.Body, &pokemon); err != nil {
		return Pokemon{}, err
	}

	pokemonVal, err := json.Marshal(pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, pokemonVal)

	return pokemon, nil
}
