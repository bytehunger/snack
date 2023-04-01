package main

import (
	"html/template"
	"strings"
)

func safeHTML(s string) template.HTML {
	return template.HTML(s)
}

func isActive(s1, s2 string) bool {
	return strings.HasPrefix(s2, strings.Trim(s1, "/"))
}

var funcMap = map[string]any{
	"safeHTML": safeHTML,
	"isActive": isActive,
}
