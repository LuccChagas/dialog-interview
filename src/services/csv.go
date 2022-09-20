package services

import (
	"dialog-interview/src/repository"
	"encoding/csv"
	"io"
	"os"
	"strings"

	"github.com/labstack/gommon/log"
)

// type Authors struct {
// 	Name string `json:"name"`
// }

const filePath = "/Users/luccas/go/projects/dialog-interview/authors.csv" // GETENV

func ReadAuthorsCSV() {
	rowsToInsert := [][]interface{}{}

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		var splited []string
		for _, rec := range record {
			splited = strings.Split(rec, ";")
		}

		for i := 0; i < len(splited); i++ {
			row := []interface{}{splited[i]}
			rowsToInsert = append(rowsToInsert, row)

		}
	}

	repository.InsertAuthorsCSV(rowsToInsert)
}
