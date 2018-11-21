package database

import (
	"database/sql"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

const connString = "root:@/godb"

// Connect creates a connection with mysql
func Connect() *sql.DB {
	db, err := sql.Open("mysql", connString)

	if err != nil {
		panic(err.Error())
	}

	return db
}
