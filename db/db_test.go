package db_test

import (
	"standardgo/db"
	"testing"
)

func TestPing(t *testing.T) {
	db, err := db.NewPostgreDB()
	if err != nil {
		t.Error(err)
	}
	if err = db.Ping(); err != nil {
		t.Error(err)
	}

}
