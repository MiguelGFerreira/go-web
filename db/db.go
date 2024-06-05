package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDatabase() *sql.DB {
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	connectionString := fmt.Sprintf("user=%s dbname=loja password=%s host=localhost sslmode=disable", DB_USER, DB_PASS)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err.Error())
	}
	return db
}
