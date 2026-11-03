package pokeapi

import (
	"encoding/json/v2"
)

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocations(pageURL *string) (LocationAreas, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// Check the cache
	if val, ok := c.cache.Get(url); ok {
		var locationAreas LocationAreas
		if err := json.Unmarshal(val, &locationAreas); err != nil {
			return LocationAreas{}, err
		}
		return locationAreas, nil

	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return LocationAreas{}, err
	}
	defer res.Body.Close()

	var locationAreas LocationAreas
	err = json.UnmarshalRead(res.Body, &locationAreas)
	if err != nil {
		return LocationAreas{}, err
	}

	locationVal, err := json.Marshal(locationAreas)
	if err != nil {
		return LocationAreas{}, err
	}
	c.cache.Add(url, locationVal)

	return locationAreas, nil
}
