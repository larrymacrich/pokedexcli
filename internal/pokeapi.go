package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
func GetLocationAreas(url string) (*LocationAreasResponse, error) {
	// request
	res, err := http.Get(url)
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

	// decode
	var data LocationAreasResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", data)

	return &data, nil
}
