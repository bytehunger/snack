package main

type Adapter interface {
	// NewSite returns a site
	NewSite() (Site, error)
}
