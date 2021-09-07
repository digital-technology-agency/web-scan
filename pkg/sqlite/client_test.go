package sqlite

import (
	"testing"
)

func TestConnect(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Test connect to DB",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Connect()
			if (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestCreateTable(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Create table",
			args: args{
				query: "CREATE TABLE IF NOT EXISTS test (id INTEGER NOT NULL PRIMARY KEY, name TEXT);",
			},
			wantErr: false,
		},
		{
			name: "Create table",
			args: args{
				query: "CREATE TABLE IF NOT EXISTS Pages (id INTEGER NOT NULL PRIMARY KEY, url TEXT);",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateTable(tt.args.query); (err != nil) != tt.wantErr {
				t.Errorf("CreateTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
