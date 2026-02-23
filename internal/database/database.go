package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgresConnection(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to PostgreSQL")

	return db, nil
}
