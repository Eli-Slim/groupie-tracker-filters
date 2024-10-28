package utils

import (
	"html/template"
	"strings"
)

func FixLocation(s string) string {
	s1 := strings.ReplaceAll(s, "-", ", ")
	s1 = strings.ReplaceAll(s1, "_", " ")
	return s1
}

var FuncMap = template.FuncMap{
	"FixLocation": FixLocation,
}
