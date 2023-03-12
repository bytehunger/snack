package main

import (
	"html/template"
	"os"
	"path/filepath"

	cp "github.com/otiai10/copy"
)

const outputDir = "build"

type Generator struct {
	Adapter Adapter
}

type PageData struct {
	*Site
	*Page
}

func (g *Generator) Generate() error {
	// Get the site from the adapter.
	site, err := g.Adapter.NewSite()
	if err != nil {
		return err
	}

	// Start the actual generating
	for _, page := range site.Pages {
		err := g.GeneratePage(
			&PageData{&site, &page},
		)

		if err != nil {
			return err
		}
	}

	assetsPath := filepath.Join("themes", site.Theme, "assets")

	if _, err = os.Stat(assetsPath); !os.IsNotExist(err) {
		err = cp.Copy(
			assetsPath,
			filepath.Join(outputDir, "assets"),
		)

		if err != nil {
			return err
		}
	}

	if _, err := os.Stat("pictures"); !os.IsNotExist(err) {
		err = cp.Copy("pictures", filepath.Join(outputDir, "pictures"))

		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Generator) GeneratePage(data *PageData) error {
	header, err := template.ParseFiles(
		filepath.Join("themes", data.Site.Theme, "header.html.tpl"),
	)

	if err != nil {
		return err
	}

	footer, err := template.ParseFiles(
		filepath.Join("themes", data.Site.Theme, "footer.html.tpl"),
	)

	if err != nil {
		return err
	}

	path := filepath.Join(outputDir, data.Page.Path)
	err = os.MkdirAll(path, 0755)

	if err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(path, "index.html"))

	if err != nil {
		return err
	}

	defer f.Close()

	err = header.Execute(f, data)

	if err != nil {
		return err
	}

	// Add global sections to page header...
	data.Page.Sections = append(
		data.Site.GlobalSections.Header,
		data.Page.Sections...,
	)

	// and footer.
	data.Page.Sections = append(
		data.Page.Sections,
		data.Site.GlobalSections.Footer...,
	)

	for _, section := range data.Page.Sections {
		basename := section.Type + ".html.tpl"

		tpl, err := template.New(basename).
			Funcs(template.FuncMap{"safeHTML": safeHTML}).
			ParseFiles(
				filepath.Join(
					"themes",
					data.Site.Theme,
					"sections", basename,
				),
			)

		if err != nil {
			return err
		}

		err = tpl.Execute(f, section)

		if err != nil {
			return err
		}
	}

	err = footer.Execute(f, data)

	if err != nil {
		return err
	}

	return nil
}
