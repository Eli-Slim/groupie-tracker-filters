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
	temp, _ := ParseTemplate("Error.html")
	err := temp.Execute(w, NewError(errCode, message))
	if err != nil {
		RenderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}
