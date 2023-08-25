package main

import (
	"net/url"
	"path/filepath"
	"strings"
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

// HasFileExtension checks if the URL path ends with an file extension
// like .html or .php.
func (p *Page) HasFileExtension() bool {
	return filepath.Ext(p.Path) != ""
}

// Pathname returns the right 'subdirectory' string.
func (p *Page) Pathname() string {
	if p.HasFileExtension() {
		return strings.TrimSuffix(p.Path, p.lastPathElement())
	}

	return p.Path
}

func (p *Page) Filename() string {
	if p.HasFileExtension() {
		return p.lastPathElement()
	}

	return "index.html"
}

func (p *Page) lastPathElement() string {
	return p.Path[strings.LastIndex(p.Path, "/")+1:]
}
