package repository

import (
	"dialog-interview/src/database"
	"dialog-interview/src/errs"
	"dialog-interview/src/models"
	"fmt"
	"time"

	"github.com/jackc/pgx"
	"github.com/labstack/gommon/log"
)

func InsertAuthorsCSV(names [][]interface{}) error {
	Conn, err := database.Conn()
	if err != nil {
		return fmt.Errorf("%s %w", errs.ERR_DB_CONN, err)
	}

	start := time.Now()
	count, err := Conn.CopyFrom(pgx.Identifier{"authors"}, []string{"author_name"}, pgx.CopyFromRows(names))
	log.Info("Done in :", time.Since(start))
	if err != nil {
		return fmt.Errorf("conn.CopyFrom %w", err)
	}

	fmt.Println("Inserted Rows:", count)
	defer Conn.Close()
	return nil
}

func GetAuthors(page string) (*models.AuthorsResultSet, error) {
	Conn, err := database.Conn()
	if err != nil {
		return nil, fmt.Errorf("%s %w", errs.ERR_DB_CONN, err)
	}

	if page == "0" {
		return nil, fmt.Errorf("page string need be greater then 0 : %w", err)
	}

	query := fmt.Sprintf(`select id, author_name from authors limit 500 OFFSET (%s - 1) * 500;`, page)

	rows, err := Conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("a error ocurred to Exec Conn.Query: %w", err)
	}

	var authors []models.Authors
	for rows.Next() {
		author := new(models.Authors)
		if err := rows.Scan(
			&author.ID,
			&author.Name,
		); err != nil {
			return nil, fmt.Errorf("a error ocurred to Bind Authors return data: %w", err)
		}
		authors = append(authors, *author)
	}

	query = `select COUNT(*) as total, COUNT(*) / 500 as pages from authors`
	rows, err = Conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("a error ocurred to Exec Conn.Query: %w", err)
	}

	pagination := new(models.Pagination)
	for rows.Next() {
		if err := rows.Scan(
			&pagination.Total,
			&pagination.Pages,
		); err != nil {
			return nil, fmt.Errorf("a error ocurred to Bind Authors - Pagination return data: %w", err)
		}
	}

	rs := models.AuthorsResultSet{
		Pagination: pagination,
		Authors:    authors,
	}

	defer Conn.Close()
	return &rs, nil
}

func GetAuthorsByName(name string) (*[]models.Authors, error) {
	Conn, err := database.Conn()
	if err != nil {
		return nil, fmt.Errorf("%s %w", errs.ERR_DB_CONN, err)
	}

	query := `select id, author_name from authors where author_name LIKE $1`

	rows, err := Conn.Query(query, "%"+name+"%")
	if err != nil {
		return nil, fmt.Errorf("a error ocurred to Exec Conn.Query: %w", err)
	}

	var authors []models.Authors
	for rows.Next() {
		author := new(models.Authors)
		if err := rows.Scan(
			&author.ID,
			&author.Name,
		); err != nil {
			return nil, fmt.Errorf("a error ocurred to Bind Authors return data: %w", err)
		}
		authors = append(authors, *author)
	}

	return &authors, nil
}
