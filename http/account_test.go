package http_test

import (
	"mytests/api/domain"
	myhttp "mytests/api/http"
	"mytests/api/tester"
	"net/http"
	"testing"
)

var accounts = []*domain.Account{&domain.Account{"test"}}

type accountRepoTest struct{}

func (a *accountRepoTest) Find() ([]*domain.Account, error) {
	return accounts, nil
}

func TestListAccounts(t *testing.T) {
	restService := myhttp.NewAccountRestService(&accountRepoTest{})
	routesTest := []*tester.RoutesTest{
		&tester.RoutesTest{
			Method:  "GET",
			Path:    "/accounts",
			Handler: restService.List,
		},
	}

	e := tester.IrisTester(routesTest, t)

	e.GET("/accounts").Expect().Status(http.StatusOK).JSON().Equal(accounts)

}
