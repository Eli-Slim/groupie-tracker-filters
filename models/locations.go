package models

type Locations struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Date      string   `json:"dates"`
}

type Index struct {
	Locations []Locations `json:"index"`
}