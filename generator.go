package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"

	cp "github.com/otiai10/copy"

	"github.com/bytehunger/snack/sitemap"
)

const outputDir = "build"

type Generator struct {
	Adapter Adapter
}

func (g *Generator) Generate() error {
	// Get the site from the adapter.
	site, err := g.Adapter.NewSite()
	if err != nil {
		return err
	}

	// Start the actual generating
	for _, page := range site.Pages {
		renderData := RenderData{
			Site: &site,
			Page: &page,
		}

		err := g.GeneratePage(&renderData)

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

	err = g.GenerateSitemap(site)

	if err != nil {
		return err
	}

	return g.GenerateContentJSON(site)
}

func (g *Generator) GeneratePage(data *RenderData) error {
	header, err := template.New("header.html.tpl").Funcs(funcMap).ParseFiles(
		filepath.Join("themes", data.Site.Theme, "header.html.tpl"),
	)

	if err != nil {
		return err
	}

	footer, err := template.New("footer.html.tpl").Funcs(funcMap).ParseFiles(
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

	// Concatenate all global and local sections.
	sections := append(
		append(
			data.Site.GlobalSections.Header,
			data.Page.Sections...,
		),
		data.Site.GlobalSections.Footer...,
	)

	for _, section := range sections {

		// Don't render inactive sections.
		if section.Inactive {
			continue
		}

		basename := section.Type + ".html.tpl"

		data.Section = section

		tpl, err := template.New(basename).
			Funcs(funcMap).
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

		err = tpl.Execute(f, data)

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

func (g *Generator) GenerateContentJSON(site Site) error {
	type ContentPage struct {
		Title       string    `json:"title"`
		Path        string    `json:"path"`
		Description string    `json:"description"`
		Sections    []Section `json:"-"`
		NoIndex     bool      `json:"-"`
		PublishedAt W3CDate   `json:"publishedAt`
		UpdatedAt   W3CDate   `json:"updatedAt"`
	}

	pages := []ContentPage{}

	for _, page := range site.Pages {

		if page.NoIndex {
			continue
		}

		pages = append(pages, ContentPage(page))
	}

	content, err := json.Marshal(pages)

	if err != nil {
		return err
	}

	path := filepath.Join(outputDir, "content.json")
	return ioutil.WriteFile(path, content, 0644)
}

func (g *Generator) GenerateSitemap(site Site) error {
	urls := []sitemap.URL{}

	for _, page := range site.Pages {

		// Do not include noindex pages.
		if page.NoIndex {
			continue
		}

		url := sitemap.URL{
			Loc: page.URL(site.Host),
		}

		if page.PublishedAt.IsDate() {
			url.LastMod = page.PublishedAt.String()
		}

		if page.UpdatedAt.IsDate() {
			url.LastMod = page.UpdatedAt.String()
		}

		// Append all pages to the URLSet.
		urls = append(urls, url)
	}

	urlset := sitemap.URLSet{
		URLs:  urls,
		XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9",
	}

	path := filepath.Join(outputDir, "sitemap.xml")
	return urlset.Write(path)
}
