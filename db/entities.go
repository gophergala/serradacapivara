package db

type Site struct {
	Id              int    `db:"id"`
	Name            string `db:"name"`
	HasPicture      bool   `db:"has_picture"`
	HasEngraving    bool   `db:"has_engraving"`
	Other           bool   `db:"other"`
	IsHistoric      bool   `db:"is_historic"`
	YearOfDiscovery int    `db:"year_of_discovery"`

	CityId         int `db:"city_id"`
	CircuitId      int `db:"circuit_id"`
	NationalParkId int `db:"national_park_id"`
	LocationId     int `db:"location_id"`
}

type City struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

type Circuit struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

type NationalPack struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

type Location struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}
