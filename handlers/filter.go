package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"groupietracker/models"
	"groupietracker/services"
	"groupietracker/utils"
)

func FilterByCreationDate(creationDateMin, creationDateMax string, ids map[int]bool, w http.ResponseWriter, artists []models.Artist) {
	min, err := strconv.Atoi(creationDateMin)
	if err != nil || min < 1900 {
		utils.RenderError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	max, err := strconv.Atoi(creationDateMax)
	if err != nil || max < 1900 || min > max {
		utils.RenderError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	for _, artist := range artists {
		if artist.CreationDate >= min && artist.CreationDate <= max {
			ids[artist.Id] = true
		}
	}
}

func FilterByFirstAlbumDate(day, month, year string, ids map[int]bool, w http.ResponseWriter, artists []models.Artist) {
	Day, err := strconv.Atoi(day)
	if err != nil || Day < 1 || Day > 31 {
		utils.RenderError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	if Day >= 1 && Day <= 9 {
		day = "0" + day
	}
	Month, err := strconv.Atoi(month)
	if err != nil || Month < 1 || Month > 12 {
		utils.RenderError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	if Month >= 1 && Month <= 9 {
		month = "0" + month
	}
	Year, err := strconv.Atoi(year)
	if err != nil || Year < 1900 || Year > 2024 {
		utils.RenderError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	date := []string{day, month, year}
	firstAlbumDate := strings.Join(date, "-")
	for _, artist := range artists {
		if artist.FirstAlbum == firstAlbumDate {
			ids[artist.Id] = true
		}
	}
}

func FilterByMembers(members []string, ids map[int]bool, w http.ResponseWriter, artists []models.Artist) {
	membersNbr := []int{}
	for _, c := range members {
		member, err := strconv.Atoi(string(c))
		if err != nil || member < 1 {
			utils.RenderError(w, http.StatusBadRequest, "Bad Request")
			return
		}
		membersNbr = append(membersNbr, member)
	}
	for _, artist := range artists {
		for _, members := range membersNbr {
			if len(artist.Membres) == members {
				ids[artist.Id] = true
			}
		}
	}
}

func FilterByLocation(location string, ids map[int]bool, w http.ResponseWriter, locations []models.Locations, artists []models.Artist) {
	Id := []int{}
	for _, local := range locations {
		for _, l := range local.Locations {
			l1 := strings.ToLower(l)
			l2 := strings.ToLower(location)
			if strings.Contains(l1, l2) {
				Id = append(Id, local.Id)
			}
		}
	}
	for _, artist := range artists {
		for _, id := range Id {
			if artist.Id == id {
				ids[artist.Id] = true
			}
		}
	}

}

func getLocations(url string) ([]models.Locations, error) {
	var artist []models.Locations
	artistsData, err := utils.Fetch(url)
	if err != nil {
		return artist, err
	}
	json.Unmarshal(artistsData, &artist)
	return artist, nil
}

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

	erro := r.ParseForm()
	if erro != nil {
		utils.RenderError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	ids := make(map[int]bool)
	results := []models.Artist{}
	creationDateMin := r.PostFormValue("creationDateMin")
	creationDateMax := r.PostFormValue("creationDateMax")
	if len(creationDateMin) != 0 && len(creationDateMax) != 0 {
		FilterByCreationDate(creationDateMin, creationDateMax, ids, w, artists)
	}
	day := r.PostFormValue("day")
	month := r.PostFormValue("month")
	year := r.PostFormValue("year")
	if len(day) != 0 && len(month) != 0 && len(year) != 0 {
		FilterByFirstAlbumDate(day, month, year, ids, w, artists)
	}
	members := r.Form["members"]
	if len(members) != 0 {
		FilterByMembers(members, ids, w, artists)
	}
	location := r.PostFormValue("location")
	if len(location) != 0 {
		locations, err := getLocations("https://groupietrackers.herokuapp.com/api/locations")
		if err != nil {
			utils.RenderError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		fmt.Println(locations)
		FilterByLocation(location, ids, w, locations, artists)
	}
	for _, artist := range artists {
		if _, exist := ids[artist.Id]; exist {
			results = append(results, artist)
		}
	}
	temp, err := utils.ParseTemplate("filter-results.html")
	if err != nil {
		utils.RenderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	err = temp.Execute(w, results)
	if err != nil {
		utils.RenderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}
