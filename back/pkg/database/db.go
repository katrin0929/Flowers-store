package database

import (
	_ "github.com/lib/pq"

	"database/sql"
	"fmt"
)

func NewDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=localhost port=5432 user=%s password=%s dbname=%s sslmode=disable",
		"user", "password", "postgres")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
