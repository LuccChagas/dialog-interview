package repository

import (
	"dialog-interview/src/models"
	"reflect"
	"testing"
)

func TestCreateBookRepository(t *testing.T) {
	type args struct {
		book models.Book
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{book: models.Book{
				ID:              1,
				Name:            "Livro",
				Edition:         2,
				PublicationYear: 2014,
				Authors:         make([]uint32, 3),
			}},
		},
		{
			name:    "Fail",
			args:    args{book: models.Book{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateBookRepository(tt.args.book); (err != nil) != tt.wantErr {
				t.Errorf("CreateBookRepository() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReadBooksRepository(t *testing.T) {
	type args struct {
		name            string
		edition         string
		publicationYear string
		author          string
	}
	tests := []struct {
		name    string
		args    args
		want    []models.Book
		wantErr bool
	}{
		{
			name: "Sucess",
			args: args{
				name: "Bethel",
			},
			want:    []models.Book{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadBooksRepository(tt.args.name, tt.args.edition, tt.args.publicationYear, tt.args.author)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadBooksRepository() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadBooksRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateBooksRepository(t *testing.T) {
	type args struct {
		book models.Book
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				book: models.Book{
					ID:              1,
					Name:            "O Livro Top",
					Edition:         2,
					PublicationYear: 2017,
					Authors:         make([]uint32, 2),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateBooksRepository(tt.args.book); (err != nil) != tt.wantErr {
				t.Errorf("UpdateBooksRepository() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteBookRepository(t *testing.T) {
	type args struct {
		ID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				ID: "1",
			},
			wantErr: false,
		},
		{
			name: "Fail",
			args: args{
				ID: "0",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteBookRepository(tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("DeleteBookRepository() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
