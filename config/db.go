package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	HOST     = "127.0.0.1"
	PORT     = "5432"
	USER     = "postgres"
	PASSWORD = "postgres"
	DBNAME   = "example"
	SSLMODE  = "disable"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", HOST, PORT, USER, PASSWORD, DBNAME, SSLMODE))

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
