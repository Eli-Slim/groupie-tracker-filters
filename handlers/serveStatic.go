package handlers

import (
	"net/http"
	"os"

	"groupietracker/utils"
)

func ServeStatic(w http.ResponseWriter, r *http.Request) {
	d, err := os.Stat("." + r.URL.Path)
	if d == nil || os.IsNotExist(err) || (d != nil && d.IsDir()) {
		utils.RenderError(w, http.StatusUnauthorized, "Page Unauthorized")
		return
	}
	fs := http.FileServer(http.Dir("./static"))
	http.StripPrefix("/static/", fs).ServeHTTP(w, r)
}
