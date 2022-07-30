package main

import (
	"html/template"
	"os"
)

type Page struct {
	Title       string   `json:"title"`
	Path        string   `json:"path"`
	Description string   `json:"description"`
	Widgets     []Widget `json:"widgets"`
}

// Render renders a single page with all its widgets.
func (p *Page) Render(site *Site) error {
	data := struct {
		*Page
		Site *Site
	}{p, site}

	header, err := template.ParseFiles("themes/" + site.Theme + "/header.html.tpl")

	if err != nil {
		return err
	}

	footer, err := template.ParseFiles("themes/" + site.Theme + "/footer.html.tpl")

	if err != nil {
		return err
	}

	path := outputDir + p.Path
	err = os.MkdirAll(path, 0755)

	if err != nil {
		return err
	}

	f, err := os.Create(path + "/index.html")

	if err != nil {
		return err
	}

	defer f.Close()

	err = header.Execute(f, data)

	if err != nil {
		return err
	}

	// Render global header widgets
	for _, widget := range site.GlobalWidgets {
		if widget.Position != "header" {
			continue
		}

		err := widget.Render(f, site)

		if err != nil {
			return err
		}
	}

	for _, widget := range p.Widgets {
		err := widget.Render(f, site)

		if err != nil {
			return err
		}
	}

	// Render global footer widgets
	for _, widget := range site.GlobalWidgets {
		if widget.Position != "footer" {
			continue
		}

		err := widget.Render(f, site)

		if err != nil {
			return err
		}
	}

	err = footer.Execute(f, data)

	return err
}
