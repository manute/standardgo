package domain_test

import (
	"mytests/api/domain"
	"testing"
)

type TestRepoAccount struct{}

func NewTestRepo() *TestRepoAccount {
	return &TestRepoAccount{}
}

func (t *TestRepoAccount) Find() ([]domain.Account, error) {
	return nil, nil
}

func TestFind(t *testing.T) {
	service := NewTestRepo()
	service.Find()
}
