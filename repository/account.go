package repository

import (
	"standardgo/db"
	"standardgo/domain"
)

type AccountRepository struct {
	Db *db.PostgreDB
}

func (ar *AccountRepository) Find() ([]*domain.Account, error) {
	accounts := []*domain.Account{}
	err := ar.Db.Select(&accounts, "SELECT username FROM account LIMIT 2")
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
