package db

import (
	"errors"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func connect() (*sqlx.DB, error) {
	connString := os.Getenv("DATABASE_URL")
	return sqlx.Connect("postgres", connString)
}

func SaveSite(site Site) error {
	d, err := connect()
	defer d.Close()

	if err != nil {
		log.Fatalln(err)
		return errors.New("Error acessing database")
	}

	tx := d.MustBegin()
	if _, err := tx.NamedExec(`INSERT INTO Sites (
                                  Name,
                                  HasPicture,
                                  HasEngraving,
                                  Other,
                                  IsHistoric, 
                                  YearOfDiscovery,
                                  City,
                                  Circuit,
                                  NationalPark,
                                  Location)
                                  VALUES (
                                     :Name,
                                     :HasPicture,
                                     :HasEngraving,
                                     :Other,
                                     :IsHistoric,
                                     :YearOfDiscovery,
                                     :City,
                                     :Circuit,
                                     :NationalPark,
                                     :Location
                                  )`, site); err != nil {
		log.Fatalln(err)
		return err
	}
	tx.Commit()

	return nil
}
