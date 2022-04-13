package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DbConnection() *sql.DB {

	// Setting the connections of the DataBase
	connection := "user=postgres dbname=store_project password=zeu$@2022 host=localhost sslmode=disable"

	// Open opens a database specified by its database driver name and a driver-specific data source name
	db, err := sql.Open("postgres", connection)

	// Catching the errors
	if err != nil {
		panic(err.Error())
	}

	return db
}
