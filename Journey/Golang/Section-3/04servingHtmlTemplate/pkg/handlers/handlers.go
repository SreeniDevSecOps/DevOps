package handlers

import (
	"github.com/SreeniDevSecOps/DevOps/servingHTML/pkg/render"
	"net/http"
)

// HomePage is the About Page Handler:
func HomePage(w http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

// AboutPage is the About Page Handler:
func AboutPage(w http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
