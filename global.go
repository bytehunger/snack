package main

// Global is a wrapper for global information given to
// the section while rendering.
type Global struct {
	// Host is the global host url.
	Host string

	// CurrentURL returns the current URL.
	CurrentURL string

	// CurrentPath returns the current path.
	CurrentPath string
}
