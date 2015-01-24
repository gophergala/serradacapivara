package db

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path"
	"runtime"
)

// Site is an archeological site
type Site struct {
	Id              string
	Name            string
	HasPicture      bool
	HasEngraving    bool
	Other           bool
	IsHistoric      bool
	YearOfDiscovery int
	Latitude        string
	Longitude       string
	City            string
	Circuit         string
	NationalPark    string
	Location        string
	Pictures        []string
}

type Database []Site

var DB Database

// Find the Site by its identifier
func FindByID(id string) (Site, error) {
	for _, s := range DB {
		if s.Id == id {
			return s, nil
		}
	}

	return Site{}, errors.New("Site not found")
}

// Initialize the database
func init() {
	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		panic("Something very bad just happen.")
	}

	file, err := ioutil.ReadFile(path.Join(path.Dir(filename), "data.json"))

	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(file, &DB); err != nil {
		panic(err)
	}
}
