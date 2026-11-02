package pokeapi

import (
	"encoding/json/v2"
	"net/http"
)

type LocatonAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocations(pageURL *string) (LocatonAreas, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	res, err := http.Get(url)
	if err != nil {
		return LocatonAreas{}, err
	}
	defer res.Body.Close()

	var locationArea LocatonAreas
	err = json.UnmarshalRead(res.Body, &locationArea)
	if err != nil {
		return LocatonAreas{}, err
	}
	return locationArea, nil
}
