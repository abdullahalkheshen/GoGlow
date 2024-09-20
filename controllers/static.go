package controllers

import (
	"net/http"

	"github.com/abdullahalkheshen/lenslocked/views"
)

// Static is a controller for serving static pages
type Static struct {
	// Template is a view used to render the HTML
	Template views.Template
}

// ServeHTTP implements the http.Handler interface
func (static Static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	static.Template.Execute(w, nil)
}

// StaticHandler returns a handler for serving static pages
func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
