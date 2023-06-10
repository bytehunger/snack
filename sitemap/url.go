package sitemap

import "encoding/xml"

type URL struct {
	XMLName    xml.Name `xml:"url"`
	Loc        string   `xml:"loc"`
	ChangeFreq string   `xml:"changefreq,omitempty"`
	LastMod    string   `xml:"lastmod,omitempty"`
	Priority   string   `xml:"priority,omitempty"`
}
