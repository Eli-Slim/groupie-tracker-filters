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
	var locations []models.Locations
	artistsData, err := fetchLocations()
	if err != nil {
		return locations, err
	}
	json.Unmarshal(artistsData, &locations)
	return locations, nil
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
