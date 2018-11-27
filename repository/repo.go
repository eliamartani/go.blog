package repository

import (
	"database/sql"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// Repo is a base concept for using database
type Repo struct {
	getQuery    string
	listQuery   string
	insertQuery string
	updateQuery string
	deleteQuery string
}

const connString = "root:123456@/godb"

// Connect creates a connection with mysql
func (r Repo) Connect() *sql.DB {
	db, err := sql.Open("mysql", connString)

	if err != nil {
		panic(err.Error())
	}

	return db
}
