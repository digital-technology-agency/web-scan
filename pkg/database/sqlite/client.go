package sqlite

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // sqlite3
)

// SqLite sqlite client.
type SqLite struct {
	db *sqlx.DB
}

// Connect connect to DB.
func (engine SqLite) Connect() (*sqlx.DB, error) {
	open, err := sqlx.Open(`sqlite3`, `./wScan.db`)
	if err != nil {
		return nil, err
	}
	engine.db = open
	return engine.db, nil
}

// Execute create table in database.
func (engine SqLite) Execute(query string) error {
	db, err := engine.Connect()
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
