package tbl

import "github.com/EdmundMartin/Colander/pkg/bst"

type ColumnFile struct {
	Filename string
	Tree *bst.BST
	ValType string // Lets make this an enum
}

type Table struct {
	ColMapping map[string]ColumnFile
}
