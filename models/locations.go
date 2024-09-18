package models

type Locations struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Date      string   `json:"date"`
}
