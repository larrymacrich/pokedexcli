package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/larrymacrich/pokedexcli/internal/pokecache"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(cacheInterval time.Duration) *Client {
	return &Client{
		httpClient: http.Client{},
		cache:      pokecache.NewCache(cacheInterval),
	}
}

type LocationAreasResponse struct {
	Count    int                  `json:"count"`
	Next     *string              `json:"next"`
	Previous *string              `json:"previous"`
	Results  []LocationAreaResult `json:"results"`
}

type LocationAreaResult struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// getLocationAreas returns 20 locations areas
func (c *Client) GetLocationAreas(url string) (*LocationAreasResponse, error) {
	// check for cached requests
	if rawBytes, ok := c.cache.Get(url); ok {
		var data LocationAreasResponse
		if err := json.Unmarshal(rawBytes, &data); err != nil {
			errMsg := fmt.Errorf("unmarshaling cached json failed: %s", err)
			return nil, errMsg
		}
		return &data, nil
	}

	// make request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		errMsg := fmt.Errorf("request failed: %v", err)
		return nil, errMsg
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		errMsg := fmt.Errorf("network error: %v", err)
		return nil, errMsg
	}
	defer res.Body.Close()

	// check return codes
	if res.StatusCode != http.StatusOK {
		errMsg := fmt.Errorf("response failed with status code: %s", res.Status)
		return nil, errMsg
	}

	// decode and add data to cache
	var data LocationAreasResponse
	rawBytes, err := io.ReadAll(res.Body)
	if err != nil {
		errMsg := fmt.Errorf("reading response body failed: %s", err)
		return nil, errMsg
	}
	if err := json.Unmarshal(rawBytes, &data); err != nil {
		errMsg := fmt.Errorf("unmarshaling json failed: %s", err)
		return nil, errMsg
	}
	c.cache.Add(url, rawBytes)

	return &data, nil
}

type LocationAreaResponse struct {
	Id                int                `json:"id"`
	Name              string             `json:"name"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (c *Client) GetLocationArea(areaName string) (*LocationAreaResponse, error) {
	url := baseURL + "/location-area/" + areaName
	// check for cached requests
	if rawBytes, ok := c.cache.Get(url); ok {
		var data LocationAreaResponse
		if err := json.Unmarshal(rawBytes, &data); err != nil {
			errMsg := fmt.Errorf("unmarshaling cached json failed: %s", err)
			return nil, errMsg
		}
		return &data, nil
	}

	// make request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		errMsg := fmt.Errorf("request failed: %v", err)
		return nil, errMsg
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		errMsg := fmt.Errorf("network error: %v", err)
		return nil, errMsg
	}
	defer res.Body.Close()

	// check return codes
	if res.StatusCode != http.StatusOK {
		errMsg := fmt.Errorf("response failed with status code: %s", res.Status)
		return nil, errMsg
	}

	// decode and add data to cache
	var data LocationAreaResponse
	rawBytes, err := io.ReadAll(res.Body)
	if err != nil {
		errMsg := fmt.Errorf("reading response body failed: %s", err)
		return nil, errMsg
	}
	if err := json.Unmarshal(rawBytes, &data); err != nil {
		fmt.Println(url)
		errMsg := fmt.Errorf("unmarshaling json failed: %s", err)
		return nil, errMsg
	}
	c.cache.Add(url, rawBytes)

	return &data, nil
}
