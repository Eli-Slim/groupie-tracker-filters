package models

type Locations struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Date      string   `json:"date"`
}

type Index struct {
	Locations []Locations `json:"index"`
}
