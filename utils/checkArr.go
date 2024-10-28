package utils

import (
	"groupietracker/models"
)

func Contains(arr []models.Artist, artist models.Artist) bool {
	for _, a := range arr {
		if a.Id == artist.Id {
			return true
		}
	}
	return false
}
