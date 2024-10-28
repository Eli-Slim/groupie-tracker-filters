package handlers

import (
	"net/http"

	"groupietracker/models"
	"groupietracker/services"
	"groupietracker/utils"
)

type pageData struct {
	Artist   models.Artist
	Relation models.Relation
}

func Band(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RenderError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	id := r.PathValue("id")
	pageData := pageData{}

	artist, err := services.GetArtistById(id)
	if err != nil {
		utils.RenderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if artist.Id == 0 {
		utils.RenderError(w, http.StatusNotFound, "BAND NOT FOUND")
		return
	}

	relation, err := services.GetRelationById(id)
	if err != nil {
		utils.RenderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	pageData.Artist = artist
	pageData.Relation = relation

	temp, err := utils.ParseTemplate("header.html", "filters.html", "band.html")
	if err != nil {
		utils.RenderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	err = temp.ExecuteTemplate(w, "band.html", pageData)
	if err != nil {
		utils.RenderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}
