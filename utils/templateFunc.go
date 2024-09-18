package utils

import (
	"html/template"
	"strings"
)

func fixLocation(s string) string {
	s1 := strings.ReplaceAll(s, "-", ", ")
	s1 = strings.ReplaceAll(s1, "_", " ")
	return s1
}

var FuncMap = template.FuncMap{
	"FixLocation": fixLocation,
}
