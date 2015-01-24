package main

import (
	"github.com/gophergala/serradacapivara/handlers"
	"github.com/zenazn/goji"
)

func main() {
	goji.Get("/", handlers.Index)
	goji.Serve()
}
