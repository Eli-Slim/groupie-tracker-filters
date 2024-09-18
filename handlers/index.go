package handlers

import (
	"net/http"

	"groupietracker/services"
	"groupietracker/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		utils.RenderError(w, http.StatusNotFound, "Page not found")
		return
	}
	if r.Method != http.MethodGet {
		utils.RenderError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	artists, err := services.GetArtists()
	if err != nil {
		utils.RenderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	
	tmp, err := utils.ParseTemplate("index.html")
	if err != nil {
		utils.RenderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	
	err = tmp.Execute(w, artists)
	if err != nil {
		utils.RenderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}
