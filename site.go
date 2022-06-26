package main

import (
	"fmt"

	cp "github.com/otiai10/copy"
)

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

// Copy assets from the theme.
func (s *Site) CopyAssets(outputDir string) error {
	err := cp.Copy("themes/"+s.Theme+"/assets", outputDir+"/assets")

	if err != nil {
		return err
	}

	return cp.Copy("pictures/", outputDir+"/pictures")
}
