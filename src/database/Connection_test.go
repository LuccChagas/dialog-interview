package database

import (
	"reflect"
	"testing"

	"github.com/jackc/pgx"
)

func TestConn(t *testing.T) {
	tests := []struct {
		name    string
		want    *pgx.Conn
		wantErr bool
	}{
		{
			name:    "Sucess Connection",
			want:    &pgx.Conn{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Conn()
			if (err != nil) != tt.wantErr {
				t.Errorf("Conn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("Conn() = %v, want %v", got, tt.want)
			}
		})
	}
}
