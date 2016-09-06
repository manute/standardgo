package http_test

import (
	"mytests/api/domain"
	myhttp "mytests/api/http"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/kataras/iris"
	"github.com/valyala/fasthttp"
)

var accounts = []*domain.Account{&domain.Account{"test"}}

type accountRepoTest struct{}

func (a *accountRepoTest) Find() ([]*domain.Account, error) {
	return accounts, nil
}

// IrisHandler creates fasthttp.RequestHandler using Iris web framework.
func IrisHandler() fasthttp.RequestHandler {
	restService := myhttp.NewAccountRestService(&accountRepoTest{})

	api := iris.New()
	api.Get("/accounts", restService.List)

	return api.ListenVirtual().Handler
}

func irisTester(t *testing.T) *httpexpect.Expect {
	handler := IrisHandler()

	return httpexpect.WithConfig(httpexpect.Config{
		BaseURL: "http://example.com",
		Client: &http.Client{
			Transport: httpexpect.NewFastBinder(handler),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})
}

func TestListAccounts(t *testing.T) {
	e := irisTester(t)

	e.GET("/accounts").Expect().Status(http.StatusOK).JSON().Equal(accounts)

}
