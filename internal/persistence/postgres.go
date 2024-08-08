package persistence

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectToDB() (*sql.DB, error) {
	connStr := "user=postgres password=@postgresql dbname=bank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return nil, err
	}
	return db, nil
}
