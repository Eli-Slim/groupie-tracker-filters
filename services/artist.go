package services

import (
	"encoding/json"

	models "groupietracker/models"
	utils "groupietracker/utils"
)

func fetchArtistById(id string) ([]byte, error) {
	return utils.FetchGroupieTracker("artists/" + id)
}

func fetchArtists() ([]byte, error) {
	return utils.FetchGroupieTracker("artists")
}

func GetArtists() ([]models.Artist, error) {
	var artist []models.Artist
	artistsData, err := fetchArtists()
	if err != nil {
		return artist, err
	}
	json.Unmarshal(artistsData, &artist)
	return artist, nil
}

func GetArtistById(id string) (models.Artist, error) {
	var artist models.Artist
	artistData, err := fetchArtistById(id)
	if err != nil {
		return artist, err
	}
	json.Unmarshal(artistData, &artist)
	return artist, nil
}
