package pokeapi

import (
	"encoding/json/v2"
	"fmt"
	"net/http"
)

func (c *Client) ExploreLocation(name string) (LocationArea, error) {
	url := baseURL + "/location-area/" + name

	if val, ok := c.cache.Get(url); ok {
		locationArea := LocationArea{}
		if err := json.Unmarshal(val, &locationArea); err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return LocationArea{}, fmt.Errorf("location area not found (status: %s)", res.Status)
	}

	var locationArea LocationArea
	err = json.UnmarshalRead(res.Body, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	locationVal, err := json.Marshal(locationArea)
	if err != nil {
		return LocationArea{}, err
	}
	c.cache.Add(url, locationVal)

	return locationArea, nil
}
