package utils

import (
	"html/template"
)

func ParseTemplate(filename ...string) (*template.Template, error) {
	temp := template.New("")
	temp.Funcs(FuncMap)
	for i, file := range filename {
		filename[i] = "./template/" + file
	}
	return temp.ParseFiles(filename...)
}
