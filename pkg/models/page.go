package models

import (
	"context"
	"fmt"

	"github.com/digital-technology-agency/web-scan/pkg/database"
)

// PageTableName ...
const PageTableName = "pages"

/*Page type of page*/
type Page struct {
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	URL         string `json:"url" db:"url"`
	Robots      string `json:"robots" db:"robots"`
	Sitemap     string `json:"sitemap" db:"sitemap"`
}

// GetTableName get table name
func (p Page) GetTableName() string {
	return PageTableName
}

// CreateTable create table
func (p Page) CreateTable(dbService database.DbService) error {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s ("+
		"url TEXT PRIMARY KEY NOT NULL,"+
		"title TEXT,"+
		"description TEXT,"+
		"robots TEXT,"+
		"sitemap TEXT"+
		")", PageTableName)
	return dbService.Execute(query)
}

// DropTable drop table
func (p Page) DropTable(dbService database.DbService) error {
	query := fmt.Sprintf("DROP TABLE IF EXISTS %s", PageTableName)
	return dbService.Execute(query)
}

// SelectAll select all rows
func (p Page) SelectAll(dbService database.DbService) ([]Page, error) {
	query := fmt.Sprintf("SELECT * FROM %s", PageTableName)
	connect, err := dbService.Connect()
	if err != nil {
		return nil, err
	}
	defer connect.Close()
	result := []Page{}
	err = connect.Select(&result, query)
	return result, err
}

// AddOrUpdate add or update row
func (p Page) AddOrUpdate(dbService database.DbService) error {
	connect, err := dbService.Connect()
	if err != nil {
		return err
	}
	defer connect.Close()
	destValue := &Page{}
	err = connect.GetContext(context.Background(), destValue, fmt.Sprintf("SELECT * from %s WHERE url=$1 LIMIT 1", PageTableName), p.URL)
	if err != nil {
		return p.Insert(dbService)
	}
	return p.Update(dbService)
}

// Insert insert data to table
func (p Page) Insert(dbService database.DbService) error {
	connect, err := dbService.Connect()
	if err != nil {
		return err
	}
	defer connect.Close()
	query := fmt.Sprintf("INSERT INTO %s (title, description, url, robots, sitemap) VALUES (:title, :description, :url, :robots, :sitemap)", PageTableName)
	_, err = connect.NamedExec(query, p)
	return err
}

// Update update data in table
func (p Page) Update(dbService database.DbService) error {
	connect, err := dbService.Connect()
	if err != nil {
		return err
	}
	defer connect.Close()
	query := fmt.Sprintf("UPDATE %s SET title=:title, description=:description, robots=:robots, sitemap=:sitemap WHERE url=:url", PageTableName)
	_, err = connect.NamedExec(query, p)
	return err
}

// Delete delete data from table
func (p Page) Delete(dbService database.DbService) error {
	connect, err := dbService.Connect()
	if err != nil {
		return err
	}
	defer connect.Close()
	query := fmt.Sprintf(`DELETE FROM %s WHERE url=:url`, PageTableName)
	_, err = connect.NamedExec(query, p)
	return err
}
