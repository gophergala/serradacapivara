package entities

type Site struct {
	Id              int
	Name            string
	HasPicture      bool
	HasEngraving    bool
	Other           bool
	IsHistoric      bool
	YearOfDiscovery int

	CityId         int
	CircuitId      int
	NationalParkId int
	LocationId     int
}

type City struct {
	Id   int
	Name string
}

type Circuit struct {
	Id   int
	Name string
}

type NationalPack struct {
	Id   int
	Name string
}

type Location struct {
	Id   int
	Name string
}
