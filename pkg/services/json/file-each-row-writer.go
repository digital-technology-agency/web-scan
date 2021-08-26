package json

import (
	"encoding/json"
	"fmt"
	"os"
)

// EachRowWriter each row writer.
type EachRowWriter struct {
	file *os.File
}

// NewEachRowWriter new writer.
func NewEachRowWriter(path string) (*EachRowWriter, error) {
	create, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return &EachRowWriter{file: create}, nil
}

// New new writer.
func (w *EachRowWriter) New(path string) (*EachRowWriter, error) {
	create, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return &EachRowWriter{file: create}, nil
}

// WriteLine write line.
func (w *EachRowWriter) WriteLine(row interface{}) error {
	byteRow, err := json.Marshal(row)
	if err != nil {
		return err
	}
	_, err = w.file.WriteString(fmt.Sprintf("%s\n", string(byteRow)))
	if err != nil {
		return err
	}
	return nil
}
