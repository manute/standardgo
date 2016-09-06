package db_test

import (
	"mytests/api/db"
	"testing"
)

func TestPing(t *testing.T) {
	db, err := db.NewPSQL()
	if err != nil {
		t.Error(err)
	}
	if err = db.Ping(); err != nil {
		t.Error(err)
	}

}
