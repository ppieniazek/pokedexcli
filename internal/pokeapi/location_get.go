package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetLocation(name string) (RespLocation, error) {
	url := baseURL + "/location-area/" + name
	val, ok := c.cache.Get(url)
	if ok {
		locationResp := RespLocation{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return RespLocation{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocation{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return RespLocation{}, errors.New("Location not found")
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocation{}, err
	}

	locationResp := RespLocation{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return RespLocation{}, err
	}

	c.cache.Add(url, dat)
	return locationResp, nil
}
