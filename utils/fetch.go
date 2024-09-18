package utils

import (
	"io"
	"net/http"

	config "groupietracker/config"
)

func Fetch(url string) ([]byte, error) {
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	return io.ReadAll(response.Body)
}

func FetchGroupieTracker(endpoint string) ([]byte, error) {
	return Fetch(config.API_URL + endpoint)
}
