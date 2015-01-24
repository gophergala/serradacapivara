package handlers

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"
)

func Index(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Serra da Capivara!")
}
