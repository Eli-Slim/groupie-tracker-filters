package utils

import (
	"html/template"
)

func ParseTemplate(filename string) (*template.Template, error) {
	temp := template.New(filename)
	temp.Funcs(FuncMap)
	return temp.ParseFiles("./template/" + filename)
}
