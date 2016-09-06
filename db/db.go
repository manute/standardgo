package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgreDB struct {
	*sqlx.DB
}

func NewPostgreDB() (*PostgreDB, error) {
	db, err := sqlx.Connect("postgres", "user=postgres password='postgres' dbname=postgres")
	if err != nil {
		return nil, err
	}

	return &PostgreDB{db}, nil
}
