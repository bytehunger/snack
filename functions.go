package main

import "html/template"

func safeHTML(s string) template.HTML {
	return template.HTML(s)
}
