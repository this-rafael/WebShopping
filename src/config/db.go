package config

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func DatabaseConnect() *sql.DB {
	connectionUrl := "user=store_user password=pgpass123 dbname=store host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connectionUrl)

	if err != nil {
		panic(err)
	}

	return db
}
