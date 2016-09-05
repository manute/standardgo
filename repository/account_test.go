package repository_test

import (
	"mytests/api/db"
	"mytests/api/repository"
	. "mytests/api/test"
	"testing"
)

var accountRepo *repository.AccountRepository

func TestMain(m *testing.M) {
	db, err := db.NewPSQL()
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
