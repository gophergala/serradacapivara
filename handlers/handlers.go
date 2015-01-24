package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gophergala/serradacapivara/db"
	"github.com/gorilla/schema"
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

// Admin Handlers

func NewSite(c web.C, w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/admin/new_site.go.html")

	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}

	t.Execute(w, nil)
}

func RegisterSite(c web.C, w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var site db.Site

	decoder := schema.NewDecoder()

	err = decoder.Decode(&site, r.PostForm)

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Println("%+v", site)

	// Validações

	if err := db.SaveSite(site); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
