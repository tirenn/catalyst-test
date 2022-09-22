package configs

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// InitDB will initialize DB connection
func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./test.db?_foreign_keys=on")
	if err != nil {
		panic(err)
	}

	return db
}
