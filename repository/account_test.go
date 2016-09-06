package repository_test

import (
	"standardgo/db"
	"standardgo/repository"
	. "standardgo/tester"
	"testing"
)

var accountRepo *repository.AccountRepository

func TestMain(m *testing.M) {
	db, err := db.NewPostgreDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	accountRepo = &repository.AccountRepository{Db: db}
	m.Run()
}

func TestAccountFind(t *testing.T) {
	accs, err := accountRepo.Find()
	Ok(t, err)
	Equals(t, 2, len(accs))
}
