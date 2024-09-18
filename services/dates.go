package services

import (
	"encoding/json"

	models "groupietracker/models"
	utils "groupietracker/utils"
)

func fetchDateById(id string) ([]byte, error) {
	return utils.FetchGroupieTracker("dates/" + id)
}

func fetchDates() ([]byte, error) {
	return utils.FetchGroupieTracker("dates")
}

func GetDates() ([]models.Dates, error) {
	var date []models.Dates
	artistsData, err := fetchDates()
	if err != nil {
		return date, err
	}
	json.Unmarshal(artistsData, &date)
	return date, nil
}

func GetDateById(id string) (models.Dates, error) {
	var date models.Dates
	dateData, err := fetchDateById(id)
	if err != nil {
		return date, err
	}
	json.Unmarshal(dateData, &date)
	return date, nil
}
