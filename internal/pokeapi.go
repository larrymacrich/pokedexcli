package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// getLocationAreas returns 20 locations areas
func GetLocationAreas(url string) (*map[string]any, error) {
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
	var data map[string]any
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
