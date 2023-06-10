package main

import "net/url"

type Page struct {
	Title       string    `yaml:"title"`
	Path        string    `yaml:"path"`
	Description string    `yaml:"description"`
	Sections    []Section `yaml:"sections"`
	NoIndex     bool      `yaml:"noIndex" json:"noIndex"`
}

// Assemble the complete URL for the page with given domain.
func (p *Page) URL(host string) string {
	url, err := url.Parse(host)

	if err != nil {
		panic(err)
	}

	url.Path = p.Path

	return url.String()
}
