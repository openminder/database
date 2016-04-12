package database

import (
	"database/sql"
	"fmt"

	// The MYSQL driver import

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// DBInformation describes the access data to the database
type DBInformation struct {
	DBName    string
	User      string
	Password  string
	Host      string
	Port      string
	Collation string
}

// GetMysqlString returns the string to connect to a mysql database
func (dbi *DBInformation) GetMysqlString() string {
	if dbi.Host == "" {
		dbi.Host = "localhost"
	}
	if dbi.Collation == "" {
		dbi.Collation = "utf8_general_ci"
	}
	if dbi.DBName == "" {
		panic("Error: Database name missing.")
	}
	if dbi.User == "" {
		return "tcp(" + dbi.Host + ":" + dbi.Port + ")" + "/" + dbi.DBName + "?collation=" + dbi.Collation
	}
	return dbi.User + ":" + dbi.Password + "@tcp(" + dbi.Host + ":" + dbi.Port + ")" + "/" + dbi.DBName + "?collation=" + dbi.Collation
}

// GetDB returning current DB connection
func GetDB(dbi DBInformation) *sql.DB {
	var err error
	if db == nil {
		fmt.Println(dbi.GetMysqlString())
		db, err = sql.Open("mysql", dbi.GetMysqlString())
		if err != nil {
			panic(err)
		}
	}

	return db
}
