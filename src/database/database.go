package database

import (
	"api-rede-social/src/config"
	"database/sql"

	_ "github.com/lib/pq" //Driver
)

// Connect create a conection with database
func Connect() (*sql.DB, error) {

	db, err := sql.Open("postgres", config.ConnStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
