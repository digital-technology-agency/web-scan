package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// Connect connect to DB.
func Connect() (*sql.DB, error) {
	return sql.Open(`sqlite3`, `./wScan.db`)
}

// CreateTable create table in database.
func CreateTable(query string) error {
	db, err := Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
