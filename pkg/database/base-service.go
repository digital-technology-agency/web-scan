package database

import "github.com/jmoiron/sqlx"

// DbService database interface.
type DbService interface {
	Connect() (*sqlx.DB, error)
	Execute(query string) error
}

// DbEntity entity interface.
type DbEntity interface {
	GetTableName() string
	CreateTable(service DbService) error
	DropTable(service DbService) error
	AddOrUpdate(service DbService) error
	Insert(service DbService) error
	Update(service DbService) error
	Delete(service DbService) error
}
