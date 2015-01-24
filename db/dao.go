package db

import (
	"errors"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func connect() (*sqlx.DB, error) {
	connString := os.Getenv("SQLX_POSTGRES")
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
	if _, err := tx.NamedExec("INSERT INTO sites (name, has_picture, has_engraving, other, is_historic, year_of_discovery, city_id, circuit_id, national_park_id, location_id) VALUES (:name, :has_picture, :has_engraving, :other, :is_historic, :year_of_discovery, :city_id, :circuit_id, :national_park_id, :location_id)", site); err != nil {
		log.Fatalln(err)
		return err
	}
	tx.Commit()

	return nil
}
