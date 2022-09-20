package services

import (
	"dialog-interview/src/models"
	"dialog-interview/src/repository"
	"errors"
	"strings"
)

func CreateBookService(book models.Book) error {

	book.Name = strings.Title(strings.ToLower(book.Name))

	err := repository.CreateBookRepository(book)
	if err != nil {
		return err
	}
	return nil
}

func ReadBooksService(name, edition, publicationYear, author string) ([]models.Book, error) {

	var finalName string
	splited := strings.Split(name, "_")
	for _, s := range splited {
		s = strings.Title(strings.ToLower(s))
		finalName += " " + s
	}

	finalName = strings.TrimLeft(finalName, " ")

	authors, err := repository.ReadBooksRepository(finalName, edition, publicationYear, author)
	if err != nil {
		return nil, err
	}

	if len(authors) < 1 {
		err := errors.New("There is no finded value for this search")
		return nil, err
	}

	return authors, nil
}

func UpdateBooksService(book models.Book) error {

	book.Name = strings.Title(strings.ToLower(book.Name))

	err := repository.UpdateBooksRepository(book)
	if err != nil {
		return err
	}

	return nil
}

func DeleteBookService(ID string) error {

	err := repository.DeleteBookRepository(ID)
	if err != nil {
		return err
	}

	return nil
}
