package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbDatabase := os.Getenv("DB_DATABASE")
	fmt.Println(dbPassword)
	connStr := "postgres://" + dbUsername + ":" + dbPassword + "@" + dbHost + "/" + dbDatabase + "?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Error connecting to db: %v", err)
		db.Close()
	}
	fmt.Println("Connected to db")
	return db
}
