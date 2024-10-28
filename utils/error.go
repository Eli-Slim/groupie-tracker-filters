package utils

import (
	"net/http"

	"groupietracker/config"
)

func NewError(errorCode int, errorMessage string) config.ErrorData {
	err := config.ErrorData{
		ErrorCode: errorCode,
		Message:   errorMessage,
	}
	return err
}



func RenderError(w http.ResponseWriter, errCode int, message string) {
	w.WriteHeader(errCode)
	if(message == "NO RESULTS FOUND" || message == "BAND NOT FOUND") {
		temp, _ := ParseTemplate("header.html", "filters.html", "no-results.html")
		err := temp.ExecuteTemplate(w, "no-results.html", NewError(errCode, message))
		if err != nil {
			RenderError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
	}else{
		temp, _ := ParseTemplate("header.html", "filters.html", "error.html")
		err := temp.ExecuteTemplate(w, "error.html", NewError(errCode, message))
		if err != nil {
			RenderError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
	}
}
