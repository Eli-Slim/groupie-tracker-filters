package main

import (
	"fmt"
	"net/http"

	"groupietracker/config"
	"groupietracker/handlers"
)

func main() {
	http.HandleFunc("/static/", handlers.ServeStatic)
	http.HandleFunc("/band/{id}", handlers.Band)
	http.HandleFunc("/locations/{id}", handlers.Location)
	http.HandleFunc("/filter-results", handlers.Filter)
	http.HandleFunc("/", handlers.Index)

	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(config.PORT, nil)
	if err != nil {
		fmt.Println(err)
	}
}
