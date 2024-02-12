package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // an underscore tells go that this package will not be used directly but under the hood. In this case by builtin sql package
)

var DB *sql.DB

func initDB() {
	DB, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}