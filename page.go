package main

import (
	"net/url"
)

type Page struct {
	Title       string    `yaml:"title" json:"title"`
	Path        string    `yaml:"path" json:"path"`
	Description string    `yaml:"description" json:"description"`
	Sections    []Section `yaml:"sections" json:"sections"`
	NoIndex     bool      `yaml:"noIndex" json:"noIndex"`
	PublishedAt W3CDate   `yaml:"publishedAt" json:"publishedAt`
	UpdatedAt   W3CDate   `yaml:"updatedAt" json:"updatedAt"`
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
