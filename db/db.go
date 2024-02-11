package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDb() *sql.DB {
	connection := "user=postgres dbname=curso_go password=root123 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
	return db
}
