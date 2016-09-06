package http_test

import (
	"net/http"
	"standardgo/domain"
	myhttp "standardgo/http"
	"standardgo/tester"
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
