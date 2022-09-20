package models

type Book struct {
	ID              uint32   `json:"id,omitempty"`
	Name            string   `json:"name,omitempty"`
	Edition         uint32   `json:"edition,omitempty"`
	PublicationYear uint32   `json:"publication_year,omitempty"`
	Authors         []uint32 `json:"authors,omitempty"`
}
