package main

import (
	"standardgo/db"
	"standardgo/http"
	"standardgo/repository"

	"github.com/iris-contrib/logger"
	mLogger "github.com/iris-contrib/middleware/logger"
	"github.com/kataras/iris"
)

func main() {
	db, err := db.NewPSQL()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	accountRestService := http.NewAccountRestService(&repository.AccountRepository{db})

	theLogger := logger.New(logger.DefaultConfig())

	iris.Use(mLogger.New(theLogger))

	iris.Get("/accounts", accountRestService.List)

	iris.Listen(":8080")
}
