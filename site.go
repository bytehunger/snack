package main

import "fmt"

type Site struct {
	Title         string   `json:"title"`
	Version       string   `json:"version"`
	Theme         string   `json:"theme"`
	Languages     []string `json:"languages"`
	Pages         []Page   `json:"pages"`
	GlobalWidgets []Widget `json:"globalWidgets"`
}

// Render renders the site with all its pages.
func (s *Site) Render() error {
	for _, page := range s.Pages {
		err := page.Render(s)

		if err != nil {
			fmt.Println(err)
		}
	}

	return nil
}
