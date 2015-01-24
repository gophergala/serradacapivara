package main

import (
	"github.com/gophergala/serradacapivara/handlers"
	"github.com/zenazn/goji"
)

func main() {
	goji.Get("/", handlers.Index)
	goji.Get("/map", handlers.Map)
	goji.Get("/site/:name", handlers.Site)
	goji.Get("/about", handlers.About)

	goji.Get("/admin/newsite", handlers.NewSite)
	goji.Post("/admin/newsite", handlers.RegisterSite)

	goji.Serve()
}
