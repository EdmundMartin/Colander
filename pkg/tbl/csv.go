package tbl

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/EdmundMartin/Colander/pkg/bst"
	"io"
	"log"
	"os"
)

func LoadFromCSV(filepath string, tblname string, mapping map[int]ColumnInfo)  error {

	emptyTable := NewTable(tblname)

	treeMap := make(map[int]*bst.BST)
	for key, value := range mapping {
		treeMap[key] = bst.NewBST(value.ColumnName)
	}

	csvFile, err := os.Open(filepath)
	if err != nil {
		return err
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	pk := 0
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		for key, _ := range mapping {
			if key < len(line) {
				ttree := treeMap[key]
				ttree.Insert(pk, line[key])
			}
		}
		pk++
	}
	for _, tree := range treeMap {
		fileName := fmt.Sprintf("%s_%s.db_col", tblname, tree.Name)
		emptyTable.ColumnFiles[tree.Name] = fileName
		tree.Save(fileName)
	}

	emptyTable.Save()

	return nil
}