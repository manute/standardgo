package domain

type Account struct {
	Username string `db:"username"`
}

type AccountService interface {
	Find() ([]*Account, error)
}
