package main

// RenderData contains all necessary information for the
// generator in order to be able to render each section.
type RenderData struct {
	// Section is an embedded struct and contains the current section.
	Section

	// Site contains the whole site.
	Site *Site

	// Page contains the current page.
	Page *Page
}

// CurrentURL returns the current URL with path.
func (rd *RenderData) CurrentURL() string {
	return rd.Site.Host + rd.Page.Path
}

// CurrentPath returns the current path.
func (rd *RenderData) CurrentPath() string {
	return rd.Page.Path
}
