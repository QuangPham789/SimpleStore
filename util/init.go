package util

import (
	"database/sql"
	"log"
)

const (
	dbDriver = "postgres"
	// dbSource = "postgressql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	dbSource = "host=localhost port=5432 user=root password=secret dbname=simple_store sslmode=disable"
)

func InitDB() *sql.DB {
	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("an error '%s' was not expected", err)
	}
	return db
}
