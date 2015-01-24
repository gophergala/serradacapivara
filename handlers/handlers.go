package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"strconv"

	"github.com/gophergala/serradacapivara/db"
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
	name := r.FormValue("name")
	hasPicture := r.FormValue("has_picture")
	hasEngraving := r.FormValue("has_engraving")
	other := r.FormValue("other")
	isHistoric := r.FormValue("is_historic")
	cityId, err := strconv.Atoi(r.FormValue("city_id"))
	circuitId, err := strconv.Atoi(r.FormValue("circuit_id"))
	nationalParkId, err := strconv.Atoi(r.FormValue("national_park_id"))
	locationId, err := strconv.Atoi(r.FormValue("location_id"))

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	site := db.Site{
		Name:           name,
		HasPicture:     hasPicture == "on",
		HasEngraving:   hasEngraving == "on",
		Other:          other == "on",
		IsHistoric:     isHistoric == "on",
		CityId:         cityId,
		CircuitId:      circuitId,
		NationalParkId: nationalParkId,
		LocationId:     locationId,
	}

	log.Println("%+v", site)

}
