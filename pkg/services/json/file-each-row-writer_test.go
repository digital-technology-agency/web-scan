package json

import (
	"os"
	"testing"
)

var testFileName = "test-file"

func TestEachRowWriter_WriteLine(t *testing.T) {
	type testRow struct {
		TestField       string
		TestSecondField int
	}
	osTempFile, err := os.Create(testFileName)
	if err != nil {
		t.Errorf("Crate test file error = %v", err)
	}
	type fields struct {
		file *os.File
	}
	type args struct {
		row interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "WriteLine",
			fields: fields{
				file: osTempFile,
			},
			args: args{
				row: testRow{
					TestField:       "Поле 1",
					TestSecondField: 100,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &EachRowWriter{
				file: tt.fields.file,
			}
			if err := w.WriteLine(tt.args.row); (err != nil) != tt.wantErr {
				t.Errorf("WriteLine() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	os.RemoveAll(testFileName)
}

func TestNewEachRowWriter(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    *EachRowWriter
		wantErr bool
	}{
		{
			name: "New writer",
			args: args{
				path: testFileName,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEachRowWriter(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewEachRowWriter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == tt.want {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
	os.RemoveAll(testFileName)
}
