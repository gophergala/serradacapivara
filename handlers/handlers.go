package handlers

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"
)

func Index(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Serra da Capivara!")
}

func Map(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Map")
}

func Site(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Site, %s", c.URLParams["name"])
}

func About(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About")
}
