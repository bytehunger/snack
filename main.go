package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"

	cp "github.com/otiai10/copy"
)

type Site struct {
	Title         string   `json:"title"`
	Version       string   `json:"version"`
	Theme         string   `json:"theme"`
	Languages     []string `json:"languages"`
	Pages         []Page   `json:"pages"`
	GlobalWidgets []Widget `json:"globalWidgets"`
}

type Page struct {
	Title       string   `json:"title"`
	Path        string   `json:"path"`
	Description string   `json:"description"`
	Widgets     []Widget `json:"widgets"`
}

type Widget struct {
	Name     string                 `json:"name"`
	Position string                 `json:"position"`
	Content  map[string]interface{} `json:"content"`
}

func main() {
	file, err := ioutil.ReadFile("site.json")

	if err != nil {
		panic("Could not read site")
	}

	var site Site
	err = json.Unmarshal([]byte(file), &site)

	if err != nil {
		panic("Could not parse site")
	}

	err = site.Render()

	if err != nil {
		panic("Could not render site")
	}

	err = cp.Copy("themes/"+site.Theme+"/assets", "public/assets")

	if err != nil {
		panic("Could not copy assets")
	}
}

// Render renders the site with all its pages.
func (s *Site) Render() error {
	for _, page := range s.Pages {
		err := page.Render(s)

		if err != nil {
			fmt.Println(err)
		}
	}

	return nil
}

// Render renders a single page with all its widgets.
func (p *Page) Render(site *Site) error {
	header, err := template.ParseFiles("themes/" + site.Theme + "/header.html.tpl")

	if err != nil {
		return err
	}

	footer, err := template.ParseFiles("themes/" + site.Theme + "/footer.html.tpl")

	if err != nil {
		return err
	}

	path := "public" + p.Path
	err = os.MkdirAll(path, 0755)

	if err != nil {
		return err
	}

	f, err := os.Create(path + "/index.html")

	if err != nil {
		return err
	}

	defer f.Close()

	err = header.Execute(f, p)

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

	err = footer.Execute(f, p)

	return err
}

func (w *Widget) Render(f io.Writer, s *Site) error {
	basename := w.Name + ".html.tpl"
	path := "themes/" + s.Theme + "/widgets/"

	t, err := template.New(basename).
		Funcs(template.FuncMap{"safeHTML": safeHTML}).
		ParseFiles(path + basename)

	if err != nil {
		return err
	}

	return t.Execute(f, w)
}

func safeHTML(s string) template.HTML {
	return template.HTML(s)
}
