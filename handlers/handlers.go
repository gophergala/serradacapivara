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

// Index is the website homepage
func Index(c web.C, w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.go.html")

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}

// Search is the page where the results of search are shown
func Search(c web.C, w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/search.go.html")

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}

// Map is the page where all sites are shown at once
func Map(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Map")
}

// Site is the page that describe the given site
func Site(c web.C, w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/site.go.html")

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	site, err := db.FindByID(c.URLParams["id"])

	if err != nil {
		log.Println(err)
		http.Error(w, "Not Found", http.StatusNotFound)
		return

	}

	t.Execute(w, site)
}

// About the project
func About(c web.C, w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/about.go.html")

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}

// Admin Handlers

func NewSite(c web.C, w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/admin/new_site.go.html")

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
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
