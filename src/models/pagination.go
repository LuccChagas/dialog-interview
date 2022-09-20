package models

type Pagination struct {
	Total uint `json:"total"`
	Pages uint `json:"pages"`
}

type AuthorsResultSet struct {
	Pagination *Pagination `json:"pagination"`
	Authors    []Authors   `json:"authors"`
}
