package main

import (
	"html/template"
	"strings"
	"time"
)

func safeHTML(s string) template.HTML {
	return template.HTML(s)
}

func isActive(s1, s2 string) bool {
	s1 = strings.Trim(s1, "/")

	// special case: root link
	if s1 == "" {
		return s1 == s2
	}

	return strings.HasPrefix(s2, s1)
}

func isDate(w W3CDate) bool {
	t := time.Time(w)
	return !t.IsZero()
}

var funcMap = map[string]any{
	"safeHTML": safeHTML,
	"isActive": isActive,
	"isDate":   isDate,
}
