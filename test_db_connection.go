package main

import (
	"fmt"
	"log"

	"github.com/fonluc/bluesoft-bank/internal/persistence"
)

func main() {
	db, err := persistence.ConnectToDB()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	fmt.Println("Database connection successful")
	db.Close()
}
