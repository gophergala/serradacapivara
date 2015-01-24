package main

import (
	"net/http"

	"github.com/gophergala/serradacapivara/handlers"
	"github.com/zenazn/goji"
)

func main() {
	// Website
	goji.Get("/", handlers.Index)
	goji.Get("/search", handlers.Search)
	goji.Get("/map", handlers.Map)
	goji.Get("/site/:id", handlers.Site)
	goji.Get("/about", handlers.About)

	// Admin
	goji.Get("/admin/newsite", handlers.NewSite)
	goji.Post("/admin/newsite", handlers.RegisterSite)

	// Static
	goji.Get("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	goji.Get("/layout/*", http.StripPrefix("/layout/", http.FileServer(http.Dir("layout"))))

	goji.Serve()
}
