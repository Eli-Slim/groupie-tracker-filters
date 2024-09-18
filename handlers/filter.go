package handlers

import (
	"net/http"

	"groupietracker/services"
	"groupietracker/utils"
)

func Filter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.RenderError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	
	artists, err := services.GetArtists()
	if err != nil {
		utils.RenderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}
