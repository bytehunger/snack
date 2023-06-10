package sitemap

import (
	"encoding/xml"
	"os"
)

type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNS   string   `xml:"xmlns,attr"`
	URLs    []URL    `xml:"url"`
}

func (u *URLSet) Write(path string) error {
	sitemap, err := xml.MarshalIndent(u, "", " ")

	if err != nil {
		return err
	}

	file, err := os.Create(path)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(sitemap)

	return err
}
