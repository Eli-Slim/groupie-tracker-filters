package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"groupietracker/models"
	"groupietracker/services"
	"groupietracker/utils"
)

const API_URL = "http://localhost:8080/"

func Suggestions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RenderError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	query := r.URL.Query().Get("search")
	if query == "" {
		utils.RenderError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	artistCh := make(chan []models.Artist)
	locationCh := make(chan []models.Locations)
	errCh := make(chan error, 2)
	go func() {
		artists, err := services.GetArtists()
		if err != nil {
			errCh <- err
			return
		}
		artistCh <- artists
	}()

	go func() {
		locations, err := services.GetLocations()
		if err != nil {
			errCh <- err
			return
		}
		locationCh <- locations
	}()
	var artists []models.Artist
	var locations []models.Locations
	for i := 0; i < 2; i++ {
		select {
		case artist := <-artistCh:
			artists = artist
		case location := <-locationCh:
			locations = location
		case err := <-errCh:
			utils.RenderError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	suggestions := []map[string]string{}
	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
			suggestions = append(suggestions, map[string]string{
				artist.Name + " - artist/band": API_URL + "/band/" + strconv.Itoa(artist.Id),
			})
		}
		if strings.Contains(strings.ToLower(strconv.Itoa(artist.CreationDate)), strings.ToLower(query)) {
			suggestions = append(suggestions, map[string]string{
				strconv.Itoa(artist.CreationDate) + " - creation date": API_URL + "/band/" + strconv.Itoa(artist.Id),
			})
		}
		if strings.Contains(strings.ToLower(artist.FirstAlbum), strings.ToLower(query)) {
			suggestions = append(suggestions, map[string]string{
				artist.FirstAlbum + " - first album": API_URL + "/band/" + strconv.Itoa(artist.Id),
			})
		}
	}
	for _, artist := range artists {
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), strings.ToLower(query)) {
				suggestions = append(suggestions, map[string]string{
					member + " - member": API_URL + "/band/" + strconv.Itoa(artist.Id),
				})
			}
		}
	}
	for _, location := range locations {
		for _, loc := range location.Locations {
			if strings.Contains(strings.ToLower(utils.FixLocation(loc)), strings.ToLower(query)) {
				suggestions = append(suggestions, map[string]string{
					utils.FixLocation(loc) + " - location": API_URL + "/band/" + strconv.Itoa(location.Id),
				})
			}
		}
	}
	suggestions = utils.RemoveDuplicates(suggestions)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(suggestions)
}
