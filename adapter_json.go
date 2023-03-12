package main

import (
	"encoding/json"
	"os"
)

type JSONAdapter struct {
	Source string
}

// NewSite returns a new site created by a given JSON file.
func (a JSONAdapter) NewSite() (site Site, err error) {
	siteFile, err := os.ReadFile(a.Source)

	if err != nil {
		return
	}

	err = json.Unmarshal(siteFile, &site)
	return
}
