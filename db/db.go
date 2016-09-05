package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	// *sql.DB
	*sqlx.DB
}

func NewPSQL() (*DB, error) {
	db, err := sqlx.Connect("postgres", "user=postgres password='postgres' dbname=postgres")
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
