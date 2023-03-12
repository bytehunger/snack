package main

type Page struct {
	Title       string    `yaml:"title"`
	Path        string    `yaml:"path"`
	Description string    `yaml:"description"`
	Sections    []Section `yaml:"sections"`
}
