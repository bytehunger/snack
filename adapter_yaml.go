package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type YAMLAdapter struct {
	Source string
}

// NewSite returns a new site created by a given JSON file.
func (a YAMLAdapter) NewSite() (site Site, err error) {
	siteFile, err := os.ReadFile(a.Source)

	if err != nil {
		return
	}

	err = yaml.Unmarshal(siteFile, &site)
	return
}
