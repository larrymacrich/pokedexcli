package cli

import "fmt"

func fetchAndDisplayLocations(cfg *config, url string) error {
	locationAreasResponse, err := cfg.pokeapiClient.GetLocationAreas(url)
	if err != nil {
		errMsg := fmt.Errorf("something went wrong: %s", err)
		return errMsg
	}

	cfg.next = locationAreasResponse.Next
	cfg.previous = locationAreasResponse.Previous

	for _, areaResult := range locationAreasResponse.Results {
		fmt.Println(areaResult.Name)
	}
	return nil
}
