package tbl

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

type Table struct {
	Name string
	ColumnFiles map[string]string
}

func NewTable(name string) *Table {
	return &Table{
		Name:        name,
		ColumnFiles: make(map[string]string),
	}
}

func (t *Table) Save() error {
	buf := &bytes.Buffer{}
	fo, _ := os.Create(fmt.Sprintf("%s.db", t.Name))
	defer fo.Close()
	err := gob.NewEncoder(buf).Encode(t)
	if err != nil {
		return err
	}
	_, err = fo.Write(buf.Bytes())
	if err != nil {
		return err
	}
	return nil
}


type ColumnType int

const (
	StringColumn ColumnType = iota
	IntegerColumn
	FloatColumn
)

type ColumnInfo struct {
	ColumnName string
	ColumnType
}
