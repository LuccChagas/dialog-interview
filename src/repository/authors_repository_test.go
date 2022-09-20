package repository

import (
	"dialog-interview/src/models"
	"testing"
)

func TestInsertAuthorsCSV(t *testing.T) {

	type args struct {
		names [][]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				names: [][]interface{}{},
			},
			wantErr: false,
		},
		{
			name:    "Fail",
			args:    args{make([][]interface{}, 500)},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InsertAuthorsCSV(tt.args.names); (err != nil) != tt.wantErr {
				t.Errorf("InsertAuthorsCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAuthors(t *testing.T) {

	type args struct {
		page string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.AuthorsResultSet
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Sucess",
			args:    args{page: "50"},
			want:    &models.AuthorsResultSet{},
			wantErr: false,
		},
		{
			name:    "zero page",
			args:    args{page: ""},
			want:    &models.AuthorsResultSet{},
			wantErr: false,
		},
		{
			name:    "Error",
			args:    args{page: "0"},
			want:    &models.AuthorsResultSet{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetAuthors(tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAuthors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
