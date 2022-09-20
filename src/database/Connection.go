package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
)

type DB struct {
	Conn *pgx.Conn
}

func LoadEnv() {
	err := godotenv.Load("/Users/luccas/go/projects/dialog-interview/.env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func Conn() (*pgx.Conn, error) {
	LoadEnv()

	DSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("HOSTNAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("PORT"))

	config, err := pgx.ParseDSN(DSN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error to parse DSN: %v\n", err)
		os.Exit(1)
	}

	conn, err := pgx.Connect(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn, err
}

// func NewConnection() (DB, error) {
// 	pgx, err := conn()
// 	if err != nil {
// 		return DB{}, err
// 	}

// 	Connect = DB{Conn: pgx}
// 	return Connect, nil
// }
