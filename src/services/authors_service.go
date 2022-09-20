package services

import (
	"dialog-interview/src/models"
	"dialog-interview/src/repository"
	"strings"
)

func GetAuthorsService(page string) (*models.AuthorsResultSet, error) {

	authors, err := repository.GetAuthors(page)
	if err != nil {
		return nil, err
	}

	return authors, nil
}

func GetAuthorsByNameService(name string) (*[]models.Authors, error) {

	name = strings.Title(strings.ToLower(name))

	authors, err := repository.GetAuthorsByName(name)
	if err != nil {
		return nil, err
	}

	return authors, nil
}
