package repository

import (
	"dialog-interview/src/database"
	"dialog-interview/src/errs"
	"dialog-interview/src/models"
	"dialog-interview/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/lib/pq"
)

func CreateBookRepository(book models.Book) error {
	Conn, err := database.Conn()
	if err != nil {
		return fmt.Errorf("%s %w", errs.ERR_DB_CONN, err)
	}

	query := `INSERT INTO public.books (book_name, edition, publication_year, authors) VALUES($1, $2, $3, $4);
	`
	_, err = Conn.Exec(query, book.Name, book.Edition, book.PublicationYear, pq.Array(book.Authors))
	if err != nil {
		return fmt.Errorf("%s %w", errs.ERR_EXEC_QUERY, err)
	}

	defer Conn.Close()
	return nil
}

func ReadBooksRepository(name, edition, publicationYear, author string) ([]models.Book, error) {
	Conn, err := database.Conn()
	if err != nil {
		return nil, fmt.Errorf("%s %w", errs.ERR_DB_CONN, err)
	}

	query := `SELECT id, book_name, edition, publication_year, authors 
	FROM public.books WHERE
	`
	initLen := len(query)

	if len(name) > 0 {
		query += fmt.Sprintf(`book_name = '%s' AND `, name)
	}

	if len(edition) > 0 {
		query += fmt.Sprintf(`edition = %s AND `, edition)
	}

	if len(publicationYear) > 0 {
		query += fmt.Sprintf(`publication_year = %s AND `, publicationYear)
	}

	if len(author) > 0 {
		query += fmt.Sprintf(`%s = any (authors) AND `, author)
	}

	query = utils.SanitizeSelectQuery(query, initLen)

	rows, err := Conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("%s %w", errs.ERR_EXEC_QUERY, err)
	}

	books := make([]models.Book, 0)
	for rows.Next() {
		var book models.Book
		rows.Scan(
			&book.ID,
			&book.Name,
			&book.Edition,
			&book.PublicationYear,
			pq.Array(&book.Authors),
		)
		books = append(books, book)
	}

	defer Conn.Close()
	return books, nil
}

func UpdateBooksRepository(book models.Book) error {
	Conn, err := database.Conn()
	if err != nil {
		return fmt.Errorf("%s %w", errs.ERR_DB_CONN, err)
	}

	var args string
	query := `UPDATE public.books SET `
	if len(book.Name) > 0 {
		args = fmt.Sprintf(`book_name='%s',`, book.Name)
		query += args
	}

	if book.Edition > 0 {
		args = fmt.Sprintf(`edition=%v,`, book.Edition)
		query += args
	}

	if book.PublicationYear > 0 {
		args = fmt.Sprintf(`publication_year=%v,`, book.PublicationYear)
		query += args
	}

	if len(book.Authors) > 0 {
		var Sauthors []string
		for _, a := range book.Authors {
			s := strconv.FormatInt(int64(a), 10)
			Sauthors = append(Sauthors, s)
		}

		args = fmt.Sprintf(`authors='{%s}',`, strings.Join(Sauthors, ","))
		query += args
	}

	query += ` WHERE id=$1`
	query = utils.SanitizeUpdateQuery(query)

	_, err = Conn.Exec(query, book.ID)
	if err != nil {
		return fmt.Errorf("%s %w", errs.ERR_EXEC_QUERY, err)
	}

	defer Conn.Close()
	return nil
}

func DeleteBookRepository(ID string) error {
	Conn, err := database.Conn()
	if err != nil {
		return fmt.Errorf("%s %w", errs.ERR_DB_CONN, err)
	}

	if ID == "0" {
		return fmt.Errorf("ID string need be greater then 0 : %w", err)
	}

	query := `DELETE FROM public.books WHERE id=$1;`

	_, err = Conn.Exec(query, ID)
	if err != nil {
		return fmt.Errorf("%s %w", errs.ERR_EXEC_QUERY, err)
	}

	defer Conn.Close()
	return nil
}
