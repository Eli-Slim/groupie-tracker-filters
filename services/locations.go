package services

import (
	"encoding/json"

	models "groupietracker/models"
	utils "groupietracker/utils"
)

func fetchLocationById(id string) ([]byte, error) {
	return utils.FetchGroupieTracker("locations/" + id)
}

func fetchLocations() ([]byte, error) {
	return utils.FetchGroupieTracker("locations")
}

func GetLocations() (models.Index, error) {
	var location models.Index
	artistsData, err := fetchLocations()
	if err != nil {
		return location, err
	}
	json.Unmarshal(artistsData, &location)
	return location, nil
}

func GetLocationById(id string) (models.Locations, error) {
	var location models.Locations
	locationData, err := fetchLocationById(id)
	if err != nil {
		return location, err
	}
	json.Unmarshal(locationData, &location)
	return location, nil
}
