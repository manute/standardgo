package tester

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/kataras/iris"
	"github.com/valyala/fasthttp"
)

type RoutesTest struct {
	Method  string
	Path    string
	Handler func(c *iris.Context)
}

// IrisHandler creates fasthttp.RequestHandler using Iris web framework.
func irisHandler(routes []*RoutesTest) fasthttp.RequestHandler {

	api := iris.New()
	for _, route := range routes {
		api.HandleFunc(route.Method, route.Path, route.Handler)
	}

	return api.ListenVirtual().Handler
}

func IrisTester(paths []*RoutesTest, t *testing.T) *httpexpect.Expect {
	handler := irisHandler(paths)

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
