package database

import "database/sql"

// The MYSQL driver import
import _ "github.com/go-sql-driver/mysql"

var db *sql.DB

// GetDB returning current DB connection
func GetDB() *sql.DB {
	var err error
	if db == nil {
		db, err = sql.Open("mysql", "")
		if err != nil {
			panic(err)
		}
	}

	return db
}
