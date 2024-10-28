package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"groupietracker/models"
	"groupietracker/services"
	"groupietracker/utils"
)

func Search(w http.ResponseWriter, r *http.Request) {
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
	foundArtists := []models.Artist{}
	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
			if !utils.Contains(foundArtists, artist) {
				foundArtists = append(foundArtists, artist)
			}
		} else if strings.Contains(strings.ToLower(strconv.Itoa(artist.CreationDate)), strings.ToLower(query)) {
			if !utils.Contains(foundArtists, artist) {
				foundArtists = append(foundArtists, artist)
			}
		} else if strings.Contains(strings.ToLower(artist.FirstAlbum), strings.ToLower(query)) {
			if !utils.Contains(foundArtists, artist) {
				foundArtists = append(foundArtists, artist)
			}
		}
	}
	for _, artist := range artists {
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), strings.ToLower(query)) {
				if !utils.Contains(foundArtists, artist) {
					foundArtists = append(foundArtists, artist)
				}
			}
		}
	}
	for _, location := range locations {
		for _, loc := range location.Locations {
			if strings.Contains(strings.ToLower(loc), strings.ToLower(query)) {
				if !utils.Contains(foundArtists, artists[location.Id-1]) {
					foundArtists = append(foundArtists, artists[location.Id-1])
				}
			}
		}
	}
	if len(foundArtists) == 0 {
		utils.RenderError(w, http.StatusNotFound, "NO RESULTS FOUND")
		return
	}
	tmp, err := utils.ParseTemplate("header.html", "filters.html", "results.html")
	if err != nil {
		utils.RenderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	err = tmp.ExecuteTemplate(w, "results.html", foundArtists)
	if err != nil {
		utils.RenderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}
