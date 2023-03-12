package main

// Site is the top-entity and represents the website.
type Site struct {
	Title          string         `yaml:"title" json:"title"`
	Description    string         `yaml:"description" json:"description"`
	Host           string         `yaml:"host" json:"host"`
	Theme          string         `yaml:"theme" json:"theme"`
	Favicon        string         `yaml:"favicon" json:"favicon"`
	Settings       Settings       `yaml:"settings" json:"settings"`
	Pages          []Page         `yaml:"pages" json:"pages"`
	GlobalSections GlobalSections `yaml:"globalSections" json:"globalSections"`
}

// GlobalSections is just a helper struct to include global sections
// into a page.
type GlobalSections struct {
	Header []Section
	Footer []Section
}
