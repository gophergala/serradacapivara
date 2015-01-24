package db

type Site struct {
	Id              int
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
}
