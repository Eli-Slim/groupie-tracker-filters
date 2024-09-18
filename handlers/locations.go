package handlers

import (
	"encoding/json"
	"fmt"
	"groupietracker/services"
	"net/http"
)



func Location(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	locations, _ := services.GetLocationById(id)
	jsonData, _ := json.Marshal(locations)
	fmt.Fprint(w, string(jsonData))
}