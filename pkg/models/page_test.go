package models

import (
	"fmt"
	"github.com/digital-technology-agency/web-scan/pkg/database"
	"github.com/digital-technology-agency/web-scan/pkg/database/sqlite"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

var (
	testPage = Page{
		Title:       "Заголовок-1",
		Description: "Описание страницы",
		Url:         "http://test.ru",
		Robots:      "*",
		Sitemap:     "<teg>Новый тег</teg>",
	}
	testSqliteDb = sqlite.SqLite{}
)

func TestPage_CreateTable(t *testing.T) {
	type fields struct {
		Title       string
		Description string
		Url         string
		Robots      string
		Sitemap     string
	}
	type args struct {
		dbService database.DbService
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Create table SQLite",
			fields: fields{
				Title:       testPage.Title,
				Description: testPage.Description,
				Url:         testPage.Url,
				Robots:      testPage.Robots,
				Sitemap:     testPage.Sitemap,
			},
			args: args{
				dbService: testSqliteDb,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Page{
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Url:         tt.fields.Url,
				Robots:      tt.fields.Robots,
				Sitemap:     tt.fields.Sitemap,
			}
			if err := p.CreateTable(tt.args.dbService); (err != nil) != tt.wantErr {
				t.Errorf("CreateTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPage_DropTable(t *testing.T) {
	type fields struct {
		Title       string
		Description string
		Url         string
		Robots      string
		Sitemap     string
	}
	type args struct {
		dbService database.DbService
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Drop table",
			fields: fields{
				Title:       testPage.Title,
				Description: testPage.Description,
				Url:         testPage.Url,
				Robots:      testPage.Robots,
				Sitemap:     testPage.Sitemap,
			},
			args: args{
				dbService: testSqliteDb,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Page{
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Url:         tt.fields.Url,
				Robots:      tt.fields.Robots,
				Sitemap:     tt.fields.Sitemap,
			}
			if err := p.DropTable(tt.args.dbService); (err != nil) != tt.wantErr {
				t.Errorf("DropTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPage_Insert(t *testing.T) {
	type fields struct {
		Title       string
		Description string
		Url         string
		Robots      string
		Sitemap     string
	}
	type args struct {
		dbService database.DbService
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Insert ",
			fields: fields{
				Title:       testPage.Title,
				Description: testPage.Description,
				Url:         testPage.Url,
				Robots:      testPage.Robots,
				Sitemap:     testPage.Sitemap,
			},
			args: args{
				dbService: testSqliteDb,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Page{
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Url:         tt.fields.Url,
				Robots:      tt.fields.Robots,
				Sitemap:     tt.fields.Sitemap,
			}
			if err := p.Insert(tt.args.dbService); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPage_Update(t *testing.T) {
	type fields struct {
		Title       string
		Description string
		Url         string
		Robots      string
		Sitemap     string
	}
	type args struct {
		dbService database.DbService
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Update",
			fields: fields{
				Title:       fmt.Sprintf("Заголовок - %s", uuid.NewString()),
				Description: fmt.Sprintf("Описание - %s", uuid.NewString()),
				Url:         testPage.Url,
				Robots:      testPage.Robots,
				Sitemap:     testPage.Sitemap,
			},
			args: args{
				dbService: testSqliteDb,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Page{
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Url:         tt.fields.Url,
				Robots:      tt.fields.Robots,
				Sitemap:     tt.fields.Sitemap,
			}
			if err := p.Update(tt.args.dbService); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPage_Delete(t *testing.T) {
	type fields struct {
		Title       string
		Description string
		Url         string
		Robots      string
		Sitemap     string
	}
	type args struct {
		dbService database.DbService
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Delete",
			fields: fields{
				Url: testPage.Url,
			},
			args: args{
				dbService: testSqliteDb,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Page{
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Url:         tt.fields.Url,
				Robots:      tt.fields.Robots,
				Sitemap:     tt.fields.Sitemap,
			}
			if err := p.Delete(tt.args.dbService); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPage_SelectAll(t *testing.T) {
	type fields struct {
		Title       string
		Description string
		Url         string
		Robots      string
		Sitemap     string
	}
	type args struct {
		dbService database.DbService
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Page
		wantErr bool
	}{
		{
			name:   "Select all",
			fields: fields{},
			args: args{
				dbService: testSqliteDb,
			},
			want:    []Page{testPage},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Page{
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Url:         tt.fields.Url,
				Robots:      tt.fields.Robots,
				Sitemap:     tt.fields.Sitemap,
			}
			got, err := p.SelectAll(tt.args.dbService)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}
