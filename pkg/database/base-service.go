package database

import "github.com/jmoiron/sqlx"

type DbService interface {
	Connect() (*sqlx.DB, error)
	Execute(query string) error
}

type DbEntity interface {
	GetTableName() string
	CreateTable(service DbService) error
	DropTable(service DbService) error
	Insert(service DbService) error
	Update(service DbService) error
	Delete(service DbService) error
}
