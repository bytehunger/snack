package main

import (
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type FileAdapter struct {
	Source string
}

// Return a new site which was read from a file structure.
func (a *FileAdapter) NewSite() (site Site, err error) {
	siteFile, err := os.ReadFile(filepath.Join(a.Source, "site.yml"))

	if err != nil {
		return
	}

	err = yaml.Unmarshal(siteFile, &site)

	if err != nil {
		return
	}

	site.Pages, err = a.LoadPages()

	if err != nil {
		return
	}

	return
}

// LoadPages load all pages from the directories and assigns them to the site.
func (a *FileAdapter) LoadPages() (pages []Page, err error) {
	err = filepath.Walk(filepath.Join(a.Source, "pages"), func(path string, info os.FileInfo, err error) error {
		// Start the iteration.
		if err != nil {
			return err
		}

		// Ignore directories
		if info.IsDir() {
			return nil
		}

		var page Page
		pageFile, err := os.ReadFile(path)

		if err != nil {
			return err
		}

		err = yaml.Unmarshal(pageFile, &page)

		if err != nil {
			return err
		}

		page.Path = normalizePagePath(path)

		pages = append(pages, page)

		return nil
	})

	return
}

// Takes a file path and returns the corresponding URL path.
func normalizePagePath(path string) string {
	// Remove the pages directory
	np := strings.TrimPrefix(path, "pages")

	// Remove the "html" file ending
	np = strings.TrimSuffix(np, filepath.Ext(path))

	// Remove the optional 'index' from the end
	np = strings.TrimSuffix(np, "index")

	// Add a trailing slash for SEO optimization.
	if !strings.HasSuffix(np, "/") {
		np = np + "/"
	}

	return np
}
