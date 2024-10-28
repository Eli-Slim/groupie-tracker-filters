package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"groupietracker/models"
	"groupietracker/services"
	"groupietracker/utils"
)

func FilterByCreationDate(creationDateMin, creationDateMax string, artist models.Artist) (bool, error) {
	found := false
	min, err := strconv.Atoi(creationDateMin)
	if err != nil || min < 1900 {
		return found, err
	}
	max, err := strconv.Atoi(creationDateMax)
	if err != nil || max < 1900 || min > max {
		return found, err
	}
	if artist.CreationDate >= min && artist.CreationDate <= max {
		found = true
	}
	return found, nil
}

func FilterByFirstAlbumDate(FirstAlbumDateMin, FirstAlbumDateMax string, artist models.Artist) (bool, error) {
	found := false
	min, err := strconv.Atoi(FirstAlbumDateMin)
	if err != nil || min < 1900 {
		return found, err
	}
	max, err := strconv.Atoi(FirstAlbumDateMax)
	if err != nil || max < 1900 || min > max {
		return found, err
	}
	FirstAlbumDate := strings.Split(artist.FirstAlbum, "-")[2]
	FirstAlbumDateInt, err := strconv.Atoi(FirstAlbumDate)
	if err != nil {
		return found, err
	}
	if FirstAlbumDateInt >= min && FirstAlbumDateInt <= max {
		found = true
	}
	return found, nil
}

func FilterByMembers(members []string, artist models.Artist) (bool, error) {
	found := false
	membersNbr := []int{}
	for _, c := range members {
		member, err := strconv.Atoi(string(c))
		if err != nil || member < 1 {
			return found, err
		}
		membersNbr = append(membersNbr, member)
	}
	for _, members := range membersNbr {
		if len(artist.Members) == members {
			found = true
			break
		}
	}
	return found, nil
}

func FilterByLocation(location string, locations models.Locations) bool{
	found := false
	for _, local := range locations.Locations {
		l1 := strings.ToLower(local)
		l2 := strings.ToLower(location)
		l2 = strings.ReplaceAll(l2, ", ", "-")
		l2 = strings.ReplaceAll(l2, " ", "_")
		if strings.Contains(l1, l2) {
			found = true
			break
		}else if strings.Contains(l2, l1) {
			found = true
			break
		}
	}
	return found
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
	locations, err := services.GetLocations()
	if err != nil {
		utils.RenderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	erro := r.ParseForm()
	if erro != nil {
		utils.RenderError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	results := []models.Artist{}
	for _, artist := range artists {
		creationDateMin := r.PostFormValue("creationDateMin")
		creationDateMax := r.PostFormValue("creationDateMax")
		if len(creationDateMin) != 0 && len(creationDateMax) != 0 {
			flag, err := FilterByCreationDate(creationDateMin, creationDateMax, artist)
			if err != nil {
				utils.RenderError(w, http.StatusBadRequest, "Bad Request")
				return
			}
			if !flag {
				continue
			}
		}
		FirstAlbumDateMin := r.PostFormValue("FirstAlbumDateMin")
		FirstAlbumDateMax := r.PostFormValue("FirstAlbumDateMax")
		if len(FirstAlbumDateMin) != 0 && len(FirstAlbumDateMax) != 0 {
			flag, err := FilterByFirstAlbumDate(FirstAlbumDateMin, FirstAlbumDateMax, artist)
			if err != nil {
				utils.RenderError(w, http.StatusBadRequest, "Bad Request")
				return
			}
			if !flag {
				continue
			}
		}
		members := r.Form["members"]
		if len(members) != 0 {
			flag, err := FilterByMembers(members, artist)
			if err != nil {
				utils.RenderError(w, http.StatusBadRequest, "Bad Request")
				return
			}
			if !flag {
				continue
			}
		}
		location := r.PostFormValue("location")
		if len(location) != 0 {
			artistLocations := locations[artist.Id-1]
			flag := FilterByLocation(location, artistLocations)
			if !flag {
				continue
			}
		}
		results = append(results, artist)
	}
	if len(results) == 0 {
		utils.RenderError(w, http.StatusNotFound, "NO RESULTS FOUND")
			return
	}
	temp, err := utils.ParseTemplate("header.html", "filters.html", "results.html")
	if err != nil {
		utils.RenderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	err = temp.ExecuteTemplate(w, "results.html", results)
	if err != nil {
		utils.RenderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}
