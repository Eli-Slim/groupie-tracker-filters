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

func GetLocations() ([]models.Locations, error) {
	var index models.Index
	locationsData, err := fetchLocations()
	if err != nil {
		return index.Locations, err
	}
	json.Unmarshal(locationsData, &index)
	return index.Locations, nil
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
