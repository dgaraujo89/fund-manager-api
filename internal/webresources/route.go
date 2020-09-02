package webresources

import "net/http"

// Route path
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

// Routes registered routes
type Routes []Route
